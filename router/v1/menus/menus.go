package menus

import (
	"encoding/json"
	"myoo/models"
	"myoo/redis"
	"myoo/router/v1/admin"

	"github.com/gin-gonic/gin"
)

type Menus struct{}
type Menuser struct{}

func (*Menus) Get(router *gin.Context) {
	template := admin.Menuser{}
	cachePath := "Myoo/Menus"
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		options := new(models.Options)
		value := options.Query("TopMenus", "NavMenus", "Links")
		json.Unmarshal([]byte(value["TopMenus"]), &template.TopMenus)
		json.Unmarshal([]byte(value["NavMenus"]), &template.NavMenus)
		json.Unmarshal([]byte(value["Links"]), &template.Links)
		redis.PutStruct(cachePath, &template, 7200)
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
