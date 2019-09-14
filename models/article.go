package models

import (
	"strconv"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id       int
	Title    string `orm:"unique"`
	Subtitle string `orm:"null"`
	Author   string
	//	Tags       [5]string
	Text       string `orm:"-"`
	Createtime int64
	Status     bool
}
type Pagination struct {
	HasPrevious, HasNext   bool
	PageIndex              int
	PreviousLink, NextLink string
}

func (u *Article) Insert() (Id int64, err error) {
	return orm.NewOrm().Insert(u)
}
func (u *Article) GetById() (*Article, error) {
	a := new(Article)

	err := orm.NewOrm().QueryTable("article").Filter("Id", u.Id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}
func (u *Article) Exist() bool {
	return orm.NewOrm().QueryTable("article").Filter("Title", u.Title).Exist()
}
func (u *Article) Delete() error {
	o := orm.NewOrm()
	_, err := o.Delete(u)
	return err
}
//NOTICE : I originally set the type of "articles" as "*[]*Article" ,which is also accepted
//by the orm function ,HOWEVER,error about reflect occurs, so I change the type to "[]*Article",
//and then everything works just fine
func ListArticleByPage(numPerPage, pageIndex int64) ([]*Article, error) {
	var articles []*Article
	if _, err := orm.NewOrm().QueryTable("article").OrderBy("-Createtime").
		Limit(numPerPage, (pageIndex-1)*numPerPage).All(&articles); err == nil {
		return articles, nil
	} else {
		return nil, err
	}

}
//Get the total number of articles in the database
func ArticleNum() (int64, error) {
	return orm.NewOrm().QueryTable("article").Count()
}
//Alright,I admit,it's NOT very good to write a function containing specific url in "models"
func (p *Pagination) Set(pageIndex,numPerPage,pagenum int64) {
	p.PageIndex = int(pageIndex)

	p.HasPrevious, p.HasNext = true, true
	if pageIndex == 1 {
		p.HasPrevious = false
	}
	if pageIndex*numPerPage >= pagenum {
		p.HasNext = false
	}

	if p.HasPrevious{
		p.PreviousLink="article/view?page_index="+strconv.Itoa(p.PageIndex-1)
	}else{
		p.PreviousLink="#"
	}

	if p.HasNext{
		p.NextLink="article/view?page_index="+strconv.Itoa(p.PageIndex+1)
	}else{
		p.NextLink="#"
	}
	
}
