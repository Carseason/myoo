package admin

import (
	"myoo/models"
	"myoo/plugins"
	"time"

	"github.com/gin-gonic/gin"
)

type Users struct{}
type Userser struct {
	Id     int64
	Email  string
	Name   string
	Date   time.Time
	Avatar string
	Lv     int64 //等级
}

func (this *Users) Get(router *gin.Context) {
	page := plugins.StringtoInt64(router.Param("page"))
	if page == 0 { //返回最大页数
		this.page(router)
	} else {
		this.pages(router, page)
	}
	return
}
func (this *Users) page(router *gin.Context) {
	user := new(models.User)
	DbValue := plugins.CountPage(user.Count())
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": DbValue})
	return

}
func (this *Users) pages(router *gin.Context, page int64) {
	user := new(models.User)
	num := plugins.NumberPage() //获取分页最大数量
	offset := (page - 1) * num
	DbValue := user.Query_All(num, offset)
	template := UsersTemplate(DbValue)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func UsersTemplate(value []models.User) []Userser {
	var all []Userser
	for v := range value {
		all = append(all, Userser{
			Id:     value[v].Id,
			Avatar: value[v].Avatar,
			Name:   value[v].Name,
			Email:  value[v].Email,
			Lv:     value[v].Lv,
			Date:   value[v].Date,
		})
	}
	return all
}
