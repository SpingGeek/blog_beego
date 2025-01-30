// Package controllers @Author Gopher
// @Date 2025/1/30 10:43:00
// @Desc
package controllers

import (
	"blog_beego/models"
	"html/template"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/russross/blackfriday/v2"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) AddGet() {
	c.TplName = "add_blog.html"
}

func (c *BlogController) AddPost() {
	title := c.GetString("title")
	tag := c.GetString("tag")
	content := c.GetString("content")

	post := models.Post{
		Title:   title,
		Content: content,
		Tag:     tag,
	}
	models.Mgr.AddPost(&post)
	c.Redirect("/blog_index", 302)
}

func (c *BlogController) ListPost() {

	posts := models.Mgr.GetAllPost()

	c.Data["posts"] = posts

	c.TplName = "blog_index.html"
}

func (c *BlogController) PostDetail() {
	s := c.GetString("pid")

	pid, _ := strconv.Atoi(s)
	p := models.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.Data["title"] = p.Title
	c.Data["content"] = template.HTML(content)

	c.TplName = "blog_detail.html"

}
