package download

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"myoo/router/v1/account"
	"time"

	"github.com/gin-gonic/gin"
)

type Download struct{}
type Downloader struct {
	account.Download
}

type Path struct {
	Id      int64
	OutTime int64
}

func DownloadEncode(id int64) string {
	template := Path{}
	template.Id = id
	template.OutTime = time.Now().Unix() + 600
	if jsonData, err := json.Marshal(template); err == nil {
		return plugins.Aes_Encode(string(jsonData))
	}
	return ""
}

func DownloadDecode(value string) int64 {
	template := Path{}
	token, err := plugins.Aes_Decode(value)
	if err != nil {
		return 0
	}
	if json.Unmarshal([]byte(token), &template) != nil {
		return 0
	}
	if template.Id == 0 || template.OutTime < time.Now().Unix() {
		return 0
	}
	return template.Id
}
func (this *Download) Get(router *gin.Context) {
	path := router.Param("path")

	id := DownloadDecode(path)
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "该链接已失效", "data": nil})
		return
	}
	if !redis.GetOptionsDownload() {
		template := this.download(id)
		router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
		return
	}
	/**************登录下载***************/
	uid := router.GetInt64("auth")
	if uid == 0 {
		router.JSON(200, gin.H{"success": 302, "message": "该下载需要先登录", "data": nil})
		return
	}
	auth := redis.Query_User(uid)
	if !plugins.Group_Download(auth.Lv) {
		router.JSON(200, gin.H{"success": 302, "message": "当前用户组权限不足", "data": nil})
		return
	}
	template := this.download(id)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return

}

func (this *Download) download(id int64) []Downloader {
	cachePath := "Myoo/Posts/" + plugins.Int64toString(id) + "/Download"
	template := []Downloader{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		meta := new(models.PostsMeta)
		DbValue := meta.Query(id, "Download")
		json.Unmarshal([]byte(DbValue.Value), &template)
		redis.PutStruct(cachePath, &template, 3600)
	}
	return template
}
