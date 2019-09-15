package posts

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"myoo/router/v1/account"
	"myoo/router/v1/download"
	"time"

	"github.com/gin-gonic/gin"
)

type Posts struct{}

type Postser struct {
	Id         int64
	Title      string
	Date       time.Time
	Category   *models.Category
	Type       string
	Content    string
	Views      int64
	Thumbnail  string
	Author     int64
	Multimedia []account.Multimedia
	Download   string
	Features
}
type Features struct {
	Favorite struct {
		Status bool
	}
	Supported struct {
		Status bool
		Count  int64
	}
}

func (this *Posts) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "该文章不存在", "data": nil})
		return
	}

	template := this.cache(id)
	if template.Id != id {
		router.JSON(200, gin.H{"success": 403, "message": "该文章不存在", "data": nil})
		return
	}

	favorite := new(models.Favorite)
	supported := new(models.PostsSupported)
	if uid := router.GetInt64("auth"); uid > 0 {
		auth := redis.Query_User(uid)
		template.Favorite.Status = favorite.Is(auth.Id, id)
		template.Supported.Status = supported.Is(auth.Id, id)
	}
	template.Supported.Count = supported.Count(id)
	template.Views = this.views(id, template.Views)

	if template.Download != "" {
		template.Download = download.DownloadEncode(id)
	}

	// if !plugins.Group_Category_Check(auth.Lv, template.Category.Id) {
	// 	router.JSON(200, gin.H{"success": 302, "message": "当前用户组无法浏览本分类文章", "data": nil})
	// 	return
	// }
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (this *Posts) cache(id int64) Postser {
	cachePath := "Myoo/Posts/" + plugins.Int64toString(id)
	template := Postser{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		posts := new(models.Posts)
		if !posts.Is(id) {
			return template
		}
		DbValue := posts.Query_One(id)
		template.Id = DbValue.Id
		template.Title = DbValue.Title
		template.Date = DbValue.Date
		template.Category = DbValue.Category
		template.Type = DbValue.Type
		template.Content = DbValue.Content
		template.Views = DbValue.Views
		template.Thumbnail = DbValue.Thumbnail
		template.Author = DbValue.Author.Id
		multimedia := []account.Multimedia{}

		meta := new(models.PostsMeta)
		if template.Type == "music" {

			music := meta.Query(DbValue.Id, "Music")
			json.Unmarshal([]byte(music.Value), &multimedia)
		} else if template.Type == "video" {
			video := meta.Query(DbValue.Id, "Video")
			json.Unmarshal([]byte(video.Value), &multimedia)
		}
		template.Multimedia = multimedia
		if meta.Is(id, "Download") {
			template.Download = "1"
		}
		defer redis.PutStruct(cachePath, &template, 7200)
	}
	return template
}

func (*Posts) views(id, views int64) int64 {
	cachePath := "Myoo/Posts/" + plugins.Int64toString(id) + "/Views"
	posts := new(models.Posts)
	if redis.Is(cachePath) {
		views = redis.GetInt64(cachePath)
		redis.Incr(cachePath)
	} else {
		redis.Put(cachePath, views, 0)
		posts.Increase_Views(id)
	}
	if views%100 == 0 {
		posts.Update_Views(id, views+1)
	}
	return views
}
