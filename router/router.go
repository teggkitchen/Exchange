package router

import (
	. "exchange/apis"
	// . "exchange/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/create", CreateMoney)
	router.PUT("/update/:id", UpdateMoney)
	router.DELETE("/delete/:id", DestroyMoney)

	// router.GET("/moneys", ShowProducts)

	return router
}
