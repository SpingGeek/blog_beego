// Package controllers @Author Gopher
// @Date 2025/1/30 10:43:00
// @Desc
package controllers

import (
	"blog_beego/models"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.html"
}
func (c *MainController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	models.Mgr.AddUser(&user)
	c.TplName = "adduser.html"

}
