package category

import (
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Category struct{}
type CategoryPage struct {
	Name string
	plugins.Page
}

func (this *Category) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	page := plugins.StringtoInt64(router.Param("page"))

	if page == 0 {
		this.page(router, id)
	} else {
		this.pages(router, id, page)
	}
	return
}

func (this *Category) page(router *gin.Context, id int64) {
	cachePath := "Myoo/Category/" + router.Param("id") + "/" + router.Param("page")
	template := CategoryPage{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		category := new(models.Category)
		DbValue := category.Query(id)
		if DbValue.Id != id {
			router.JSON(200, gin.H{"success": 301, "message": "不存在的分类", "data": nil})
			return
		}
		template.Name = DbValue.Name
		posts := new(models.Posts)
		value := plugins.CountPage(posts.Count_Categor(id))
		template.MaxPage = value.MaxPage
		template.MaxCount = value.MaxCount
		defer redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func (this *Category) pages(router *gin.Context, id, page int64) {
	cachePath := "Myoo/Category/" + router.Param("id") + "/" + router.Param("page")
	template := []plugins.Posts{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		posts := new(models.Posts)
		num := plugins.NumberPage()
		offset := (page - 1) * num
		value := posts.Query_Category(id, num, offset, "-Id")
		template = plugins.CheckPostsAll(value)
		go redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
