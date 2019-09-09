package models

import "github.com/astaxie/beego/orm"

func NewCategory() *Category {
	return new(Category)
}

/***********查询分类是否存在**************/
func (*Category) Is(value interface{}) bool {
	o := orm.NewOrm()
	switch v := value.(type) {
	case int64:
		return o.QueryTable("Category").Filter("Id", v).Exist()
	case string:
		return o.QueryTable("Category").Filter("Name", v).Exist()
	}
	return false
}

/***********新建分类**************/
func (*Category) Insert(name string) int64 {
	o := orm.NewOrm()
	value := new(Category)
	value.Name = name
	if id, err := o.Insert(value); err == nil {
		return id
	}
	return 0
}

/***********查询所有分类**************/
func (*Category) QueryAll() []Category {
	o := orm.NewOrm()
	all := []Category{}
	o.QueryTable("Category").Limit(-1).All(&all)
	return all
}

/***********查询分类**************/
func (*Category) Query(value interface{}) *Category {
	o := orm.NewOrm()
	this := new(Category)
	switch v := value.(type) {
	case int64:
		o.QueryTable("Category").Filter("Id", v).One(this)
	case string:
		o.QueryTable("Category").Filter("Name", v).One(this)
	}
	return this
}

/***********更新分类**************/
func (*Category) Update(id int64, name string) bool {
	o := orm.NewOrm()
	value := new(Category)
	value.Id = id
	value.Name = name
	_, err := o.Update(value)
	return err == nil
}

/***********分类数量**************/
func (*Category) Count() int64 {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable("Category").Count(); err == nil {
		return cnt
	}
	return 0
}

/***********删除分类**************/
func (*Category) Delete(value interface{}) bool {
	o := orm.NewOrm()
	switch v := value.(type) {
	case string:
		if _, err := o.QueryTable("Category").Filter("Name", v).Delete(); err == nil {
			return true
		}
	case int64:
		if _, err := o.QueryTable("Category").Filter("Id", v).Delete(); err == nil {
			return true
		}
	}
	return false
}

/***********查询分类下是否有文章**************/
func (*Category) QueryPosts(id int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("Posts").Filter("Category", id).Exist()
}
