package controllers

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

type UserTargetStatus struct {
	StarringRepos []string
	StarredRepos  []string
	Score         int
}

func getUserTargetStatus(user string, target string) UserTargetStatus {
	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	starringRepos := getIntersect(targetRepos, userStarringRepos)

	userRepos := getUserRepos(user)
	targetStarringRepos := getUserStarringRepos(target)
	starredRepos := getIntersect(userRepos, targetStarringRepos)

	score := len(starredRepos) - len(starringRepos)

	return UserTargetStatus{StarringRepos: starringRepos, StarredRepos: starredRepos, Score: score}
}

func (c *MainController) GetUserTargetStatus() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	c.Data["json"] = getUserTargetStatus(user, target)
	c.ServeJSON()
}

func (c *MainController) GetUserTargetPool() {
	user := c.GetString(":user")
	target := c.GetString(":target")

	targetRepos := getUserRepos(target)
	userStarringRepos := getUserStarringRepos(user)
	poolRepos := getSubtract(targetRepos, userStarringRepos)

	c.Data["json"] = poolRepos
	c.ServeJSON()
}
