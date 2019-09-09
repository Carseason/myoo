package account

import "github.com/gin-gonic/gin"

type Account struct{}

func (*Account) Get(router *gin.Context) {
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": nil})
	return
}
