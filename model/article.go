package model

import "time"

type Article struct {
	Id             int
	Title          string `v:"required"`
	Content        string `v:"required"`
	CreateUserId   int
	CreateUserName string
	CreateTime     time.Time
	ModifyTime     time.Time
}
