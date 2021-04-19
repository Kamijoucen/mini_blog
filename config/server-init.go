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

	server.BindHandler("/", func(request *ghttp.Request) {
		_ = request.Response.WriteTpl("index.html")
	})

	server.BindObject("/auth", router.AuthRouter{})
	server.BindObject("/article", router.ArticleRouter{})
	server.BindObject("/user", router.UserRouter{})
}

func config(server *ghttp.Server) {

	_ = g.View().SetPath("./template")
	g.View().SetDelimiters("${", "}")
	g.View().SetAutoEncode(true)

	server.SetServerRoot("./static")
	server.SetPort(80)
}
