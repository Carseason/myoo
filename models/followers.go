package models

import "github.com/astaxie/beego/orm"

/***********粉丝关注**************/
/***********工厂方法**************/
func NewFollowers() *Followers {
	return new(Followers)
}

/***********查询用户粉丝数量**************/
func (*Followers) Count_Fans(id int64) int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Followers").Filter("Followers", id).Count()
	if err != nil {
		return 0
	}
	return cnt
}

/***********查询用户关注数量**************/
func (*Followers) Count_Followers(id int64) int64 {
	o := orm.NewOrm()
	cnt, err := o.QueryTable("Followers").Filter("Fans", id).Count()
	if err != nil {
		return 0
	}
	return cnt
}

/***********查询是否已关注**************/
func (*Followers) Is(id, uid int64) bool { //接收当前用户ID，目标ID
	o := orm.NewOrm()
	return o.QueryTable("Followers").Filter("Fans", id).Filter("Followers", uid).Exist()
}

/***********取消关注**************/
func (*Followers) Del(id, uid int64) bool { //接收当前用户ID，目标ID
	o := orm.NewOrm()
	_, err := o.QueryTable("Followers").Filter("Fans", id).Filter("Followers", uid).Delete()
	return err == nil
}

/***********添加关注**************/
func (*Followers) Insert(id, uid int64) bool { //接收当前用户ID，目标ID
	followers := User{
		Id: uid,
	}
	fans := User{
		Id: id,
	}
	this := Followers{
		Followers: &followers,
		Fans:      &fans,
	}
	o := orm.NewOrm()
	_, err := o.Insert(&this)
	return err == nil
}

/***********查询关注的用户**************/
func (*Followers) Query_Followers(id, number, page int64) []User {
	o := orm.NewOrm()
	this := []Followers{}
	o.QueryTable("Followers").Filter("Fans", id).Limit(number, page).RelatedSel("Followers").All(&this, "Followers")
	value := []User{}
	for v := range this {
		value = append(value, *this[v].Followers)
	}
	return value
}

/**************查询用户的粉丝*********************/
func (*Followers) Query_Fans(id, number, page int64) []User {
	o := orm.NewOrm()
	this := []Followers{}
	o.QueryTable("Followers").Filter("Followers", id).Limit(number, page).RelatedSel("Fans").All(&this, "Fans")
	value := []User{}
	for v := range this {
		value = append(value, *this[v].Fans)
	}
	return value
}
