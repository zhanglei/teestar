package api

type UserFollowStatus struct {
	User                 string
	QQ                   string
	Nickname             string
	Hitter               string
	FollowingTargets     []string
	FollowedTargets      []string
	Score                int
	CanFollowTargets     []string
	CanBeFollowedTargets []string
}
