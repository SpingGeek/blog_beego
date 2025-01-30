// Package controllers @Author Gopher
// @Date 2025/1/30 10:44:00
// @Desc
package controllers

import (
	"blog_beego/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) RegisterGet() {
	c.TplName = "register.html"
}

func (c *UserController) LoginGet() {
	c.TplName = "login.html"

}

func (c *UserController) LoginPost() {
	username := c.GetString("username")
	password := c.GetString("password")
	u := models.Mgr.Login(username)

	fmt.Printf("u: %v\n", u)

	if u.Username == "" {
		c.Data["error"] = "用户名不存在！"
		c.TplName = "login.html"
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			c.Data["error"] = "密码错误"
			c.TplName = "login.html"
			fmt.Println("密码错误")
		} else {
			fmt.Println("登录成功")
			c.Redirect("/", 302)
		}
	}
}

func (c *UserController) RegisterPost() {
	username := c.GetString("username")
	password := c.GetString("password")
	user := models.User{
		Username: username,
		Password: password,
	}

	models.Mgr.AddUser(&user)

	c.Redirect("/", 302)
}
