package models

import (
	"github.com/astaxie/beego/orm"
)

type user struct {
	id         int
	username   string `orm:"unique"`
	emailaddr  string
	password   string
	createtime int64
	status     bool
}

func NewUser(username, emailaddr, password string) *user {
	return &user{
		username:  username,
		emailaddr: emailaddr,
		password:  password,
		status:    true,
	}
}

func (u *user) Insert() (Id int64, err error) {
	return orm.NewOrm().Insert(u)
}

// func (u *user) GetByName() (*user, error) {
// 	a := new(user)

// 	err := orm.NewOrm().QueryTable("user").Filter("Username", u.Username).One(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return a, nil
// }
// func (u *user) Exist() bool {
// 	return orm.NewOrm().QueryTable("user").Filter("Username", u.Username).Exist()
// }
