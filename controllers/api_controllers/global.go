package api_controllers

import (
	"encoding/base64"

	"github.com/hsluoyz/gitstar/api"
)

// Global API
type GlobalController struct {
	APIController
}

// @Title UpdateStarringRepos
// @Description update all the repos starred by each user into GitStar cache
// @Success 200 {object} controllers.api_controller.Response The Response object
// @router /starring-repos/update [post]
func (c *GlobalController) UpdateStarringRepos() {
	if c.requireLogin() {
		return
	}

	var resp Response

	affected := api.UpdateStarringRepos()

	if affected {
		resp = Response{Code: 200, Msg: "ok", Data: ""}
	} else {
		resp = Response{Code: 200, Msg: "not affected", Data: ""}
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// @Title GetRecommend
// @Description Get the recommend repos for each user
// @Success 200 {object} []api.Entry2 The list of Entry2 objects
// @router /recommend [get]
func (c *GlobalController) GetRecommend() {
	if c.requireLogin() {
		return
	}

	c.Data["json"] = api.GetRecommend()
	c.ServeJSON()
}

// @Title GetOwe
// @Description Get the details that each user owes another user stars
// @Success 200 {object} []*api.UserTargetStatus The list of UserTargetStatus objects
// @router /owe [get]
func (c *GlobalController) GetOwe() {
	if c.requireLogin() {
		return
	}

	c.Data["json"] = api.GetOwe()
	c.ServeJSON()
}

// @Title GetSystemMessages
// @Description Get the system messages from admin
// @Success 200 {object} []api.Message The list of Message objects, text is base64-encoded
// @router /messages [get]
func (c *GlobalController) GetSystemMessages() {
	if c.requireLogin() {
		return
	}

	messages := api.GetSystemMessages()
	for i := range messages {
		messages[i].Text = base64.StdEncoding.EncodeToString([]byte(messages[i].Text))
	}

	c.Data["json"] = messages
	c.ServeJSON()
}
