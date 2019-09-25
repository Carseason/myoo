package admin

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"myoo/router/v1/account"

	"github.com/gin-gonic/gin"
)

type ModifyPosts struct{}
type ModifyPostser struct {
	Posts struct {
		Id        int64
		Title     string
		Content   string
		Type      string
		Thumbnail string
		Status    string
		Category  int64
	}
	Video    []Multimedia
	Music    []Multimedia
	Download []Download
}

type Multimedia struct {
	account.Multimedia
}
type Download struct {
	account.Download
}

func (*ModifyPosts) Get(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "不存在的文章", "data": nil})
		return
	}
	template := ModifyPostser{}
	posts := new(models.Posts)
	DbValue := posts.Query(id)
	template.Posts.Id = DbValue.Id
	template.Posts.Title = DbValue.Title
	template.Posts.Content = DbValue.Content
	template.Posts.Type = DbValue.Type
	template.Posts.Thumbnail = DbValue.Thumbnail
	template.Posts.Status = DbValue.Status
	template.Posts.Category = DbValue.Category.Id

	meta := new(models.PostsMeta)
	if DbValue.Type == "music" {
		music := meta.Query(DbValue.Id, "Music")
		json.Unmarshal([]byte(music.Value), &template.Music)
	} else if DbValue.Type == "video" {
		video := meta.Query(DbValue.Id, "Video")
		json.Unmarshal([]byte(video.Value), &template.Video)
	}

	if meta.Is(DbValue.Id, "Download") {
		download := meta.Query(DbValue.Id, "Download")
		json.Unmarshal([]byte(download.Value), &template.Download)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return

}
func (this *ModifyPosts) Post(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "不存在的文章", "data": nil})
		return
	}
	switch router.PostForm("types") {
	case "update":
		this.update(router, id)
	case "delete":
		this.delete(router, id)
	default:
		router.JSON(200, gin.H{"success": 403, "message": "不存在的文章", "data": nil})
		return
	}
	cachePath := "Myoo/Posts/" + router.Param("id") + `(/*)?`
	redis.DelAll(cachePath)
	return
}

func (this *ModifyPosts) update(router *gin.Context, id int64) {

	title := router.PostForm("title")
	types := router.PostForm("type")
	thumbnail := router.PostForm("thumbnail")
	content := router.PostForm("content")
	music := []Multimedia{}
	video := []Multimedia{}
	download := []Download{}
	status := router.DefaultPostForm("status", "pedding")
	if title == "" {
		router.JSON(200, gin.H{"success": 403, "message": "投稿标题不能为空", "data": nil})
		return
	}
	if types != "music" && types != "standard" && types != "video" && types != "image" {
		router.JSON(200, gin.H{"success": 403, "message": "投稿类型异常", "data": nil})
		return
	}
	if status != "pedding" && status != "publish" {
		router.JSON(200, gin.H{"success": 403, "message": "文章状态异常", "data": nil})
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

	template := new(models.Posts)
	template.Id = id
	template.Title = title
	template.Content = content
	template.Type = types
	template.Thumbnail = thumbnail
	template.Category = &models.Category{
		Id: category,
	}
	template.Status = status
	posts := new(models.Posts)
	if posts.Update(template) {
		/***************多媒体资源*****************/
		if types == "music" {
			this.multimedia(id, "Music", music)
		} else if types == "video" {
			this.multimedia(id, "Video", video)
		}
		/***************下载资源*****************/
		this.download(id, download)

		router.JSON(200, gin.H{"success": 200, "message": "已更新", "data": nil})
	} else {
		router.JSON(200, gin.H{"success": 403, "message": "更新失败", "data": nil})
	}
	return
}

func (*ModifyPosts) multimedia(id int64, name string, music []Multimedia) {
	meta := new(models.PostsMeta)
	if len(music) > 0 {
		template := new(models.PostsMeta)
		template.PostsId = id
		template.Name = "Music"
		if jsonData, err := json.Marshal(music); err == nil {
			template.Value = string(jsonData)
		}
		if meta.Is(id, "Music") {
			meta.Update(id, "Music", template.Value)
		} else {
			meta.Insert(template)
		}
	} else if meta.Is(id, "Music") {
		meta.Delete(id, "Music")
	}
}
func (*ModifyPosts) download(id int64, download []Download) {
	meta := new(models.PostsMeta)
	if len(download) > 0 {
		down := new(models.PostsMeta)
		down.PostsId = id
		down.Name = "Download"
		if jsonData, err := json.Marshal(download); err == nil {
			down.Value = string(jsonData)
		}
		if meta.Is(id, "Download") {
			meta.Update(id, "Download", down.Value)
		} else {
			meta.Insert(down)
		}
	} else {
		if meta.Is(id, "Download") {
			meta.Delete(id, "Download")
		}
	}
}

func (*ModifyPosts) delete(router *gin.Context, id int64) {
	posts := new(models.Posts)
	if posts.Delete(id) {
		new(models.Comments).Delete_All(id)
		new(models.PostsMeta).Delete_All(id)
		new(models.Favorite).Delete_All(id)
		new(models.PostsSupported).Delete_All(id)
		router.JSON(200, gin.H{"success": 200, "message": "已删除", "data": nil})
	} else {
		router.JSON(200, gin.H{"success": 403, "message": "删除失败", "data": nil})
	}

	return
}

/***********文章略缩图提取**************/
func extractThumbnail(value string) string {
	return plugins.RegexFindString(value, `(https?:\/\/[^\/]+[^\.]+\.(jpg|png|gif|jpge))`)
}
