package model

type User struct {
	Id       int
	Name     string `v:"required"`
	ShowName string
	Enable   bool
	Icon     string
}
