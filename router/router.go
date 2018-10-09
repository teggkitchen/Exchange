package router

import (
	. "exchange/apis"
	// . "exchange/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/money", CreateMoney)
	router.PUT("/money/:id", UpdateMoney)
	router.DELETE("/money/:id", DestroyMoney)

	// router.GET("/moneys", ShowProducts)

	return router
}
