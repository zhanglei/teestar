package api

func GetUserFollowingTargets(user string) []string {
	var userTargets []UserFollowingTarget
	err := adapter.engine.Find(&userTargets, &UserFollowingTarget{User: user})
	if err != nil {
		panic(err)
	}

	repos := []string{}
	for _, userTarget := range userTargets {
		repos = append(repos, userTarget.Target)
	}

	return repos
}

func UpdateUserFollowingTargets(user string) bool {
	hitter := GetUserOrHitter(user)

	affected, err := adapter.engine.Delete(&UserFollowingTarget{User: user})
	if err != nil {
		panic(err)
	}

	targets := ListFollowingTargets(hitter)
	userTargets := []UserFollowingTarget{}
	for _, target := range targets {
		userTargets = append(userTargets, UserFollowingTarget{User: user, Target: target})
	}

	affected, err = adapter.engine.Insert(&userTargets)
	if err != nil {
		panic(err)
	}

	return affected != 0
}

func UpdateFollowingTargets() bool {
	affected := false
	users := GetFollowableUsers()

	for _, user := range users {
		if UpdateUserFollowingTargets(user) {
			affected = true
		}
	}

	return affected
}

func IsUserFollowingTarget(user string, target string) bool {
	userFollowingTarget := UserFollowingTarget{User: user, Target: target}
	has, err := adapter.engine.Get(&userFollowingTarget)
	if err != nil {
		panic(err)
	}

	return has
}

func GetUserFollowStatus(user string) UserFollowStatus {
	otherUsers := GetOtherFollowableUsers(user)
	followingTargets := GetIntersect(otherUsers, GetUserFollowingTargets(user))
	canFollowTargets := GetSubtract(otherUsers, followingTargets)

	followedTargets := []string{}
	for _, target := range otherUsers {
		if IsUserFollowingTarget(target, user) {
			followedTargets = append(followedTargets, target)
		}
	}
	followedTargets = GetIntersect(otherUsers, followedTargets)
	canBeFollowedTargets := GetSubtract(otherUsers, followedTargets)

	userFollowStatus := UserFollowStatus{
		FollowingTargets: followingTargets,
		FollowedTargets: followedTargets,
		CanFollowTargets: canFollowTargets,
		CanBeFollowedTargets: canBeFollowedTargets}

	return userFollowStatus
}
