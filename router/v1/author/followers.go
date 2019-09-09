package author

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Followers struct{}

func (this *Followers) Get(router *gin.Context) {
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

func (this *Followers) page(router *gin.Context, id int64) {
	cachePath := "Myoo/Author/" + router.Param("id") + "/Followers/" + router.Param("page")
	template := plugins.Page{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		followers := new(models.Followers)
		DbValue := followers.Count_Followers(id)
		template = plugins.CountPage(DbValue)
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return

}
func (this *Followers) pages(router *gin.Context, id, page int64) {
	cachePath := "Myoo/Author/" + router.Param("id") + "/Followers/" + router.Param("page")

	template := []plugins.User{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		followers := new(models.Followers)
		num := plugins.NumberPage()
		offset := (page - 1) * num

		DbValue := followers.Query_Followers(id, num, offset)
		template = plugins.CheckUsers(DbValue)
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})

}
