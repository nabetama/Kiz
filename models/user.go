package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	kiz "github.com/nabetama/Kiz/lib"
)

var engine *xorm.Engine

type User struct {
	ID       int    `json:"id" xorm:"'id'"`
	UserName string `json:"name" xrom:"'nickname'"`
}

type UserRepository struct{}

func init() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root@/bookshelf")
	kiz.PanicIf(err)
}

func NewUser(id int, username string) User {
	return User{
		ID:       id,
		UserName: username,
	}
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (m UserRepository) GetById(id int) *User {
	var user = User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}
