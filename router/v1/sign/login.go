package sign

import (
	"myoo/jwt"
	"myoo/models"
	"myoo/plugins"

	"github.com/gin-gonic/gin"
)

type Login struct{}

func (*Login) Get(router *gin.Context) {

}
func (*Login) Post(router *gin.Context) {
	email := router.PostForm("email")
	pwd := router.PostForm("pwd")
	if !plugins.An_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "邮箱格式错误", "data": nil})
		return
	}
	user := new(models.User)
	if !user.Is_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "该邮箱尚未注册", "data": nil})
		return
	}
	if len(pwd) < 5 || plugins.An_Pwd(pwd) {
		router.JSON(200, gin.H{"success": 403, "message": "账号或密码错误", "data": nil})
		return
	}
	DbValue := user.Query_Email(email)
	if models.BcryotSet(pwd, DbValue.Pwd) {
		if jwt_token := jwt.Encode(DbValue.Id); jwt_token != "" {
			router.JSON(200, gin.H{"success": 200, "message": "登录成功", "data": jwt_token})
			return
		}
	}
	router.JSON(200, gin.H{"success": 403, "message": "账号或密码错误", "data": nil})
	return
}
