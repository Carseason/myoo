package admin

import (
	"encoding/json"
	"myoo/models"

	//	"myoo/plugins"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Group struct{}
type Grouper struct {
	Lv       int64
	NickName string
	//	Medal        []string
	Category     []int64
	NewPosts     bool
	Comments     bool
	Supported    bool
	Favorite     bool
	Followers    bool
	Download     bool
	Checkin      bool
	UpdateAvatar bool
}

func (*Group) Get(router *gin.Context) {
	groups := []Grouper{}
	options := new(models.Options)
	DbValue := options.Query("Group")
	json.Unmarshal([]byte(DbValue["Group"]), &groups)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": groups})
	return
}
func (*Group) Post(router *gin.Context) {
	groups := []Grouper{}
	data := router.PostForm("data")
	if err := json.Unmarshal([]byte(data), &groups); err != nil {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})
		return
	}
	t := len(groups)
	for v := range groups {
		if groups[v].Lv <= 0 {
			router.JSON(200, gin.H{"success": 403, "message": "用户组等级异常", "data": nil})
			return
		}
		if groups[v].NickName == "" {
			router.JSON(200, gin.H{"success": 403, "message": "用户组别名不能为空", "data": nil})
			return
		}
		for i := v + 1; i < t; i++ {
			if groups[v].Lv == groups[i].Lv {
				router.JSON(200, gin.H{"success": 403, "message": "用户组等级不可重复", "data": nil})
				return
			}
		}
		// for j := range groups[v].Medal {
		// 	if !plugins.An_Image(groups[v].Medal[j]) {
		// 		router.JSON(200, gin.H{"success": 403, "message": "用户勋章只接受网络URL图像地址", "data": nil})
		// 		return
		// 	}
		// }
	}
	if jsonData, err := json.Marshal(groups); err == nil {
		DbValue := make(map[string]string)
		DbValue["Group"] = string(jsonData)
		if models.NewOptions().Update(DbValue) {
			defer redis.Del("Myoo/Group")
			router.JSON(200, gin.H{"success": 200, "message": "保存成功", "data": nil})
			return
		}
	}
	router.JSON(200, gin.H{"success": 403, "message": "保存失败", "data": nil})
	return
}
