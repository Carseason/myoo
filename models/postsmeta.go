package models

import (
	"github.com/astaxie/beego/orm"
)

func NewPostsMeta() *PostsMeta {
	return new(PostsMeta)
}

/******************检查文章媒体数据****************/
func (*PostsMeta) Is(uid int64, name string) bool {
	o := orm.NewOrm()
	return o.QueryTable("PostsMeta").Filter("PostsId", uid).Filter("Name", name).Exist()
}

/******************插入文章媒体数据****************/
func (*PostsMeta) Insert(value *PostsMeta) int64 {
	o := orm.NewOrm()
	if id, err := o.Insert(value); err == nil {
		return id
	}
	return 0
}

/******************更新文章媒体数据****************/
func (*PostsMeta) Update(uid int64, name, value string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("PostsMeta").Filter("PostsId", uid).Filter("Name", name).Update(orm.Params{
		"Value": value,
	})
	return err == nil
}

/******************删除文章媒体数据****************/
func (*PostsMeta) Delete(uid int64, name string) bool {
	o := orm.NewOrm()
	o.QueryTable("PostsMeta").Filter("PostsId", uid).Filter("Name", name).Delete()
	return true
}

/******************删除文章所有媒体数据****************/
func (*PostsMeta) Delete_All(uid int64) bool {
	o := orm.NewOrm()
	o.QueryTable("PostsMeta").Filter("PostsId", uid).Delete()
	return true
}

/***********查询文章媒体数据**************/
func (*PostsMeta) Query(id int64, name string) *PostsMeta {
	o := orm.NewOrm()
	value := new(PostsMeta)
	o.QueryTable("PostsMeta").Filter("PostsId", id).Filter("Name", name).One(value)
	return value
}

/***********查询文章媒体数据**************/
func (*PostsMeta) Query_All(id int64) []PostsMeta {
	o := orm.NewOrm()
	value := []PostsMeta{}
	o.QueryTable("PostsMeta").Filter("PostsId", id).One(&value)
	return value
}
