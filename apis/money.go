package apis

import (
	code "exchange/config"
	msg "exchange/config"
	model "exchange/models"
	. "exchange/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// go run 使用
// var productImgPath = config.IMAGE_PATH

//  go build 使用
// var productImgPath = GetAppPath() + config.IMAGE_PATH2

func TestCh(c *gin.Context) {

	ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, "result")
}

// // 展示全部外幣
// func ShowProducts(c *gin.Context) {
// 	var product model.Product
// 	// 執行-查詢全部外幣
// 	result, err := product.QueryProducts()

// 	if err != nil {
// 		ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
// 		return
// 	}
// 	ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)

// }

// 增加外幣
func CreateMoney(c *gin.Context) {
	var money model.Money

	// 取得參數
	tempNameStr := c.Request.FormValue("name")
	tempBuyStr := c.Request.FormValue("buy")
	tempBuy, err := strconv.ParseFloat(tempBuyStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR+"321")
		return
	}
	tempSellStr := c.Request.FormValue("sell")
	tempSell, err := strconv.ParseFloat(tempSellStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR+"123")
		return
	}

	//參數是否有值
	if len(tempNameStr) > 0 && tempBuy > 0 && tempSell > 0 {

		// 執行-增加外幣
		err := money.InsertMoneyName(tempNameStr, tempBuy, tempSell)
		if err != nil {
			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}

		ShowJsonDATA(c, code.SUCCESS, msg.CREATE_SUCCESS, "")

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

}

// 修改外幣
func UpdateMoney(c *gin.Context) {
	var money model.Money
	// 取得參數
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	tempNameStr := c.Request.FormValue("name")
	tempBuyStr := c.Request.FormValue("buy")
	tempBuy, err := strconv.ParseFloat(tempBuyStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR+"321")
		return
	}
	tempSellStr := c.Request.FormValue("sell")
	tempSell, err := strconv.ParseFloat(tempSellStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR+"123")
		return
	}

	//參數是否有值
	if len(tempNameStr) > 0 && tempBuy > 0 && tempSell > 0 {

		// 執行-修改外幣
		err = money.UpdateMoneyMarket(tempNameStr, tempBuy, tempSell)
		if err != nil {
			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}
		ShowJsonDATA(c, code.SUCCESS, msg.UPDATE_SUCCESS, "")

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

}

//刪除外幣
func DestroyMoney(c *gin.Context) {
	var money model.Money

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// 執行-刪除外幣
	err = money.DestroyMoneyMarket(id)
	if err != nil {
		//刪除失敗
		ShowJsonMSG(c, code.ERROR, msg.DELETE_ERROR)
		return
	}

	ShowJsonDATA(c, code.SUCCESS, msg.DELETE_SUCCESS, "")

}
