package apis

import (
	code "exchange/config"
	msg "exchange/config"
	model "exchange/models"
	. "exchange/utils"
	"fmt"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

// 增加外幣
func CreateMoney(c *gin.Context) {
	var money model.Money

	// 取得參數
	tempNameStr := c.Request.FormValue("name")
	tempBuyStr := c.Request.FormValue("buy")
	tempBuy, err := strconv.ParseFloat(tempBuyStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}
	tempSellStr := c.Request.FormValue("sell")
	tempSell, err := strconv.ParseFloat(tempSellStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	//參數是否有值
	if len(tempNameStr) > 0 && tempBuy > 0 && tempSell > 0 {

		// 執行-檢查外幣名稱
		err := money.CheckMoneyName(tempNameStr)
		if err != nil {
			ShowJsonMSG(c, code.ERROR, msg.DATA_REPEAT_ERROR)
			return
		}

		// 執行-增加外幣
		err = money.InsertMoneyName(tempNameStr, tempBuy, tempSell)
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

// 查詢單一外幣
func QueryMoney(c *gin.Context) {
	var money model.Money
	// var moneyMarket model.MoneyMarket
	// 取得參數
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	//參數是否有值
	if id >= 0 {

		// 執行-查詢單一幣別行情
		result, err := money.QueryMoney(id)
		fmt.Println(result)
		if err != nil {
			ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
			return
		}
		ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}
}

// 查詢所有外幣
func QueryMoneys(c *gin.Context) {
	var money model.Money

	// 執行-查詢單一幣別行情
	result, err := money.QueryMoneys()
	fmt.Println(result)
	if err != nil {
		ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
		return
	}
	ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)

}

// 修改外幣
func UpdateMoney(c *gin.Context) {
	var money model.Money
	// 取得參數
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}
	tempBuyStr := c.Request.FormValue("buy")
	tempBuy, err := strconv.ParseFloat(tempBuyStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}
	tempSellStr := c.Request.FormValue("sell")
	tempSell, err := strconv.ParseFloat(tempSellStr, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	//參數是否有值
	if id >= 0 && tempBuy > 0 && tempSell > 0 {

		// 檢查外幣金額是否重複
		isRepeat := money.IsCheckMoneyMarket(id, tempBuy, tempSell)
		if isRepeat == true {
			ShowJsonMSG(c, code.ERROR, msg.PRICE_REPEAT_ERROR)
			return
		}

		// 執行-修改外幣
		err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
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

/////////////////////////
////    		     ////
////    以下測試用     ////
////      未整理      ////
////    		     ////
/////////////////////////

// 修改外幣 - 鎖測試
func UpdateMoneyGoroutine(id int64, tempBuy float64, tempSell float64) {
	var money model.Money
	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	var mu sync.Mutex     //宣告互斥鎖

	//參數是否有值
	if id >= 0 && tempBuy > 0 && tempSell > 0 {

		// 檢查外幣金額是否重複
		isRepeat := money.IsCheckMoneyMarket(id, tempBuy, tempSell)
		if isRepeat == true {
			fmt.Println(msg.PRICE_REPEAT_ERROR)
			return
		}
		wg.Add(1)
		mu.Lock() // 鎖住
		err := money.UpdateMoneyMarket(id, tempBuy, tempSell)
		if err != nil {
			mu.Unlock() // 解鎖
			wg.Done()
			wg.Wait()
			fmt.Println(msg.WRITE_ERROR)
			return
		}
		mu.Unlock() // 解鎖
		wg.Done()
		wg.Wait()
		fmt.Println(msg.UPDATE_SUCCESS)

	} else {
		// 缺少參數
		fmt.Println(msg.ARGS_ERROR)
		return
	}
}

// 修改外幣 - Goroutine測試
func UpdateMoneyGoroutine2(id int64, tempBuy float64, tempSell float64) {
	var money model.Money
	var err error

	//參數是否有值
	if id >= 0 && tempBuy >= 0 && tempSell >= 0 {
		go func() {
			err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
			if err != nil {
				fmt.Println("出錯:", msg.WRITE_ERROR)
				return
			}
			fmt.Println("變更完成")
			fmt.Println("成功:", msg.UPDATE_SUCCESS)
		}()
	} else {
		// 缺少參數
		fmt.Println("出錯:", msg.ARGS_ERROR)
		return
	}

}

var WorkingCountUpdate int

// 修改外幣 - Channel測試
func UpdateMoneyGoroutine3(id int64, tempBuy float64, tempSell float64) {
	var money model.Money
	var err error

	//參數是否有值
	if id >= 0 && tempBuy >= 0 && tempSell >= 0 {
		WorkingCountUpdate++
		finishChan := make(chan int)

		go func(c chan int) {
			err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
			if err != nil {
				fmt.Println("出錯:", msg.WRITE_ERROR)
				return
			}
			fmt.Println("變更完成")
			fmt.Println("成功:", msg.UPDATE_SUCCESS)

			c <- 1
		}(finishChan)

		finishLoop := false

		for {
			if finishLoop {
				break
			}
			select {
			case n := <-finishChan:
				WorkingCountUpdate -= n
				if WorkingCountUpdate == 0 {
					finishLoop = true
				}
			}
		}

	} else {
		// 缺少參數
		fmt.Println("出錯:", msg.ARGS_ERROR)
		return
	}

}

var WorkingCountQuery int

// 查詢外幣 - Channel測試
func QueryMoneyGoroutine() {
	var money model.Money
	// var err error

	//參數是否有值

	WorkingCountQuery++
	finishChan := make(chan int)

	go func(c chan int) {
		// 執行-查詢單一幣別行情
		result, err := money.QueryMoneys()
		fmt.Println(result)
		if err != nil {
			fmt.Println("出錯:", msg.NOT_FOUND_DATA_ERROR)
			return
		}
		fmt.Println(msg.EXEC_SUCCESS)

		c <- 1
	}(finishChan)

	finishLoop := false

	for {
		if finishLoop {
			break
		}
		select {
		case n := <-finishChan:
			WorkingCountQuery -= n
			if WorkingCountQuery == 0 {
				finishLoop = true
			}
		}
	}
}

// 併發 更新(未處理 條件競爭)
func TestGoroutineUpdate(c *gin.Context) {
	// 取得參數
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	fmt.Println("ID", id)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	// 測試併發 Update
	a := 1
	for a <= 1000 {
		a++
		fmt.Println("目前：", float64(a))
		UpdateMoneyGoroutine3(id, 1000, float64(a))
	}
}

// 併發 查詢
func TestGoroutineQuery(c *gin.Context) {
	// 測試併發 Query
	a := 1
	for a <= 1000 {
		a++
		fmt.Println("目前：", float64(a))
		QueryMoneyGoroutine()
	}

}
