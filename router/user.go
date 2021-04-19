package router

import (
	"MyBlog/model"
	"MyBlog/service"
	"github.com/gogf/gf/net/ghttp"
)

type UserRouter struct{}

func (userRouter *UserRouter) Save(request *ghttp.Request) {
	user := new(model.User)

	if err := request.Parse(user); err != nil {
		request.Response.Writef(err.Error())
		return
	}

	if err := service.User.Save(user); err != nil {
		request.Response.Writef(err.Error())
		return
	}

	_ = request.Response.WriteJson(model.Success())
}
