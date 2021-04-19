package router

import "github.com/gogf/gf/net/ghttp"

type AuthRouter struct {
}

func (receiver *AuthRouter) Login(request *ghttp.Request) {

	if true {
		_ = request.Session.Set("user", nil)
	}
}
