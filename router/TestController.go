package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

type TestController struct {
}

func (t *TestController) CreateTest(request *ghttp.Request) {
	_, err := g.DB().Model("db_test").Data(g.Map{"name": "李思岑", "age": 55}).Insert()
	if err != nil {
		request.Response.Writef("error!")
		return
	}
	request.Response.Writef("success!")
}
