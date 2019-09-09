package account

import (
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type Favorite struct{}

type Favoriteer struct {
	Id        int64
	Date      time.Time
	Title     string
	Type      string
	Thumbnail string
	Category  *models.Category
	Status    bool
}

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

func DbPostsToFavorite(value []models.Posts) []Favoriteer {
	all := []Favoriteer{}
	for v := range value {
		all = append(all, Favoriteer{
			Id:        value[v].Id,
			Date:      value[v].Date,
			Title:     value[v].Title,
			Type:      value[v].Type,
			Thumbnail: plugins.CheckThumbnail(value[v].Thumbnail),
			Status:    true,
			Category:  value[v].Category,
		})
	}
	return all
}
