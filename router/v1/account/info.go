package account

import (
	"encoding/base64"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Info struct{}

func (this *Info) Post(router *gin.Context) {
	switch router.PostForm("type") {
	case "description":
		this.description(router)
		return
	case "pwd":
		this.pass(router)
		return
	case "avatar":
		this.avatar(router)
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
	return
}

func (this *Info) description(router *gin.Context) {
	uid := router.GetInt64("auth")
	description := router.PostForm("description")
	if len(description) > 100 {
		router.JSON(200, gin.H{"success": 403, "message": "签名长度超出限制", "data": nil})
		return
	}
	value := redis.Query_User(uid)
	if description == value.Description {
		router.JSON(200, gin.H{"success": 200, "message": "没有要更新的", "data": nil})
		return
	}
	user := new(models.User)
	if user.Update_Value(uid, "Description", description) {
		redis.Del_User(uid)
		router.JSON(200, gin.H{"success": 200, "message": "更新成功", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "更新失败", "data": nil})
	return

}
func (this *Info) pass(router *gin.Context) {
	uid := router.GetInt64("auth")
	pwd := router.PostForm("pwd")
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
	user := new(models.User)
	if user.Update_Pwd(uid, pwd) {
		redis.Del_User(uid)
		router.JSON(200, gin.H{"success": 200, "message": "密码更新成功,请重新登录", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "更新失败", "data": nil})
	return
}
func (this *Info) avatar(router *gin.Context) {
	uid := router.GetInt64("auth")
	avatar := router.PostForm("avatar")
	if avatar == "" {
		router.JSON(200, gin.H{"success": 403, "message": "无法更新头像", "data": nil})
		return
	}
	baseSpli := strings.SplitAfter(avatar, ";base64,")
	if len(baseSpli) != 2 {
		router.JSON(200, gin.H{"success": 403, "message": "无法更新头像", "data": nil})
		return
	}
	if len(baseSpli[1]) > 4096000 {
		router.JSON(200, gin.H{"success": 403, "message": "上传的文件占用过大!", "data": nil})
		return
	}

	auth := redis.Query_User(uid)
	if !plugins.Group_UpdateAvatar(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组权限不足", "data": nil})
		return
	}

	imgs, err := base64.StdEncoding.DecodeString(baseSpli[1])
	if err != nil {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
		return
	}

	pathId := strconv.FormatInt(auth.Id, 10)
	pathFile := "src/avatar/" + pathId + ".jpg"
	f, err := os.OpenFile(pathFile, 0666, os.ModePerm)
	if err != nil { //文件不存在，创建文件流
		f, err = os.Create(pathFile)
		if err != nil {
			router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
			return
		}
	}
	defer f.Close()
	f.Write(imgs)
	user := new(models.User)
	user.Update_Avatar(uid, "/src/avatar/"+pathId+".jpg")
	redis.Del_User(uid)
	router.JSON(200, gin.H{"success": 200, "message": "头像已更新", "data": nil})
	return
}
