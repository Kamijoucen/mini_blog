package service

import "MyBlog/model"

var User = new(UserService)

type UserService struct {
}

func (service *UserService) Save(user *model.User) error {

	return nil
}
