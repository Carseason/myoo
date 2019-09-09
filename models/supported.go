package models

import (
	"github.com/astaxie/beego/orm"
)

/********************************
			文章点赞操作
********************************/
func NewSupported() *PostsSupported {
	return new(PostsSupported)
}

/***************点赞查询*****************/
func (*PostsSupported) Is(id, postId int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("PostsSupported").Filter("Author", id).Filter("Posts", postId).Exist()
}

/***************取消点赞*****************/
func (*PostsSupported) Delete(id, postsId int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("PostsSupported").Filter("Author", id).Filter("Posts", postsId).Delete()
	return err == nil
}

/***************删除文章点赞*****************/
func (*PostsSupported) Delete_All(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("PostsSupported").Filter("Posts", id).Delete()
	return err == nil
}

/***************增加点赞*****************/
func (*PostsSupported) Insert(id, postsId int64) bool {
	author := User{
		Id: id,
	}
	posts := Posts{
		Id: postsId,
	}
	this := PostsSupported{
		Author: &author,
		Posts:  &posts,
	}
	o := orm.NewOrm()
	if _, err := o.Insert(&this); err == nil {
		return true
	}

	return false
}

/***************点赞数量*****************/
func (*PostsSupported) Count(id int64) int64 {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable("PostsSupported").Filter("Posts", id).Count(); err == nil {
		return cnt
	}
	return 0
}
