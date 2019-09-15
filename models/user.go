package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type user struct {
	Id         int
	Username   string `orm:"unique"`
	Emailaddr  string
	Password   string
	Createtime int64
	Status     bool
}

func NewUser(username, emailaddr, password string) *user {
	return &user{
		Username:   username,
		Emailaddr:  emailaddr,
		Createtime: time.Now().Unix(),
		Password:   password,
		Status:     true,
	}
}

func (u *user) GetName() string {
	return u.Username
}
func (u *user) Insert() (id int64, err error) {
	return orm.NewOrm().Insert(u)
}

func UserEmailExist(email string) bool {
	return orm.NewOrm().QueryTable("user").Filter("Emailaddr", email).Exist()
}

// func (u *user) GetByName() (*user, error) {
// 	a := new(user)

// 	err := orm.NewOrm().QueryTable("user").Filter("Username", u.Username).One(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return a, nil
// }
func (u *user) Exist() bool {
	return orm.NewOrm().QueryTable("user").Filter("Username", u.Username).Exist()
}

func (u *user) VerifyPassword() (bool, error) {
	a := new(user)
	err := orm.NewOrm().QueryTable("user").Filter("Username", u.Username).One(a)
	if err != nil {
		return false, err
	}

	return a.Password == u.Password, nil
}
