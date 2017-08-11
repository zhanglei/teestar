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

type FollowEntryList []*FollowEntry

func (list FollowEntryList) Len() int {
	return len(list)
}

func (list FollowEntryList) Less(i, j int) bool {
	if list[i].Followed && !list[j].Followed {
		return true
	} else {
		return false
	}
}

func (list FollowEntryList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
