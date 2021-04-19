package service

import (
	"MyBlog/db"
	"MyBlog/model"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

var Article = new(ArticleService)

type ArticleService struct{}

func (receiver *ArticleService) Save(article *model.Article) error {
	article.CreateTime = time.Now()
	article.ModifyTime = time.Now()
	_, err := g.DB().Model(db.ARTICLE).Data(article).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (receiver *ArticleService) ListMain() (gdb.List, error) {
	result, err := g.DB().Model(db.ARTICLE).Fields(
		"id, title, create_time, modify_time").All()
	if err != nil {
		return nil, err
	}
	return result.List(), nil
}
