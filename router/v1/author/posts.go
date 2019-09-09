package author

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Posts struct{}

func (this *Posts) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	page := plugins.StringtoInt64(router.Param("page"))
	auth := redis.Query_User(id)
	if auth.Id != id {
		router.JSON(200, gin.H{"success": 403, "message": "该用户不存在", "data": nil})
		return
	}

	if page == 0 {
		this.page(router, id)
	} else {
		this.pages(router, id, page)
	}
	return

}

func (this *Posts) page(router *gin.Context, id int64) {
	cachePath := "Myoo/Author/" + router.Param("id") + "/Posts/" + router.Param("page")
	template := plugins.Page{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		posts := new(models.Posts)
		DbValue := posts.Count_User(id)
		template = plugins.CountPage(DbValue)
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return

}
func (this *Posts) pages(router *gin.Context, id, page int64) {
	cachePath := "Myoo/Author/" + router.Param("id") + "/Posts/" + router.Param("page")
	template := []plugins.Posts{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		posts := new(models.Posts)
		num := plugins.NumberPage()
		offset := (page - 1) * num
		DbValue := posts.Query_User(id, num, offset)
		template = plugins.CheckPostsAll(DbValue)
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
