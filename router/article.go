package router

import (
	"MyBlog/model"
	"MyBlog/service"
	"github.com/gogf/gf/net/ghttp"
)

type ArticleRouter struct {
}

func (receiver *ArticleRouter) Save(request *ghttp.Request) {
	article := new(model.Article)
	if err := request.Parse(article); err != nil {
		request.Response.Writef(err.Error())
		return
	}
	if err := service.Article.Save(article); err != nil {
		request.Response.Writef(err.Error())
		return
	}
	_ = request.Response.WriteJson(model.Success())
}

func (receiver *ArticleRouter) ListMain(request *ghttp.Request) {

	list, err := service.Article.ListMain()

	if err != nil {
		request.Response.Writef(err.Error())
		return
	}
	_ = request.Response.WriteJson(list)
}
