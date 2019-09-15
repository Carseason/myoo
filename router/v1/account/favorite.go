package account

import (
	"myoo/models"
	"myoo/plugins"

	"github.com/gin-gonic/gin"
)

type Favorite struct{}

func (this *Favorite) Get(router *gin.Context) {
	page := plugins.StringtoInt64(router.Param("page"))
	if page == 0 {
		this.page(router)
	} else {
		this.pages(router, page)
	}
	return
}

func (*Favorite) page(router *gin.Context) {
	uid := router.GetInt64("auth")
	favorite := new(models.Favorite)
	DbValue := favorite.Count_User(uid)
	template := plugins.CountPage(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func (*Favorite) pages(router *gin.Context, page int64) {
	uid := router.GetInt64("auth")
	favorite := new(models.Favorite)
	num := plugins.NumberPage()
	offset := (page - 1) * num
	DbValue := favorite.Query_User(uid, num, offset)
	template := plugins.CheckPostsAll(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
}
