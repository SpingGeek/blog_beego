// Package routers @Author Gopher
// @Date 2025/1/30 10:42:00
// @Desc
package routers

import (
	"blog_beego/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{})

	beego.Router("/register", &controllers.UserController{}, "get:RegisterGet")
	beego.Router("/register", &controllers.UserController{}, "post:RegisterPost")

	beego.Router("/login", &controllers.UserController{}, "get:LoginGet")
	beego.Router("/login", &controllers.UserController{}, "post:LoginPost")

	beego.Router("/add_blog", &controllers.BlogController{}, "get:AddGet")
	beego.Router("/add_blog", &controllers.BlogController{}, "post:AddPost")
	beego.Router("/blog_list", &controllers.BlogController{}, "get:ListPost")

	beego.Router("/blog_detail", &controllers.BlogController{}, "get:PostDetail")

}
