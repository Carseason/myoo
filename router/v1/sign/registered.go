package sign

import (
	"myoo/jwt"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Registered struct{}

func (*Registered) Get(router *gin.Context) {
	if !redis.GetOptionsRegistered() {
		router.JSON(200, gin.H{"success": 200, "message": "当前注册已关闭", "data": false})
		return
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": true})
	return

}
func (*Registered) Post(router *gin.Context) {
	email := router.PostForm("email")
	name := router.PostForm("name")
	pwd := router.PostForm("pwd")
	pwd2 := router.PostForm("pwd2")
	if !redis.GetOptionsRegistered() {
		router.JSON(200, gin.H{"success": 403, "message": "当前注册已关闭", "data": nil})
		return
	}
	if pwd != pwd2 {
		router.JSON(200, gin.H{"success": 403, "message": "两次输入的密码不一致", "data": nil})
		return
	}
	if len(pwd) < 5 {
		router.JSON(200, gin.H{"success": 403, "message": "请输入至少五位的密码", "data": nil})
		return
	}
	if len(pwd) > 25 {
		router.JSON(200, gin.H{"success": 403, "message": "密码不能太长了", "data": nil})
		return
	}
	if plugins.An_Pwd(pwd) {
		router.JSON(200, gin.H{"success": 403, "message": "密码不能使用 . _ 以外的特殊符号", "data": nil})
		return
	}
	if !plugins.An_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "邮箱格式错误", "data": nil})
		return
	}
	if len(name) < 3 {
		router.JSON(200, gin.H{"success": 403, "message": "用户名太短了", "data": nil})
		return
	}
	if len(name) > 20 {
		router.JSON(200, gin.H{"success": 403, "message": "用户名太长了", "data": nil})
		return
	}
	user := new(models.User)
	if user.Is_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "该邮箱已注册", "data": nil})
		return
	}
	if user.Is_Name(name) {
		router.JSON(200, gin.H{"success": 403, "message": "该用户名已被使用", "data": nil})
		return
	}
	if id := user.Insert(&models.User{
		Email: email,
		Name:  name,
		Pwd:   pwd,
		Admin: false,
		Lv:    redis.GetDefaultLv(),
	}); id > 0 {
		router.JSON(200, gin.H{"success": 200, "message": "注册成功", "data": jwt.Encode(id)})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "注册失败", "data": nil})
	return
}
