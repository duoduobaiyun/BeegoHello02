package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "baidu.me"
	c.Data["Email"] = "liulianwanfan.com"
	c.TplName = "index.tpl"
}
