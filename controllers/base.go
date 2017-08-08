package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) GetSessionUser() string {
	user := c.GetSession("username")
	if user == nil {
		return ""
	}

	return user.(string)
}

func (c *BaseController) SetSessionUser(user string) {
	c.SetSession("username", user)
}

func (c *BaseController) RequireLogin() bool {
	if c.GetSessionUser() == "" {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("请先登录"))
		return true
	}

	return false
}

func (c *BaseController) RequireUser(user string) bool {
	if c.RequireLogin() {
		return true
	} else if c.GetSessionUser() != user {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("当前登录用户无权限为其他用户执行此操作"))
		return true
	}

	return false
}

func (c *BaseController) RequireAdmin() bool {
	if c.RequireLogin() {
		return true
	} else if !api.IsUserAdmin(c.GetSessionUser()) {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("当前登录用户不是管理员，无权限执行此操作"))
		return true
	}

	return false
}
