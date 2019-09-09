package models

import (
	"github.com/astaxie/beego/orm"
)

/********************************
			文章收藏操作
********************************/
func NewFavorite() *Favorite {
	return new(Favorite)
}

/***************判断收藏*****************/
func (*Favorite) Is(id, postId int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("Favorite").Filter("Author", id).Filter("Posts", postId).Exist()
}

/***************新增收藏*****************/
func (*Favorite) Insert(id, postsId int64) bool { //收藏
	author := User{
		Id: id,
	}
	posts := Posts{
		Id: postsId,
	}
	this := Favorite{
		Author: &author,
		Posts:  &posts,
	}
	o := orm.NewOrm()
	_, err := o.Insert(&this)
	return err == nil
}

/***********取消收藏**************/
func (*Favorite) Delete(id, postId int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Favorite").Filter("Author", id).Filter("Posts", postId).Delete()
	return err == nil
}

/***********删除文章收藏**************/
func (*Favorite) Delete_All(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Favorite").Filter("Posts", id).Delete()
	return err == nil
}

/****************查询用户收藏********************/
func (*Favorite) Count_User(id int64) int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Favorite").Filter("Author", id).Count()
	if err != nil {
		return 0
	}
	return cnt
}
func (*Favorite) Query_User(id int64, number, page int64) []Posts {
	o := orm.NewOrm()
	all := []Favorite{}
	o.QueryTable("Favorite").Filter("Author", id).OrderBy("-Id").Limit(number, page).Distinct().RelatedSel("Posts").All(&all)
	dbValue := []Posts{}
	for v := range all {
		dbValue = append(dbValue, *all[v].Posts)
	}
	return dbValue
}
