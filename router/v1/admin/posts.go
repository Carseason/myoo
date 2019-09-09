package admin

import (
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type Posts struct{}
type Postser struct {
	Id        int64
	Date      time.Time //发布时间
	Title     string    //文章标题
	Type      string    //文章类型
	Thumbnail string    //文章略缩图
	Status    string    //文章状态，如已发布publish，待审核pending
	Author    PostsAuthorCategory
	Category  PostsAuthorCategory
}
type PostsAuthorCategory struct {
	Id   int64
	Name string
}

func (this *Posts) Get(router *gin.Context) {
	page := plugins.StringtoInt64(router.Param("page"))

	if page == 0 { //返回最大页数
		this.page(router)
	} else {
		this.pages(router, page)
	}
	return

}
func (this *Posts) page(router *gin.Context) {
	posts := new(models.Posts)
	DbValue := plugins.CountPage(posts.Count())
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": DbValue})
	return
}
func (this *Posts) pages(router *gin.Context, page int64) {
	posts := new(models.Posts)
	num := plugins.NumberPage() //获取分页最大数量
	offset := (page - 1) * num
	value := posts.Query_All(num, offset)
	template := PostsTemplate(value)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func PostsTemplate(value []models.Posts) []Postser {
	var all []Postser
	for v := range value {
		all = append(all, Postser{
			Id:        value[v].Id,
			Title:     value[v].Title,
			Thumbnail: value[v].Thumbnail,
			Date:      value[v].Date,
			Type:      value[v].Type,
			Category: PostsAuthorCategory{
				Id:   value[v].Category.Id,
				Name: value[v].Category.Name,
			},
			Status: value[v].Status,
			Author: PostsAuthorCategory{
				Id:   value[v].Author.Id,
				Name: value[v].Author.Name,
			},
		})
	}
	return all
}
