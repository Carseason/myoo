package account

import (
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type Posts struct{}
type Postser struct {
	Id        int64
	Date      time.Time
	Title     string
	Type      string
	Thumbnail string
	Status    string
	Category  *models.Category
}

func (this *Posts) Get(router *gin.Context) {
	page := plugins.StringtoInt64(router.Param("page"))
	if page == 0 {
		this.page(router)
	} else {
		this.pages(router, page)
	}
	return
}
func (this *Posts) page(router *gin.Context) {
	uid := router.GetInt64("auth")
	posts := new(models.Posts)
	DbValue := posts.Count_Account(uid)
	template := plugins.CountPage(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (this *Posts) pages(router *gin.Context, page int64) {
	uid := router.GetInt64("auth")
	posts := new(models.Posts)
	num := plugins.NumberPage()
	offset := (page - 1) * num
	DbValue := posts.Query_Account(uid, num, offset)
	template := DbPostsToAccountPosts(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})

}

func DbPostsToAccountPosts(value []models.Posts) []Postser {
	all := []Postser{}
	for v := range value {
		all = append(all, Postser{
			Id:        value[v].Id,
			Date:      value[v].Date,
			Title:     value[v].Title,
			Type:      value[v].Type,
			Thumbnail: plugins.CheckThumbnail(value[v].Thumbnail),
			Status:    value[v].Status,
			Category:  value[v].Category,
		})
	}
	return all

}
