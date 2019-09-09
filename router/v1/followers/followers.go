/**************************************************
@DCommit
@Description: 关注点赞
@MethodAuthor:  Carseason
@Date: 2019-09-04 21:53:23*
**************************************************/

package followers

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Followers struct{}

func (*Followers) Post(router *gin.Context) {
	uid := router.GetInt64("auth")

	if uid == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "请先登陆", "data": nil})
		return
	}
	id := plugins.StringtoInt64(router.PostForm("id"))
	if uid == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "非法的用户", "data": nil})
		return
	}
	if id == uid {
		router.JSON(200, gin.H{"success": 403, "message": "你不能关注自己!", "data": nil})
		return
	}
	auth := redis.Query_User(uid)
	if !plugins.Group_Followers(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组权限不足", "data": nil})
		return
	}

	gz := new(models.Followers)
	if gz.Is(auth.Id, id) {
		gz.Del(auth.Id, id) //取消关注
		router.JSON(200, gin.H{"success": 200, "message": "已取消关注", "data": nil})
	} else {
		gz.Insert(auth.Id, id) //添加关注
		router.JSON(200, gin.H{"success": 200, "message": "已关注", "data": nil})
	}
	return

}
