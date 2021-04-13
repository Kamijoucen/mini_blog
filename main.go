package main

import (
	"MyBlog/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	server := g.Server()
	config(server)
	bind(server)
	server.Run()
}

func bind(server *ghttp.Server) {
	server.SetRewrite("/", "/test1.html")
	server.BindObject("/test", router.TestController{})
	server.BindObject("/article", router.ArticleRouter{})
}

func config(server *ghttp.Server) {
	server.SetServerRoot("./static")
	server.SetPort(80)
}
