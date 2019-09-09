package admin

import (
	"encoding/json"
	"myoo/models"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Options struct{}

type Optionser struct {
	Sitename        string
	Sitedescription string
	Registered      bool
	Download        bool
	Search          string
	MobileSearch    string
	Email           Email
	Comments        Comments
	DefaultCategory int64
	Ad              Ad
	Message         string
	Code            Code
	DefaultLv       int64
}
type Video struct {
	Keywords   string
	Url        string
	SourceType int64
}
type Email struct {
	Url           string
	FromEmail     string
	FromUname     string
	EmailUsername string
	EmailPassword string
	EmailHost     string
	EmailPort     int
}
type Comments struct {
	KeywordsFilter      string
	SensitivewordBefore string
	SensitivewordRear   string
}
type Code struct {
	Style  string
	Script string
}
type Ad struct {
	Header string
	Footer string
}

func (*Options) Get(router *gin.Context) {
	options := new(models.Options)
	MapData := options.Query("Sitename", "Sitedescription", "Registered", "Download", "Search", "MobileSearch", "Email", "Comments", "DefaultCategory", "Ad", "Message", "DefaultLv")
	template := new(Optionser)
	template.Sitename = MapData["Sitename"]
	template.Sitedescription = MapData["Sitedescription"]
	template.Search = MapData["Search"]
	template.MobileSearch = MapData["MobileSearch"]
	template.Message = MapData["Message"]

	if MapData["Registered"] == "" {
		template.Registered = true
	} else {
		json.Unmarshal([]byte(MapData["Registered"]), &template.Registered)
	}
	json.Unmarshal([]byte(MapData["DefaultLv"]), &template.DefaultLv)
	json.Unmarshal([]byte(MapData["Download"]), &template.Download)
	json.Unmarshal([]byte(MapData["Email"]), &template.Email)
	json.Unmarshal([]byte(MapData["Comments"]), &template.Comments)
	json.Unmarshal([]byte(MapData["DefaultCategory"]), &template.DefaultCategory)
	json.Unmarshal([]byte(MapData["Ad"]), &template.Ad)
	json.Unmarshal([]byte(MapData["Code"]), &template.Code)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}

func (*Options) Post(router *gin.Context) {
	data := router.PostForm("data")
	var template Optionser
	if json.Unmarshal([]byte(data), &template) != nil {
		router.JSON(200, gin.H{"success": 403, "message": "数据异常", "data": nil})
		return
	}
	if template.Sitename == "" {
		router.JSON(200, gin.H{"success": 403, "message": "网站名称不能为空", "data": nil})
		return
	}
	if template.Sitedescription == "" {
		router.JSON(200, gin.H{"success": 403, "message": "网站副标题不能为空", "data": nil})
		return
	}
	if template.Search == "" || template.MobileSearch == "" {
		router.JSON(200, gin.H{"success": 403, "message": "搜索链接不能为空", "data": nil})
		return
	}
	if template.DefaultLv <= 0 {
		router.JSON(200, gin.H{"success": 403, "message": "默认用户等级过低", "data": nil})
		return
	}

	MapData := make(map[string]string)
	MapData["Sitename"] = template.Sitename
	MapData["Sitedescription"] = template.Sitedescription
	MapData["Search"] = template.Search
	MapData["Search"] = template.Search
	MapData["MobileSearch"] = template.MobileSearch
	MapData["Message"] = template.Message

	if jsonData, err := json.Marshal(template.DefaultLv); err == nil {
		MapData["DefaultLv"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.Registered); err == nil {
		MapData["Registered"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.Download); err == nil {
		MapData["Download"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.Email); err == nil {
		MapData["Email"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.Comments); err == nil {
		MapData["Comments"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.DefaultCategory); err == nil {
		MapData["DefaultCategory"] = string(jsonData)
	}
	if jsonData, err := json.Marshal(template.Ad); err == nil {
		MapData["Ad"] = string(jsonData)
	}
	options := new(models.Options)
	if options.Update(MapData) {
		defer redis.DelAll("Myoo/Options/*")
		router.JSON(200, gin.H{"success": 200, "message": "保存成功", "data": nil})
	} else {
		router.JSON(200, gin.H{"success": 200, "message": "保存失败", "data": nil})
	}
	return
}
