package models

import (
	"github.com/astaxie/beego/orm"
)

/***********工厂方法**************/
func NewComments() *Comments {
	return new(Comments)
}

/***********查询评论是否存在**************/
func (*Comments) Is(id int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("Comments").Filter("Id", id).Exist()
}

/***********单条评论查询**************/
func (*Comments) Query_One(id int64) Comments {
	o := orm.NewOrm()
	this := Comments{}
	o.QueryTable("Comments").Filter("Id", id).RelatedSel("Author").One(&this)
	return this
}

/***********查询全部评论数量**************/
func (*Comments) Count_All() int64 {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable("Comments").Count(); err == nil {
		return cnt
	}
	return 0
}

/***********新增评论**************/
func (*Comments) Insert(value *Comments) int64 {
	o := orm.NewOrm()
	if id, err := o.Insert(value); err == nil {
		return id
	}
	return 0
}

/***********查询时间段评论数量**************/
func (*Comments) Count_Date(status string) int64 {
	o := orm.NewOrm()
	var maps []orm.Params
	var err error
	var num int64
	switch status {
	case "day":
		num, err = o.Raw("select * from `comments` where to_days(date) = to_days(now());").Values(&maps)
	case "week":
		num, err = o.Raw("SELECT * FROM `comments` where DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(date)").Values(&maps)
	case "month":
		num, err = o.Raw("SELECT * FROM `comments` where DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(date)").Values(&maps)
	case "year":
		num, err = o.Raw("select * from `comments` where YEAR(date)=YEAR(NOW());").Values(&maps)
	default:
		num, err = o.Raw("select * from `comments` where to_days(date) = to_days(now());").Values(&maps)
	}
	if err != nil {
		return 0
	}
	return num
}

/***********条件查询评论数量**************/
func (*Comments) Query_Count(name string, value interface{}) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("Comments").Filter(name, value).Count()
	return cnt
}

/***********文章评论查询**************/
func (*Comments) Count_Posts(id int64) int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Comments").Filter("Posts", id).Count()
	if err != nil {
		return 0
	}
	return cnt
}
func (*Comments) Query_Posts(id, number, page int64) []Comments { //文章ID,查询的数量，翻页
	o := orm.NewOrm()
	this := []Comments{}
	o.QueryTable("Comments").Filter("Posts", id).RelatedSel("Author").Limit(number, page).OrderBy("Id").All(&this)
	return this
}

/***********查询作者的评论**************/
func (*Comments) Count_User(id int64) int64 {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable("Comments").Filter("Author", id).Count(); err == nil {
		return cnt
	}
	return 0
}
func (*Comments) Query_User(id int64, number, page int64) []Comments {
	o := orm.NewOrm()
	all := []Comments{}
	o.QueryTable("Comments").Filter("Author", id).OrderBy("-Id").Limit(number, page).Distinct().RelatedSel("Posts").All(&all)
	return all
}

/***********文章评论删除**************/
func (*Comments) Delete_All(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Comments").Filter("Posts", id).Delete()
	return err == nil
}

/*********************END***************************/
