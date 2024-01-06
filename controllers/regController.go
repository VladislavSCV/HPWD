package controllers

import (
	"github.com/astaxie/beego"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) Get() {
	c.TplName = "reg.html"
}
