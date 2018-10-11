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
	////      測試用      ////
	////    		     ////
	/////////////////////////

	// 併發 查詢
	router.GET("/testQuery", TestGoroutineQuery)

	// 併發 更新(未處理 條件競爭)
	router.GET("/testUpdate/:id", TestGoroutineUpdate)

	return router
}
