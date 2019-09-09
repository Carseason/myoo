package account

import (
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type Comments struct{}
type Commentser struct {
	Id      int64
	Posts   CommentsPosts
	Content string
	Date    time.Time
}
type CommentsPosts struct {
	Id    int64
	Title string
}

func (this *Comments) Get(router *gin.Context) {
	page := plugins.StringtoInt64(router.Param("page"))
	if page == 0 {
		this.page(router)
	} else {
		this.pages(router, page)
	}
	return
}

func (*Comments) page(router *gin.Context) { //返回最大分页
	uid := router.GetInt64("auth")
	com := new(models.Comments)
	DbValue := com.Count_User(uid)
	template := plugins.CountPage(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (*Comments) pages(router *gin.Context, page int64) { //返回评论分页数据
	uid := router.GetInt64("auth")
	com := new(models.Comments)
	num := plugins.NumberPage() //获取分页最大数量
	offset := (page - 1) * num
	DbValue := com.Query_User(uid, num, offset)
	template := DbCommentsToAccountComments(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func DbCommentsToAccountComments(value []models.Comments) []Commentser {
	all := []Commentser{}
	for v := range value {
		all = append(all, Commentser{
			Id:      value[v].Id,
			Content: value[v].Content,
			Date:    value[v].Date,
			Posts: CommentsPosts{
				Id:    value[v].Posts.Id,
				Title: value[v].Posts.Title,
			},
		})
	}
	return all
}
