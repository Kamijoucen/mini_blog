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
	server.BindObject("/", router.TestController{})
}

func config(server *ghttp.Server) {
	//server.SetServerRoot("/home/sftpuser/web_host/MyBlog/static")
	server.SetPort(80)
}
