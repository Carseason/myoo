package posts

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

func (this *Posts) Post(router *gin.Context) {
	uid := router.GetInt64("auth")
	if uid == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "请先登录", "data": nil})
		return
	}

	id := plugins.StringtoInt64(router.Param("id"))
	posts := new(models.Posts)
	if id == 0 || !posts.Is(id) {
		router.JSON(200, gin.H{"success": 403, "message": "该文章不存在", "data": nil})
		return
	}

	types := router.PostForm("type")
	switch types {
	case "favorite":
		this.favorite(router, id, uid)
	case "supported":
		this.supported(router, id, uid)
	default:
		router.JSON(200, gin.H{"success": 403, "message": "尚未开放", "data": nil})
	}
	return

}
func (this *Posts) favorite(router *gin.Context, id, uid int64) {
	auth := redis.Query_User(uid)
	if !plugins.Group_Favorite(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组权限不足", "data": nil})
		return
	}
	favorite := new(models.Favorite)
	if favorite.Is(auth.Id, id) {
		favorite.Delete(auth.Id, id)
	} else {
		favorite.Insert(auth.Id, id)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": nil})
	return

}
func (this *Posts) supported(router *gin.Context, id, uid int64) {
	auth := redis.Query_User(uid)
	if !plugins.Group_Supported(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组权限不足", "data": nil})
		return
	}
	supported := new(models.PostsSupported)
	if supported.Is(auth.Id, id) {
		supported.Delete(auth.Id, id)
	} else {
		supported.Insert(auth.Id, id)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": nil})
	return

}
