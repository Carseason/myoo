package comments

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"time"

	"github.com/gin-gonic/gin"
)

type Comments struct{}

type Commentser struct {
	Id      int64
	Author  plugins.User
	Content string
	Date    time.Time
}

func (this *Comments) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
		return
	}
	page := plugins.StringtoInt64(router.Param("page"))
	if page == 0 {
		this.page(router, id)
	} else {
		this.pages(router, id, page)
	}
	return

}

func (*Comments) page(router *gin.Context, id int64) {
	cachePath := "Myoo/Comments/" + router.Param("id") + "/" + router.Param("page")
	template := plugins.Page{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		comments := new(models.Comments)
		DbValue := comments.Count_Posts(id)
		template = plugins.CountPage(DbValue)
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (*Comments) pages(router *gin.Context, id, page int64) {
	cachePath := "Myoo/Comments/" + router.Param("id") + "/" + router.Param("page")
	template := []Commentser{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		comments := new(models.Comments)
		num := plugins.NumberPage()
		offset := (page - 1) * num
		DbValue := comments.Query_Posts(id, num, offset)
		for v := range DbValue {
			template = append(template, Commentser{
				Id:      DbValue[v].Id,
				Author:  plugins.CheckUser(*DbValue[v].Author),
				Content: DbValue[v].Content,
				Date:    DbValue[v].Date,
			})
		}
		defer redis.PutStruct(cachePath, &template, 3600)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return

}
