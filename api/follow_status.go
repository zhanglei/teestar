package api

type UserFollowStatus struct {
	User                 string
	FollowingTargets     []string
	FollowedTargets      []string
	Score                int
	CanFollowTargets     []string
	CanBeFollowedTargets []string
}

type FollowEntry struct {
	User     string
	QQ       string
	Nickname string
	Followed bool
}
