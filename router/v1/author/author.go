package author

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type Author struct{}
type Authorer struct {
	Id          int64
	Name        string
	Avatar      string
	Description string
	Date        time.Time
	Group       struct {
		Lv       int64
		NickName string
		//		Medal    []string
	}

	Posts struct {
		Count int64
	}

	Fans struct {
		Count int64
	}
	Followers struct {
		Count  int64
		Status bool
	}
	Comments struct {
		Count int64
	}
}

func (*Author) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
		return
	}

	cachePath := "Myoo/Author/" + router.Param("id")

	template := Authorer{}
	gz := new(models.Followers)
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		value := redis.Query_User(id)
		if value.Id != id {
			router.JSON(200, gin.H{"success": 403, "message": "该用户不存在", "data": nil})
			return
		}
		template.Id = value.Id
		template.Name = value.Name
		template.Date = value.Date
		template.Avatar = plugins.CheckAvatar(value.Avatar)
		template.Description = plugins.CheckDescription(value.Description)
		/***********Group**************/
		group := plugins.CheckGroup(value.Lv)
		template.Group.Lv = group.Lv
		template.Group.NickName = group.NickName
		//		template.Group.Medal = group.Medal

		/***********Posts**************/
		posts := new(models.Posts)
		template.Posts.Count = posts.Count_User(id)

		/***********Fans**************/
		template.Fans.Count = gz.Count_Fans(id)
		/***********Comments**************/
		template.Comments.Count = models.NewComments().Count_User(id)
		/***********Followers**************/
		template.Followers.Count = gz.Count_Followers(id)

		redis.PutStruct(cachePath, &template, 7200)
	}

	if uid := router.GetInt64("auth"); uid > 0 && id != uid {
		template.Followers.Status = gz.Is(uid, id) //判断是否已关注
	}

	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
