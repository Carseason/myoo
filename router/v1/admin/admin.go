package admin

import "github.com/gin-gonic/gin"

type Admin struct{}

func (*Admin) Get(router *gin.Context) {
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": nil})
	return
}
