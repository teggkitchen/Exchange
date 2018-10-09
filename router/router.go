package router

import (
	. "exchange/apis"
	// . "exchange/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/create_money", CreateMoney)
	// router.GET("/moneys", ShowProducts)
	// router.PUT("/money/:id", UpdateProduct)
	// router.DELETE("/money/:id", DestroyProduct)

	return router
}
