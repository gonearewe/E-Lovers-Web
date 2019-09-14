package models

import "github.com/astaxie/beego/orm"

type administrator struct {
	id         int
	username   string `orm:"unique"`
	emailaddr  string
	password   string
	createtime int64
	status     bool
}

func NewAdministrator(username, emailaddr, password string) *administrator {
	return &administrator{
		username:  username,
		emailaddr: emailaddr,
		password:  password,
		status:    true,
	}
}

func (u *administrator) Insert() (Id int64, err error) {
	return orm.NewOrm().Insert(u)
}

// func (u *administrator) GetByName() (*administrator, error) {
// 	a := new(administrator)

// 	err := orm.NewOrm().QueryTable("administrator").Filter("Username", u.Username).One(a)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return a, nil
// }
// func (u *administrator) Exist() bool {
// 	return orm.NewOrm().QueryTable("administrator").Filter("Username", u.Username).Exist()
// }
