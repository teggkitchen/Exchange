package router

import (
	. "exchange/apis"
	// . "exchange/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 增加外幣
	router.POST("/money", CreateMoney)

	// 修改外幣
	router.PUT("/money/:id", UpdateMoney)

	// 刪除外幣
	router.DELETE("/money/:id", DestroyMoney)

	// 查詢單一外幣
	router.GET("/money/:id", QueryMoney)

	// 查詢所有外幣
	router.GET("/moneys", QueryMoneys)

	/////////////////////////
	////    		     ////
	////    以下測試用     ////
	////    		     ////
	/////////////////////////
	router.PUT("/moneygo/:id", UpdateMoneyGoroutine)
	router.PUT("/moneygotest/:id", TestUpdateMoneyGoroutine)
	router.GET("/test/:id", TestUpdateMoney)

	return router
}
