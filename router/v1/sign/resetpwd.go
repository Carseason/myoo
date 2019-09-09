package sign

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type ResetPwd struct{}
type ResetPwder struct {
	Id         int64
	IssuedTime int64
	OutTime    int64
}

func (*ResetPwd) Post(router *gin.Context) {
	pwd := router.PostForm("pwd")
	pwd2 := router.PostForm("pwd2")
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

	token, err := plugins.Aes_Decode(router.Param("path"))
	if err != nil {
		router.JSON(200, gin.H{"success": 403, "message": "该链接已失效", "data": nil})
		return
	}

	template := ResetPwder{}
	if json.Unmarshal([]byte(token), &template); err != nil || template.Id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "该链接已失效", "data": nil})
		return
	}

	if time.Now().Unix() > template.OutTime {
		router.JSON(200, gin.H{"success": 403, "message": "该链接已失效", "data": nil})
		return
	}

	user := new(models.User)
	if user.Update_Pwd(template.Id, pwd) {
		router.JSON(200, gin.H{"success": 200, "message": "重置密码成功", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "重置密码失败", "data": nil})
	return
}
