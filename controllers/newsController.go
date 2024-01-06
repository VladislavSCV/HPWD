package controllers

import (
	"github.com/astaxie/beego"
)

type NewsController struct {
	beego.Controller
}

func (c *NewsController) News1(){
	c.TplName = "News/newsFirst.html"
}

func (c *NewsController) News2(){
	c.TplName = "News/newsSecond.html"
}

func (c *NewsController) News3(){
	c.TplName = "News/newsThird.html"
}