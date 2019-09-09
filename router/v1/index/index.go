package index

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"myoo/router/v1/admin"

	"github.com/gin-gonic/gin"
)

type Index struct{}
type Indexer struct {
	Title  string
	Url    string
	Icon   string
	Module int //模块类型
	Posts  []plugins.Posts
	Widget struct {
		Enabled int
		Title   string
		Posts   []plugins.Posts
		Code    string
	}
	Code string
}

func (this *Index) Get(router *gin.Context) {
	cachePath := "Myoo/Index"
	template := []Indexer{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		template = this.CheckTemplate()
		defer redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func (*Index) CheckTemplate() []Indexer {
	box := []admin.Boxer{}
	options := new(models.Options)
	DbValue := options.Query("Box")
	json.Unmarshal([]byte(DbValue["Box"]), &box)
	template := []Indexer{}
	posts := new(models.Posts)
	for v := range box {
		tmp := Indexer{}
		tmp.Posts = plugins.CheckPostsAll(posts.Query_Filter(box[v].Category, box[v].Number, box[v].Orderby, 0))
		if len(tmp.Posts) == 0 {
			continue
		}
		tmp.Module = box[v].Module
		tmp.Title = box[v].Title
		tmp.Icon = box[v].Icon
		tmp.Url = box[v].Url
		tmp.Code = box[v].Code
		/******小工具*****/
		tmp.Widget.Enabled = box[v].Widget.Enabled
		if box[v].Widget.Enabled > 0 {
			tmp.Widget.Title = box[v].Widget.Title
			if tmp.Widget.Enabled == 3 {
				tmp.Widget.Code = box[v].Widget.Code
			} else {
				tmp.Widget.Posts = plugins.CheckPostsAll(posts.Query_Filter(box[v].Category, box[v].Widget.Number, box[v].Widget.Orderby, 0))
			}
		}
		template = append(template, tmp)
	}
	return template
}
