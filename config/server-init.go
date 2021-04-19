package config

import (
	"MyBlog/router"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func init() {
	server := g.Server()

	config(server)

	bind(server)
}

func bind(server *ghttp.Server) {
	server.SetRewrite("/", "/index.html")
	server.BindObject("/test", router.TestController{})
	server.BindObject("/article", router.ArticleRouter{})
}

func config(server *ghttp.Server) {
	server.SetServerRoot("./static")
	server.SetPort(80)
}
