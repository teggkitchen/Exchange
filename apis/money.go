package apis

import (
	code "exchange/config"
	msg "exchange/config"
	model "exchange/models"
	. "exchange/utils"
	"fmt"
	"math/rand"
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
////    以下測試用    ////
////    		     ////
/////////////////////////

// 修改外幣 - 併發
func UpdateMoneyGoroutine(c *gin.Context) {
	var money model.Money
	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	var mu sync.Mutex     //宣告互斥鎖

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
		wg.Add(1)
		mu.Lock() // 鎖住
		err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
		if err != nil {
			mu.Unlock() // 解鎖
			wg.Done()
			wg.Wait()
			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}
		mu.Unlock() // 解鎖
		wg.Done()
		wg.Wait()
		ShowJsonDATA(c, code.SUCCESS, msg.UPDATE_SUCCESS, "")

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}
}

type UpdateTemp struct {
	MoneyId int64   `json:"moneyId"` // 外幣Id
	Buy     float64 `json:"buy"`     // 買入
	Sell    float64 `json:"sell"`    // 賣出
}

var updateTemps []UpdateTemp

func UpdateMoneyGoroutine2(id int64, tempBuy float64, tempSell float64) {
	var money model.Money
	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	var mu sync.Mutex     //宣告互斥鎖
	var resId int64
	var err error

	//參數是否有值
	if id >= 0 && tempBuy >= 0 && tempSell >= 0 {
		resId = -1
		// 檢查外幣金額是否重複
		// isRepeat := money.IsCheckMoneyMarket(id, tempBuy, tempSell)
		// if isRepeat == true {
		// 	fmt.Println("出錯:", msg.PRICE_REPEAT_ERROR)
		// 	return
		// }
		fmt.Println("待寫入數量", len(updateTemps))
		updateTemp := UpdateTemp{MoneyId: id, Buy: tempBuy, Sell: tempSell}
		updateTemps = append(updateTemps, updateTemp)

		wg.Add(1)
		fmt.Println("變更準備")
		mu.Lock() //互斥鎖 - 鎖住
		fmt.Println("鎖住")

		resId, err = money.UpdateMoneyMarketTemp(id, tempBuy, tempSell)
		if err != nil {
			fmt.Println("出錯:", msg.WRITE_ERROR)
			return
		}
		for len(updateTemps) > 0 {
			updateTemps = updateTemps[1:len(updateTemps)]
			// fmt.Println("當前", updateTemps)
			// fmt.Println("當前數量", len(updateTemps))
		}
		fmt.Println(resId)
		mu.Unlock() //互斥鎖 - 解鎖
		fmt.Println("解鎖")
		fmt.Println("變更完成")
		wg.Done()
		fmt.Println("成功:", msg.UPDATE_SUCCESS)
		wg.Wait()

	} else {
		// 缺少參數
		fmt.Println("出錯:", msg.ARGS_ERROR)
		return
	}

}

func UpdateMoneyGoroutineTemp(id int64, tempBuy float64, tempSell float64) {
	var money model.Money
	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	var mu sync.Mutex     //宣告互斥鎖
	var resId int64
	var err error

	//參數是否有值
	if id >= 0 && tempBuy > 0 && tempSell > 0 {
		resId = -1
		// 檢查外幣金額是否重複
		isRepeat := money.IsCheckMoneyMarket(id, tempBuy, tempSell)
		if isRepeat == true {
			fmt.Println("出錯:", msg.PRICE_REPEAT_ERROR)
			return
		}
		wg.Add(1)
		fmt.Println("變更準備")
		mu.Lock() //互斥鎖 - 鎖住
		fmt.Println("鎖住")
		resId, err = money.UpdateMoneyMarketTemp(id, tempBuy, tempSell)
		if err != nil {
			fmt.Println("出錯:", msg.WRITE_ERROR)
			return
		}
		if resId >= 0 {
			mu.Unlock() //互斥鎖 - 解鎖
			fmt.Println("解鎖")
			fmt.Println("變更完成")
			wg.Done()
			fmt.Println("成功:", msg.UPDATE_SUCCESS)
			wg.Wait()
		}

	} else {
		// 缺少參數
		fmt.Println("出錯:", msg.ARGS_ERROR)
		return
	}

}

// 測試Update
func TestUpdateMoneyGoroutine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		// 參數錯誤
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	for i := 0; i < 1000; i++ {
		go func(ii int) {
			p := rand.Intn(300)
			fmt.Printf("Hello %d\n", p)
			UpdateMoneyGoroutine2(id, 100, float64(p))
		}(i)
	}
	// intArr := make([]int, 0)
	// for i := 0; i < 10; i++ {
	// 	intArr = append(intArr, i)
	// }

	// fmt.Println("-初始-", intArr)
	// for len(intArr) > 0 {
	// 	intArr = intArr[1:len(intArr)]
	// 	fmt.Println("當前", intArr)
	// 	fmt.Println("當前數量", len(intArr))
	// }

	// fmt.Println("結果", intArr)

}

func TestMoney(c *gin.Context) {

	// // Channel實作
	// var NowCount int = 100 //總數100
	// coubnt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// rChan := make(chan int, len(coubnt))
	// defer close(rChan)

	// for i, _ := range coubnt {
	// 	go func() {
	// 		NowCount = (NowCount - 10)
	// 		rChan <- NowCount
	// 		fmt.Println("A-第", i, "次")
	// 	}()
	// }
	// for i, _ := range coubnt {
	// 	params := <-rChan
	// 	fmt.Println("B-第", i, "次")
	// 	fmt.Println("B-結果：", params)
	// }

	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	a := 0
	for a <= 1000 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			a++
			var money model.Money

			// 執行-查詢單一幣別行情
			result, err := money.QueryMoneys()
			fmt.Println(result)

			if err != nil {
				wg.Wait()
				ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
				return
			}
			wg.Wait()
			ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)
		}()
	}

	// // Channel實作
	// var NowCount int = 100 //總數100
	// coubnt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// rChan := make(chan int, len(coubnt))
	// defer close(rChan)

	// for i, _ := range coubnt {
	// 	go func() {
	// 		NowCount = (NowCount - 10)
	// 		rChan <- NowCount
	// 		fmt.Println("A-第", i, "次")
	// 	}()
	// }
	// for i, _ := range coubnt {
	// 	params := <-rChan
	// 	fmt.Println("B-第", i, "次")
	// 	fmt.Println("B-結果：", params)
	// }

	// rChan := make(chan int, len(coubnt))
	// var wg sync.WaitGroup  //存放Thread的空間，歸0則運行主程式
	// var NowCount int = 100 //總數100

	// var mu sync.Mutex //宣告互斥鎖
	// wg.Add(10)        // 在WaitGroup可以容納Thread的數目

	// fmt.Println("總金額：", NowCount)
	// for true {

	// 	fmt.Println("提領任務準備")
	// 	go func() {
	// 		if NowCount <= 0 {
	// 			fmt.Println("提領任務開始")
	// 			mu.Lock() //互斥鎖 - 鎖住
	// 			fmt.Println("A - 提領中")
	// 			// NowCount -= 10                     //如果直接計算NowCount，可以正常交易
	// 			transCount := NowCount             //定義transCount
	// 			time.Sleep(500 * time.Millisecond) //暫停兩秒，等待第二個Thread執行
	// 			transCount -= 10                   //transCount = transCount-10, 每次扣10
	// 			NowCount = transCount              //transCount的值回放到NowCount
	// 			mu.Unlock()                        //互斥鎖 - 解鎖
	// 			fmt.Println("A - 提領後：", NowCount)
	// 			fmt.Println("A - 提領完成")
	// 			wg.Done()
	// 		}
	// 	}()
	// 	time.Sleep(100 * time.Millisecond) //暫停兩秒，等待第二個Thread執行
	// }

	// go func() {
	// 	mu.Lock() //互斥鎖 - 鎖住
	// 	fmt.Println("B - 提領中")
	// 	// NowCount -= 10 //如果直接計算NowCount，可以正常交易

	// 	transCount := NowCount             //定義transCount
	// 	time.Sleep(500 * time.Millisecond) //暫停兩秒，等待第二個Thread執行
	// 	transCount -= 10                   //transCount = transCount-10, 每次扣10
	// 	NowCount = transCount              //transCount的值回放到NowCount
	// 	mu.Unlock()                        //互斥鎖 - 解鎖
	// 	fmt.Println("B - 提領後：", NowCount)
	// 	fmt.Println("B - 提領完成")
	// 	wg.Done()
	// }()

	// go func() {
	// 	mu.Lock() //互斥鎖 - 鎖住
	// 	fmt.Println("C - 提領中")
	// 	// NowCount -= 10 //如果直接計算NowCount，可以正常交易

	// 	transCount := NowCount             //定義transCount
	// 	time.Sleep(500 * time.Millisecond) //暫停兩秒，等待第二個Thread執行
	// 	transCount -= 10                   //transCount = transCount-10, 每次扣10
	// 	NowCount = transCount              //transCount的值回放到NowCount
	// 	mu.Unlock()                        //互斥鎖 - 解鎖
	// 	fmt.Println("C - 提領後：", NowCount)
	// 	fmt.Println("C - 提領完成")
	// 	wg.Done()
	// }()

	// wg.Wait() //等待WaitGroup裡的Thread，全部執行完畢後，再繼續執行以下

	// fmt.Println("結束結束")
}

// 修改外幣
func UpdateMoneyGoroutine222(c *gin.Context) {
	var money model.Money
	var wg sync.WaitGroup //存放Thread的空間，歸0則運行主程式
	var mu sync.Mutex     //宣告互斥鎖

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
		// updateTemp := UpdateTemp{MoneyId: id, Buy: tempBuy, Sell: tempSell}
		// updateTemps = append(updateTemps, updateTemp)
		wg.Add(1)
		// // Channel實作
		// myChan := make(chan int, len(updateTemps))
		// defer close(myChan)

		// for i, _ := range updateTemps {
		// 	go func() {
		// 		// 執行-修改外幣
		// 		err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
		// 		if err != nil {
		// 			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
		// 			return
		// 		}
		// 		myChan <- NowCount

		// 	}()
		// }
		// for i, _ := range updateTemps {
		// 	params := <-myChan
		// 	fmt.Println("B-第", i, "次")
		// 	fmt.Println("B-結果：", params)
		// }

		fmt.Println("變更準備")
		mu.Lock() //互斥鎖 - 鎖住
		fmt.Println("鎖住")
		err = money.UpdateMoneyMarket(id, tempBuy, tempSell)
		if err != nil {
			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
			return
		}
		mu.Unlock() //互斥鎖 - 解鎖
		fmt.Println("解鎖")
		fmt.Println("變更完成")
		wg.Done()
		ShowJsonDATA(c, code.SUCCESS, msg.UPDATE_SUCCESS, "")
		wg.Wait()

	} else {
		// 缺少參數
		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
		return
	}

	// func run(task_id, sleeptime int, ch chan string) {

	// 	time.Sleep(time.Duration(sleeptime) * time.Second)
	// 	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	// 	return
	// }

	// 	input := []int{3, 2, 1}
	//     ch := make(chan string)
	//     fmt.Println("Multirun start")
	//     for i, sleeptime := range input {
	//         go run(i, sleeptime, ch)
	//     }

	//     for range input {
	//         fmt.Println(<-ch)
	// 	}

	// ////----------

	// 	input := []int{3, 2, 1}
	//     chs := make([]chan string, len(input))
	//     startTime := time.Now()
	//     fmt.Println("Multirun start")
	//     for i, sleeptime := range input {
	//         chs[i] = make(chan string)
	//         go run(i, sleeptime, chs[i])
	//     }

	//     for _, ch := range chs {
	//         fmt.Println(<-ch)
	//     }

}
