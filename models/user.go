package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

func BcryotSet(s, s1 string) bool { //密码比较
	if err := bcrypt.CompareHashAndPassword([]byte(s1), []byte(s)); err == nil { //s原密码,S1为数据库密码
		return true
	}
	return false
}

func BcryotPut(s string) string { //密码生成
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost); err == nil { //原文,加密的等级
		return string(hashedPassword)
	} else {
		return ""
	}

}

/***********工厂方法**************/
func NewUser() *User {
	return new(User)
}

/***********判断Id是否存在**************/
func (*User) Is(id int64) bool {
	o := orm.NewOrm()
	return o.QueryTable("User").Filter("Id", id).Exist()
}

/***********判断邮箱是否存在**************/
func (*User) Is_Email(email string) bool {
	o := orm.NewOrm()
	return o.QueryTable("User").Filter("Email", email).Exist()
}

/***********判断用户名是否存在**************/
func (*User) Is_Name(name string) bool {
	o := orm.NewOrm()
	return o.QueryTable("User").Filter("Name", name).Exist()
}

/***********创建用户**************/
func (*User) Insert(value *User) int64 {
	value.Pwd = BcryotPut(value.Pwd)
	o := orm.NewOrm()
	if id, err := o.Insert(value); err == nil {
		return id
	} else {
		fmt.Println(err)
	}
	return 0
}

/***********查询用户**************/
func (*User) Query_One(value interface{}) *User {
	o := orm.NewOrm()
	this := NewUser()
	switch v := value.(type) {
	case int64:
		o.QueryTable("User").Filter("Id", v).One(this)
	case string:
		o.QueryTable("User").Filter("Name", v).One(this)
	}
	return this
}

/***********邮箱查询用户**************/
func (*User) Query_Email(email string) *User {
	o := orm.NewOrm()
	this := NewUser()
	o.QueryTable("User").Filter("Email", email).One(this)
	return this
}

/***********查询用户数量**************/
func (*User) Count() int64 {
	o := orm.NewOrm()
	if cnt, err := o.QueryTable("User").Count(); err == nil {
		return cnt
	}
	return 0
}

/***********日期条件查询数量**************/
func (*User) Count_Date(status string) int64 {
	o := orm.NewOrm()
	var maps []orm.Params
	var err error
	var num int64
	switch status {
	case "day":
		num, err = o.Raw("select *  from `user` where to_days(date) = to_days(now());").Values(&maps)
	case "week":
		num, err = o.Raw("SELECT * FROM `user` where DATE_SUB(CURDATE(), INTERVAL 7 DAY) <= date(date)").Values(&maps)
	case "month":
		num, err = o.Raw("SELECT * FROM `user` where DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(date)").Values(&maps)
	case "year":
		num, err = o.Raw("select * from `user` where YEAR(date)=YEAR(NOW());").Values(&maps)
	default:
		return 0
	}
	if err != nil {
		return 0
	}
	return num
}

/***********查询30篇文章**************/
func (*User) Query_All(count, offset int64) []User {
	o := orm.NewOrm()
	all := []User{}
	o.QueryTable("User").Limit(count, offset).Distinct().All(&all)
	return all
}

/***********用户更新**************/
func (*User) Update(value *User) bool {
	value.Pwd = BcryotPut(value.Pwd)
	o := orm.NewOrm()
	_, err := o.Update(value)
	return err == nil
}

/***********用户资料更新**************/
func (*User) Update_Value(id int64, name string, value interface{}) bool {
	if name == "pwd" || name == "Pwd" {
		value = BcryotPut(value.(string))
	}
	o := orm.NewOrm()
	_, err := o.QueryTable("User").Filter("Id", id).Update(map[string]interface{}{
		name: value,
	})
	return err == nil
}

/***********用户密码更新**************/
func (*User) Update_Pwd(id int64, pwd string) bool {
	pwd = BcryotPut(pwd)
	o := orm.NewOrm()
	_, err := o.QueryTable("User").Filter("Id", id).Update(map[string]interface{}{
		"Pwd": pwd,
	})
	return err == nil
}

/***********用户头像更新更新**************/
func (*User) Update_Avatar(id int64, avatar string) bool {
	o := orm.NewOrm()
	_, err := o.QueryTable("User").Filter("Id", id).Update(map[string]interface{}{
		"Avatar": avatar,
	})
	return err == nil
}


/***********用户资料更新更新**************/
func (*User) Update_Map(value map[string]interface{}) bool {
	id, ok := value["Id"].(int64)
	if !ok {
		return false
	}
	if value["Pwd"] != nil {
		pwd := value["Pwd"].(string)
		value["Pwd"] = BcryotPut(pwd)
	}
	o := orm.NewOrm()
	if _, err := o.QueryTable("User").Filter("Id", id).Update(value); err != nil {
		return false
	}
	return true
}
