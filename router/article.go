package router

import (
	"MyBlog/db"
	"MyBlog/model"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"time"
)

type ArticleRouter struct {
}

func (receiver *ArticleRouter) Save(request *ghttp.Request) {
	article := new(model.Article)
	if err := request.Parse(article); err != nil {
		request.Response.Writef(err.Error())
		return
	}
	article.CreateTime = time.Now()
	article.ModifyTime = time.Now()
	_, err := g.DB().Model(db.ARTICLE).Data(article).Insert()
	if err != nil {
		request.Response.Writef(err.Error())
		return
	}
	_ = request.Response.WriteJson(Success())
}

func (receiver *ArticleRouter) Detail(request *ghttp.Request) {

}

func (receiver *ArticleRouter) List(request *ghttp.Request) {
	result, err := g.DB().Model(db.ARTICLE).Fields(
		"id, title, create_time, modify_time").All()
	if err != nil {
		request.Response.Writef(err.Error())
		return
	}
	_ = request.Response.WriteJson(result.List())
}
