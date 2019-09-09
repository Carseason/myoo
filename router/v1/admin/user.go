package admin

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type User struct{}

func (*User) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	users := new(models.User)
	value := users.Query_One(id)
	value.Pwd = ""
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": value})
	return
}
func (this *User) Post(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		this.insert(router)
	} else if id > 0 {
		this.update(router, id)
	} else {
		router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
	}
	return
}

func (*User) insert(router *gin.Context) {
	name := router.PostForm("name")
	email := router.PostForm("email")
	pwd := router.PostForm("pwd")
	description := router.PostForm("description")
	lv := plugins.StringtoInt64(router.Param("lv"))
	if name == "" || plugins.An_Name(name) {
		router.JSON(200, gin.H{"success": 403, "message": "用户名不能为空或带有非法符号", "data": nil})
		return
	}
	if email == "" || !plugins.An_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "邮箱不能为空或非法邮箱", "data": nil})
		return
	}
	if pwd == "" || !plugins.An_Pwd(pwd) {
		router.JSON(200, gin.H{"success": 403, "message": "密码不能为空或存在 . _外的非法符号", "data": nil})
		return
	}
	user := new(models.User)
	if user.Is_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "该邮箱已被注册", "data": nil})
		return
	}
	if user.Is_Name(name) {
		router.JSON(200, gin.H{"success": 403, "message": "该用户名已被使用", "data": nil})
		return
	}
	if id := user.Insert(&models.User{
		Name:        name,
		Email:       email,
		Pwd:         pwd,
		Admin:       false,
		Lv:          lv,
		Description: description,
	}); id > 0 {
		router.JSON(200, gin.H{"success": 200, "message": "用户创建成功", "data": id})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "用户创建失败", "data": nil})
	return
}

func (*User) update(router *gin.Context, id int64) {
	name := router.PostForm("name")
	email := router.PostForm("email")
	pwd := router.PostForm("pwd")
	description := router.PostForm("description")

	lv := plugins.StringtoInt64(router.PostForm("lv"))
	data := make(map[string]interface{})
	user := new(models.User)
	value := user.Query_One(id)
	if name != value.Name {
		if name == "" || plugins.An_Name(name) {
			router.JSON(200, gin.H{"success": 403, "message": "用户名不能为空或带有非法符号", "data": nil})
			return
		}
		if user.Is_Name(name) {
			router.JSON(200, gin.H{"success": 403, "message": "该用户名已存在", "data": nil})
			return
		}
		data["Name"] = name
	}

	if email != value.Email {
		if email == "" || !plugins.An_Email(email) {
			router.JSON(200, gin.H{"success": 403, "message": "邮箱不能为空或非法邮箱", "data": nil})
			return
		}
		if user.Is_Email(email) {
			router.JSON(200, gin.H{"success": 403, "message": "该邮箱已存在", "data": nil})
			return
		}
		data["Email"] = email
	}
	if pwd != "" {
		if !plugins.An_Pwd(pwd) {
			router.JSON(200, gin.H{"success": 403, "message": "密码存在 . _外的非法符号", "data": nil})
			return
		}
		data["Pwd"] = pwd

	}
	if description != value.Description {
		data["Description"] = description
	}
	if lv != value.Lv {
		data["Lv"] = lv
	}
	data["Id"] = id
	if user.Update_Map(data) {
		redis.Del_User(id)
		router.JSON(200, gin.H{"success": 200, "message": "已更新", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "更新失败", "data": nil})
	return
}
