package admin

import (
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Cache struct{}

func (*Cache) Post(router *gin.Context) {
	defer redis.DelAll("Myoo/*")
	router.JSON(200, gin.H{"success": 200, "message": "已清空缓存", "data": nil})
	return
}
