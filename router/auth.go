package router

import "github.com/gogf/gf/net/ghttp"

type AuthRouter struct {
}

func (receiver AuthRouter) Login(request *ghttp.Request) {
	// todo check pwd

	if true {
		request.Session.Set("user", nil)
	}
}
