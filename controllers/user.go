package controllers

import "github.com/nabetama/Kiz/models"

type User struct {
}

func NewUser() User {
	return User{}
}

func (c User) Get(n int) interface{} {
	repo := models.NewUserRepository()
	user := repo.GetById(n)
	return user
}
