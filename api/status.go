package api

type UserTargetStatus struct {
	User              string
	Target            string
	QQ                string
	Nickname          string
	Hitter            string
	StarringRepos     []string
	StarredRepos      []string
	Score             int
	CanStarRepos      []string
	CanBeStarredRepos []string
}

type StatusList []*UserTargetStatus

func (list StatusList) Len() int {
	return len(list)
}

func (list StatusList) Less(i, j int) bool {
	if list[i].Score > list[j].Score {
		return true
	} else {
		return false
	}
}

func (list StatusList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}
