package admin

import (
	"encoding/json"
	"myoo/models"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Box struct{}

type Boxer struct {
	Title    string
	Url      string
	Icon     string
	Module   int //模块类型
	Number   int //数量
	Orderby  int //排序方式
	Category []int64
	Widget   Widget
	Code     string
}
type Widget struct {
	Orderby int
	Enabled int
	Title   string
	Number  int
	Code    string
}

func (*Box) Get(router *gin.Context) {
	box := []Boxer{}
	options := new(models.Options)
	DbValue := options.Query("Box")
	json.Unmarshal([]byte(DbValue["Box"]), &box)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": box})
	return
}
func (*Box) Post(router *gin.Context) {
	box := []Boxer{}
	data := router.PostForm("data")
	if err := json.Unmarshal([]byte(data), &box); err != nil {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
		return
	}
	for v := range box {
		if box[v].Module > 1 || box[v].Module < 0 {
			router.JSON(200, gin.H{"success": 403, "message": "模块数据异常", "data": nil})
			return
		}
		if box[v].Number < 1 {
			router.JSON(200, gin.H{"success": 403, "message": "模块数量不能少于1", "data": nil})
			return
		}
		if box[v].Widget.Enabled < 0 || box[v].Widget.Enabled > 3 {
			router.JSON(200, gin.H{"success": 403, "message": "小工具开关异常", "data": nil})
			return
		}
		if box[v].Widget.Number < 0 {
			router.JSON(200, gin.H{"success": 403, "message": "小工具数量异常", "data": nil})
			return
		}
		if len(box[v].Category) == 0 {
			router.JSON(200, gin.H{"success": 403, "message": "分类不能为空", "data": nil})
			return
		}
	}
	if jsonData, err := json.Marshal(box); err == nil {
		DbValue := make(map[string]string)
		DbValue["Box"] = string(jsonData)
		options := new(models.Options)
		if options.Update(DbValue) {
			defer redis.Del("Myoo/Index")
			router.JSON(200, gin.H{"success": 200, "message": "保存成功", "data": nil})
			return
		}
	}
	router.JSON(200, gin.H{"success": 403, "message": "保存失败", "data": nil})
	return
}
