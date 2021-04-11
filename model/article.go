package model

import "time"

type Article struct {
	Id             int
	Title          string
	Content        string
	CreateUserId   int
	CreateUserName string
	CreateTime     time.Time
	ModifyTime     time.Time
}
