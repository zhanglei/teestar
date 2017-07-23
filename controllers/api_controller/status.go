package api_controller

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
	Score         int
	CanStarRepos  []string
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
