package controllers

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Get() {
	c.TplName = "auth.html"
}