package account

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type NewPosts struct{}

type Multimedia struct {
	Title string
	Url   string
}
type Download struct {
	Title       string
	Url         string
	DownloadPwd string
	ExtractPwd  string
}

type NewPostser struct {
	Category []models.Category
}

func (*NewPosts) Get(router *gin.Context) { //返回所有分类
	cachePath := "Myoo/NewPosts"
	template := NewPostser{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		template.Category = models.NewCategory().QueryAll()
		defer redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (this *NewPosts) Post(router *gin.Context) {
	uid := router.GetInt64("auth")
	auth := redis.Query_User(uid)
	if !plugins.Group_NewPosts(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组权限不足", "data": nil})
		return
	}

	title := router.PostForm("title")
	types := router.PostForm("type")
	thumbnail := router.PostForm("thumbnail")
	content := router.PostForm("content")
	music := []Multimedia{}
	video := []Multimedia{}
	download := []Download{}
	if title == "" {
		router.JSON(200, gin.H{"success": 403, "message": "投稿标题不能为空", "data": nil})
		return
	}
	if types != "music" && types != "standard" && types != "video" && types != "image" {
		router.JSON(200, gin.H{"success": 403, "message": "投稿类型异常", "data": nil})
		return
	}
	category := plugins.StringtoInt64(router.PostForm("category"))
	if category == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "投稿分类异常", "data": nil})
		return
	}
	if !models.NewCategory().Is(category) {
		router.JSON(200, gin.H{"success": 403, "message": "投稿分类异常", "data": nil})
		return
	}

	if json.Unmarshal([]byte(router.PostForm("music")), &music) != nil {
		router.JSON(200, gin.H{"success": 403, "message": "投稿音乐资源异常", "data": nil})
		return
	}
	for v := range music {
		if music[v].Title == "" {
			router.JSON(200, gin.H{"success": 403, "message": "音乐资源名称不能为空", "data": nil})
			return
		}
	}
	for v := range video {
		if video[v].Title == "" {
			router.JSON(200, gin.H{"success": 403, "message": "视频资源名称不能为空", "data": nil})
			return
		}
	}

	if json.Unmarshal([]byte(router.PostForm("video")), &video) != nil {
		router.JSON(200, gin.H{"success": 403, "message": "投稿视频资源异常", "data": nil})
		return
	}
	if json.Unmarshal([]byte(router.PostForm("download")), &download) != nil {
		router.JSON(200, gin.H{"success": 403, "message": "投稿下载资源异常", "data": nil})
		return
	}
	if thumbnail != "" && !plugins.An_Image(thumbnail) {
		router.JSON(200, gin.H{"success": 403, "message": "文章略缩图只支持网络图像URL", "data": nil})
		return
	}
	if thumbnail != "" {
		thumbnail = extractThumbnail(thumbnail)
	}
	if plugins.RegexpMatch(content, `(<\s{0,}script|script\s{0,}>)`) {
		router.JSON(200, gin.H{"success": 403, "message": "文章内容异常", "data": nil})
		return
	}

	posts := new(models.Posts)
	posts.Title = title
	posts.Content = content
	posts.Type = types
	posts.Thumbnail = thumbnail
	posts.Views = 1
	if auth.Admin {
		posts.Status = "publish"
	} else {
		posts.Status = "pedding"
	}
	posts.Author = &models.User{
		Id: uid,
	}
	posts.Category = &models.Category{
		Id: category,
	}

	if v := models.NewPosts().Insert(posts); v > 0 {
		/***************多媒体资源*****************/
		meta := new(models.PostsMeta)
		meta.PostsId = v
		if types == "music" && len(music) > 0 {
			meta.Name = "Music"
			if jsonData, err := json.Marshal(music); err == nil {
				meta.Value = string(jsonData)
			}
			new(models.PostsMeta).Insert(meta)
		} else if types == "video" && len(video) > 0 {
			meta.Name = "Video"
			if jsonData, err := json.Marshal(video); err == nil {
				meta.Value = string(jsonData)
			}
			new(models.PostsMeta).Insert(meta)
		}
		/***************下载资源*****************/
		if len(download) > 0 {
			down := new(models.PostsMeta)
			down.PostsId = v
			down.Name = "Download"
			if jsonData, err := json.Marshal(download); err == nil {
				down.Value = string(jsonData)
			}
			models.NewPostsMeta().Insert(down)
		}
		router.JSON(200, gin.H{"success": 200, "message": "文章投稿成功,请静候审核...", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "文章投稿失败", "data": nil})
	return
}

/***********文章略缩图提取**************/
func extractThumbnail(value string) string {
	return plugins.RegexFindString(value, `https?:\/\/[^\/]+[^\.]+\.(jpg|png|gif|jpge|jpeg|webp)`)
}
