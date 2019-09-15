package recommend

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Recommend struct{}

func (*Recommend) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	cachePath := "Myoo/Recommend/" + router.Param("id")
	template := []plugins.Posts{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		posts := new(models.Posts)
		DbValue := posts.Query_Rand([]int64{id}, 12)

		template = plugins.CheckPostsAll(DbValue)
		redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
