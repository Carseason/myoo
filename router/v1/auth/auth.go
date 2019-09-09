package auth

import (
	"myoo/plugins"
	"myoo/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type Auth struct{}
type Auther struct {
	Id          int64
	Name        string
	Email       string
	Date        time.Time
	Avatar      string
	Description string
	Group       struct {
		Lv       int64
		NickName string
		//		Medal    []string
	}
}

func (*Auth) Get(router *gin.Context) {
	id := router.GetInt64("auth")
	if id == 0 {
		router.JSON(200, gin.H{"success": 302, "message": "请先登录", "data": nil})
		return
	}
	user := redis.Query_User(id)
	if user.Id != id {
		router.JSON(200, gin.H{"success": 302, "message": "请先登录", "data": nil})
		return
	}

	auth := new(Auther)
	auth.Id = user.Id
	auth.Name = user.Name
	auth.Avatar = plugins.CheckAvatar(user.Avatar)
	auth.Description = plugins.CheckDescription(user.Description)
	auth.Date = user.Date
	auth.Email = user.Email
	group := plugins.CheckGroup(user.Lv)
	auth.Group.Lv = group.Lv
	auth.Group.NickName = group.NickName
	//	auth.Group.Medal = group.Medal
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": auth})
	return
}
