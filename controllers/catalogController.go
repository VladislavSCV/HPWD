package controllers

import (
	"github.com/astaxie/beego"
)

type CatalogController struct {
	beego.Controller
}

func (c *CatalogController) Get() {
	c.TplName = "catalog.html"
	c.Data["countVacancy"] = 0
}