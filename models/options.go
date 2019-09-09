package models

import (
	"github.com/astaxie/beego/orm"
)

func NewOptions() *Options {
	return new(Options)
}
func (*Options) Query(name ...string) map[string]string {
	o := orm.NewOrm()
	value := new(Options)
	valueMap := make(map[string]string)
	for v := range name {
		if o.QueryTable("Options").Filter("Name", name[v]).One(value) == nil {
			valueMap[name[v]] = value.Value
		}
	}
	return valueMap
}

func (*Options) Update(value map[string]string) bool {
	o := orm.NewOrm()
	for v, k := range value {
		if o.QueryTable("Options").Filter("Name", v).Exist() {
			o.QueryTable("Options").Filter("Name", v).Update(map[string]interface{}{
				"Name":  v,
				"Value": k,
			})
		} else {
			this := new(Options)
			this.Name = v
			this.Value = k
			o.Insert(this)
		}
	}
	return true
}
