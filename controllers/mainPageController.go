package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["NewsFirst"] = "Важные новости"
	c.Data["NewsSecond"] = "Возрадуйтесь смертные огурцы растут как на дрожжах"
	c.Data["NewsThird"] = "ДА"
	c.TplName = "index.html"
}
