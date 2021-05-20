package controllers

import (
    "github.com/beego/beego/v2/server/web"
)

type MainController struct {
    web.Controller
}

func (this *MainController) ShowHomePage() {
    this.Ctx.WriteString("Home Page")
}

func (this *MainController) Test() {
    this.TplName = "test.tpl"
}
