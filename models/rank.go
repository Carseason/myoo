package models

import (
	"github.com/astaxie/beego/orm"
)

func Query_Rank(value string) []Posts {
	o := orm.NewOrm()
	all := []Posts{}
	o.QueryTable("Posts").Filter("Status", "publish").OrderBy(value).Limit(100).Distinct().All(&all)
	return all
}

func Query_Rank_Posts(id []int64) []Posts {
	t := len(id)
	posts := make([]Posts, t)
	o := orm.NewOrm()
	for i := 0; i < t; i++ {
		o.QueryTable("Posts").Filter("Id", id[i]).One(&posts[i])
	}
	return posts
}

func Query_Rank_Support() []Posts {
	o := orm.NewOrm()
	var id []int64
	o.Raw("select posts_id from posts_supported group by posts_id order by count(posts_id) DESC LIMIT 100;").QueryRows(&id)
	return Query_Rank_Posts(id)
}

func Query_Rank_Comments() []Posts {
	o := orm.NewOrm()
	var id []int64
	o.Raw("select posts_id from comments group by posts_id order by count(posts_id) DESC LIMIT 100;").QueryRows(&id)
	return Query_Rank_Posts(id)
}
