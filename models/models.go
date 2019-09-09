package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //驱动
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Options struct {
	Id    int64
	Name  string `orm:"unique;index"`
	Value string `orm:"type(text)"`
}
type Category struct { //分类主键
	Id   int64  //分类ID
	Name string `orm:"unique;"` //分类名称
}
type User struct { //用户表
	Id          int64
	Email       string    `orm:"unique;index"` //邮箱
	Name        string    `orm:"unique;index"` //名字
	Pwd         string    //密码
	Date        time.Time `orm:"type(datetime);auto_now"` //注册时间
	Avatar      string    `orm:"type(text)"`
	Admin       bool      //是否管理员
	Lv          int64     //等级
	Description string    `orm:"size(100)"` //签名
}
type Favorite struct { //作者文章收藏
	Id     int64
	Author *User     `orm:"rel(fk);index"`
	Posts  *Posts    `orm:"rel(fk);index"`
	Date   time.Time `orm:"type(datetime);auto_now"` //发布时间
}
type Followers struct { //关注表
	Id        int64
	Followers *User `orm:"rel(fk);index"` //关注
	Fans      *User `orm:"rel(fk);index"` //粉丝
}
type Posts struct {
	Id        int64
	Date      time.Time `orm:"type(datetime);auto_now"` //发布时间
	Title     string    `orm:"type(text)"`              //文章标题
	Content   string    `orm:"type(text)"`              //文章内容
	Type      string    `orm:"index"`                   //文章类型
	Thumbnail string    //文章略缩图
	Status    string    `orm:"index"`         //文章状态，如已发布publish，待审核pending
	Author    *User     `orm:"rel(fk);index"` //作者
	Category  *Category `orm:"rel(fk);index"` //分类
	Views     int64     `orm:"index"`
}
type PostsMeta struct {
	Id      int64
	PostsId int64  `orm:"index"`
	Name    string `orm:"index"`
	Value   string `orm:"index;type(text)"`
}

type PostsSupported struct { //点赞用户
	Id     int64
	Author *User  `orm:"rel(fk);index"`
	Posts  *Posts `orm:"rel(fk);index"`
}

type Comments struct {
	Id      int64
	Posts   *Posts    `orm:"rel(fk);index"`
	Author  *User     `orm:"rel(fk);index"`
	Date    time.Time `orm:"type(datetime);auto_now"` //发布时间
	Content string    `orm:"type(text)"`              //内容
}

const (
	Path = "root:@/debugseason?charset=utf8&loc=Local"
	//root:@/one?charset=utf8
)

func RegisterDB(addr string) {
	//注册 model
	orm.RegisterModel(
		new(Options),
		new(Category),

		new(User),
		new(Followers),
		new(Favorite),

		new(Posts),
		new(PostsMeta),
		new(PostsSupported),

		new(Comments),
	)

	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	orm.RegisterDataBase("default", "mysql", addr) //数据库名字,数据库类型,数据库路径地址,账号:密码@/localhost
	orm.Debug = false                              //日志信息
	orm.RunSyncdb("default", false, false)
}
