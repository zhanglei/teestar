package api

import "sort"

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
	otherUsers := GetOtherEnabledUsers(user)
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

func GetUserFollowingStatus(user string) []FollowEntry {
	otherUsers := GetOtherEnabledUsers(user)
	followingTargets := GetIntersect(otherUsers, GetUserFollowingTargets(user))

	followEntries := []FollowEntry{}
	for _, target := range followingTargets {
		if !IsUserFollowingTarget(target, user) {
			followEntries = append(followEntries, FollowEntry{User: target, Followed: false})
		}
	}

	return followEntries
}

func GetUserFollowedStatus(user string) []FollowEntry {
	otherUsers := GetOtherEnabledUsers(user)
	followingTargets := GetIntersect(otherUsers, GetUserFollowingTargets(user))
	followingTargetsMap := NewSet()
	for _, target := range followingTargets {
		followingTargetsMap.Add(target)
	}

	followedTargets := []string{}
	for _, target := range otherUsers {
		if IsUserFollowingTarget(target, user) {
			followedTargets = append(followedTargets, target)
		}
	}
	followedTargets = GetIntersect(otherUsers, followedTargets)

	followEntries := []FollowEntry{}
	for _, target := range followedTargets {
		followed := followingTargetsMap.Has(target)
		followEntries = append(followEntries, FollowEntry{User: target, Followed: followed})
	}

	return followEntries
}

func GetUserCanFollowStatus(user string) FollowEntryList {
	otherUsers := GetOtherFollowableUsers(user)
	followingTargets := GetIntersect(otherUsers, GetUserFollowingTargets(user))
	canFollowTargets := GetSubtract(otherUsers, followingTargets)

	followEntryList := FollowEntryList{}
	for _, target := range canFollowTargets {
		followed := IsUserFollowingTarget(target, user)
		followEntryList = append(followEntryList, &FollowEntry{User: target, Followed: followed})
	}

	sort.Sort(followEntryList)
	return followEntryList
}
