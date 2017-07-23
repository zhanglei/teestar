package controllers

import  "sort"

func getIntersect(a []string, b []string) []string {
	res := []string{}
	for _, ia := range a {
		found := false
		for _, ib := range b {
			if ia == ib {
				found = true
				break
			}
		}

		if found {
			res = append(res, ia)
		}
	}

	return res
}

func getSubtract(a []string, b []string) []string {
	res := []string{}
	for _, ia := range a {
		found := false
		for _, ib := range b {
			if ia == ib {
				found = true
				break
			}
		}

		if !found {
			res = append(res, ia)
		}
	}

	return res
}

func (c *MainController) GetUserTarget() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	c.Data["json"] = getIntersect(targetRepos, userStarringRepos)
	c.ServeJSON()
}

func getUserTargetStatus(user string, target string) UserTargetStatus {
	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	starringRepos := getIntersect(targetRepos, userStarringRepos)

	userRepos := getUserRepos(user)
	targetStarringRepos := getUserStarringRepos(target)
	starredRepos := getIntersect(userRepos, targetStarringRepos)

	score := len(starredRepos) - len(starringRepos)

	canStarRepos := getSubtract(targetRepos, userStarringRepos)

	return UserTargetStatus{StarringRepos: starringRepos, StarredRepos: starredRepos, Score: score, CanStarRepos: canStarRepos}
}

func (c *MainController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = getUserTargetStatus(user, target)
	c.ServeJSON()
}

func getUserTargetPool(user string, target string) []string {
	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	poolRepos := getSubtract(targetRepos, userStarringRepos)
	return poolRepos
}

func (c *MainController) GetUserTargetPool() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = getUserTargetPool(user, target)
	c.ServeJSON()
}

func getUserStatus(user string) StatusList {
	statusList := StatusList{}
	otherUsers := getOtherUsers(user)
	for _, otherUser := range otherUsers {
		status := getUserTargetStatus(user, otherUser)
		statusList = append(statusList, &status)
	}

	sort.Sort(statusList)
	return statusList
}

func (c *MainController) GetUserStatus() {
	user := c.GetString(":user")

	statusList := getUserStatus(user)

	c.Data["json"] = statusList
	c.ServeJSON()
}

func (c *MainController) GetUserRecommend() {
	user := c.GetString(":user")

	repos := []string{}
	statusList := getUserStatus(user)
	for _, status := range statusList {
		repos = append(repos, status.CanStarRepos...)
	}

	c.Data["json"] = repos
	c.ServeJSON()
}
