package controllers

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
	Score         int
	CanStarRepos  []string
}

type StatusList []*UserTargetStatus
