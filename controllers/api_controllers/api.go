package api_controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsluoyz/gitstar/api"
)

type APIController struct {
	beego.Controller
}

func (c *APIController) getSessionUser() string {
	user, ok := c.GetSecureCookie(CookieSecret, CookieKey)
	if !ok {
		return ""
	}

	return user
}

func (c *APIController) setSessionUser(user string) {
	c.SetSecureCookie(CookieSecret, CookieKey, user)
}

func (c *APIController) requireLogin() bool {
	if c.getSessionUser() == "" {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("请先登录"))
		return true
	}

	return false
}

func (c *APIController) requireUser(user string) bool {
	if c.requireLogin() {
		return true
	} else if c.getSessionUser() != user {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("当前登录用户无权限为其他用户执行此操作"))
		return true
	}

	return false
}

func (c *APIController) requireAdmin() bool {
	if c.requireLogin() {
		return true
	} else if !api.IsUserAdmin(c.getSessionUser()) {
		c.Ctx.ResponseWriter.WriteHeader(403)
		c.Ctx.ResponseWriter.Write([]byte("当前登录用户不是管理员，无权限执行此操作"))
		return true
	}

	return false
}
