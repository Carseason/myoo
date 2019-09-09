package models

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego/orm"
)

/***********工厂方法**************/
func NewPosts() *Posts {
	return new(Posts)
}

/***********查询所有文章数量**************/
func (*Posts) Count() int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Posts").Count()
	if err != nil {
		return 0
	}
	return cnt
}

/***********查询30篇文章**************/
func (*Posts) Query_All(count, offset int64) []Posts {
	o := orm.NewOrm()
	var all []Posts
	o.QueryTable("Posts").OrderBy("-Id").Limit(count, offset).Distinct().RelatedSel("Author", "Category").All(&all)
	return all
}

/***********查询文章是否存在**************/
func (*Posts) Is(id int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("Posts").Filter("Id", id).Exist()
}

/***********查询文章**************/
func (*Posts) Query(id int64) *Posts {
	o := orm.NewOrm()
	this := new(Posts)
	o.QueryTable("Posts").Filter("Id", id).RelatedSel("Category", "Author").One(this)
	return this
}

func (*Posts) Query_One(id int64) *Posts {
	o := orm.NewOrm()
	this := new(Posts)
	o.QueryTable("Posts").Filter("Id", id).Filter("Status", "publish").RelatedSel("Category", "Author").One(this)
	return this
}

/***********更新浏览量**************/
func (*Posts) Update_Views(id, views int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Posts").Filter("Id", id).Update(orm.Params{
		"Views": views,
	})
	return err == nil
}
func (*Posts) Increase_Views(id int64) {
	o := orm.NewOrm()
	o.Raw("UPDATE posts SET views = views +1  WHERE id = ?;", id).Exec()
}

/***********更新文章**************/
func (*Posts) Update(value *Posts) bool {
	o := orm.NewOrm()
	_, err := o.Update(value, "Title", "Content", "Type", "Thumbnail", "Category", "Status")
	return err == nil
}

/***********插入文章**************/
func (*Posts) Insert(value *Posts) int64 {
	o := orm.NewOrm()
	if id, err := o.Insert(value); err == nil {
		return id
	}
	return 0
}

/****************Account*****************/
/****************查询用户中心文章数量*****************/
func (*Posts) Count_Account(id int64) int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Posts").Filter("Author", id).Count()
	if err != nil {
		return 0
	}
	return cnt
}

/**********查询用户中心文章*************/
func (*Posts) Query_Account(id int64, number, page int64) []Posts {
	o := orm.NewOrm()
	all := []Posts{}
	o.QueryTable("Posts").Filter("Author", id).OrderBy("-Id").Limit(number, page).Distinct().RelatedSel("Category").All(&all)
	return all
}

/****************Author*****************/
/**********查询用户文章数量*************/
func (*Posts) Count_User(id int64) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("Posts").Filter("Status", "publish").Filter("Author", id).Count()
	return cnt
}

/**********查询用户文章*************/
func (*Posts) Query_User(id int64, number, page int64) []Posts {
	o := orm.NewOrm()
	all := []Posts{}
	o.QueryTable("Posts").Filter("Status", "publish").Filter("Author", id).OrderBy("-Id").Limit(number, page).Distinct().RelatedSel("Category", "Author").All(&all)
	return all
}

/***********日期条件查询数量**************/
func (*Posts) Count_Date(status string) int64 {
	o := orm.NewOrm()
	var maps []orm.Params
	var err error
	var num int64
	switch status {
	case "day":
		num, err = o.Raw("select * from `posts` where to_days(date) = to_days(now());").Values(&maps)
	case "week":
		num, err = o.Raw("SELECT * FROM `posts` where DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(date)").Values(&maps)
	case "month":
		num, err = o.Raw("SELECT * FROM `posts` where DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(date)").Values(&maps)
	case "year":
		num, err = o.Raw("select * from `posts` where YEAR(date)=YEAR(NOW());").Values(&maps)
	default:
		num, err = o.Raw("select * from `posts` where to_days(date) = to_days(now());").Values(&maps)
	}
	if err != nil {
		return 0
	}
	return num
}

/***********分类文章数量查询**************/
func (*Posts) Count_Categor(id int64) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable("Posts").Filter("Status", "publish").Filter("Category", id).Count()
	return cnt
}
func (*Posts) Query_Category(id, Number, Offset int64, Orderby string) []Posts {
	all := []Posts{}
	o := orm.NewOrm()
	o.QueryTable("Posts").Filter("Status", "publish").Filter("Category", id).RelatedSel("Category", "Author").Limit(Number, Offset).Distinct().OrderBy(Orderby).All(&all)
	return all
}

/**********模块查询文章**************/
/*接收分类数组,数量,偏移个数,查询排序方式*/
func (*Posts) Query_Box(id []int64, Number, Offset int, Orderby string) []Posts {
	all := []Posts{}
	o := orm.NewOrm()
	o.QueryTable("Posts").Filter("Status", "publish").Filter("Category__in", id).RelatedSel("Category", "Author").Limit(Number, Offset).Distinct().OrderBy(Orderby).All(&all)
	return all
}

/**********随机查询文章**************/
func (*Posts) Query_Rand(id []int64, Number int) []Posts {
	var data string
	var all []Posts
	if len(id) > 0 {
		if jsonData, err := json.Marshal(id); err != nil {
			return all
		} else {
			data = string(jsonData)
			data = strings.Replace(data, "[", "(", -1)
			data = strings.Replace(data, "]", ")", -1)
		}
	} else {
		data = "()"
	}
	o := orm.NewOrm()
	o.Raw("SELECT  * FROM posts WHERE  status = 'publish' and category_id IN "+data+" order by rand() LIMIT ?", Number).QueryRows(&all)
	for v := range all {
		o.LoadRelated(&all[v], "Author")
		o.LoadRelated(&all[v], "Category")
	}

	return all
}

/***********模块文章处理**************/
func (this *Posts) Query_Filter(id []int64, Number, Orderby, Offset int) []Posts {

	all := []Posts{}
	switch Orderby { //排序方式
	case 1: //随机
		all = this.Query_Rand(id, Number)
	case 2: //浏览
		all = this.Query_Box(id, Number, 0, "-Views")
	default:
		all = this.Query_Box(id, Number, 0, "-Id")
	}
	return all
}

/***********文章删除**************/
func (this *Posts) Delete(id int64) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("Posts").Filter("Id", id).Delete()
	return err == nil
}
