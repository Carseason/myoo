package inc

import (
	"myoo/redis"

	"github.com/gin-gonic/gin"
)

type Inc struct{}
type Template struct {
	Sitename        string
	Sitedescription string
	Message         string
	Search          string
	MobileSearch    string
	Ad              redis.Ad
}

func (*Inc) Get(this *gin.Context) {
	cachePath := "Myoo/Options/Inc"
	template := Template{}
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &template)
	} else {
		template.Sitename = redis.GetOptionsSitename()
		template.Sitedescription = redis.GetOptionsSitedescription()
		template.Message = redis.GetOptionsMessage()
		template.Search = redis.GetOptionsSearch()
		template.MobileSearch = redis.GetOptionsMobileSearch()
		template.Ad = redis.GetOptionsAd()
		redis.PutStruct(cachePath, &template, 7200)
	}
	this.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
