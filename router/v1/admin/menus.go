package admin

import (
	"encoding/json"
	"myoo/models"
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Menus struct{}
type Menuser struct {
	TopMenus []NavMenus
	NavMenus []NavMenus
	Links    []Child
}
type NavMenus struct {
	Title string
	Url   string
	Icon  string
	Child []Child
}
type Child struct {
	Title string
	Url   string
}

func (*Menus) Get(router *gin.Context) {
	options := new(models.Options)
	MapData := options.Query("TopMenus", "NavMenus", "Links")
	template := new(Menuser)
	json.Unmarshal([]byte(MapData["TopMenus"]), &template.TopMenus)
	json.Unmarshal([]byte(MapData["NavMenus"]), &template.NavMenus)
	json.Unmarshal([]byte(MapData["Links"]), &template.Links)
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
func (*Menus) Post(router *gin.Context) {
	template := new(Menuser)
	data := router.PostForm("data")
	if err := json.Unmarshal([]byte(data), &template); err == nil {
		/*数据验证*/
		for v := range template.NavMenus {
			if template.NavMenus[v].Title == "" {
				router.JSON(200, gin.H{"success": 403, "message": "导航菜单名称不能为空", "data": nil})
				return
			}
			for j := range template.NavMenus[v].Child {
				if template.NavMenus[v].Child[j].Title == "" {
					router.JSON(200, gin.H{"success": 403, "message": "二级菜单名称不能为空", "data": nil})
					return
				}
			}
		}
		for v := range template.TopMenus {
			if template.TopMenus[v].Title == "" {
				router.JSON(200, gin.H{"success": 403, "message": "置顶菜单名称不能为空", "data": nil})
				return
			}
			for j := range template.TopMenus[v].Child {
				if template.TopMenus[v].Child[j].Title == "" {
					router.JSON(200, gin.H{"success": 403, "message": "二级菜单名称不能为空", "data": nil})
					return
				}
			}
		}

		for v := range template.Links {
			if template.Links[v].Title == "" {
				router.JSON(200, gin.H{"success": 403, "message": "友联名称不能为空", "data": nil})
				return
			}
		}
		MapData := make(map[string]string)
		if jsonData, err := json.Marshal(template.NavMenus); err == nil {
			MapData["NavMenus"] = string(jsonData)
		}
		if jsonData, err := json.Marshal(template.Links); err == nil {
			MapData["Links"] = string(jsonData)
		}
		if jsonData, err := json.Marshal(template.TopMenus); err == nil {
			MapData["TopMenus"] = string(jsonData)
		}
		options := new(models.Options)
		if options.Update(MapData) {
			defer redis.Del("Myoo/Menus")
			router.JSON(200, gin.H{"success": 200, "message": "保存成功", "data": nil})
		}
	} else {
		router.JSON(200, gin.H{"success": 403, "message": err, "data": nil})

	}
	return
}
