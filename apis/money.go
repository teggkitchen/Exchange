package apis

import (
	code "exchange/config"
	msg "exchange/config"
	model "exchange/models"
	. "exchange/utils"
	"fmt"
	"os"
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

// // 展示全部商品
// func ShowProducts(c *gin.Context) {
// 	var product model.Product
// 	// 執行-查詢全部商品
// 	result, err := product.QueryProducts()

// 	if err != nil {
// 		ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_DATA_ERROR)
// 		return
// 	}
// 	ShowJsonDATA(c, code.SUCCESS, msg.EXEC_SUCCESS, result)

// }

// 增加商品
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

	// file, header, err := c.Request.FormFile("productImage")
	// if err != nil {
	// 	ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
	// 	return
	// }

	//參數是否有值
	if len(tempNameStr) > 0 && tempBuy > 0 && tempSell > 0 {

		// 執行-增加商品
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

// // 修改商品
// func UpdateProduct(c *gin.Context) {
// 	var product model.Product
// 	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	product.Name = c.Request.FormValue("productName")
// 	product.Price = c.Request.FormValue("productPrice")
// 	product.UpdatedAt = time.Now()

// 	file, header, err := c.Request.FormFile("productImage")
// 	if err != nil {
// 		ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
// 		return
// 	}

// 	//參數是否有值
// 	if len(product.Name) > 0 && len(product.Price) > 0 {
// 		filename := header.Filename
// 		if file == nil && len(filename) <= 0 {
// 			//找不到圖片
// 			ShowJsonMSG(c, code.ERROR, msg.NOT_FOUND_IMAGE)
// 			return
// 		}

// 		// 執行-查詢原本圖檔的名稱
// 		oldImgName, err := product.GetProductImg(id)
// 		if err != nil {
// 			//找不到圖片刪除，仍繼續
// 			fmt.Println(err)
// 		}

// 		//圖檔重新命名
// 		product.Img = fileRename(filename)

// 		// 執行-修改商品
// 		err = product.Update(id)
// 		if err != nil {
// 			//如果出錯，就刪除剛存的圖片
// 			os.Remove(productImgPath + product.Img)
// 			ShowJsonMSG(c, code.ERROR, msg.WRITE_ERROR)
// 			return
// 		}

// 		//刪除原本照片
// 		err = os.Remove(productImgPath + oldImgName)
// 		if err != nil {
// 			//找不到圖片刪除，仍繼續
// 			fmt.Println(msg.CONTINUE_NOT_FOUND_IMAGE)
// 		}

// 		// 新增圖片
// 		err = addImg(c, file, product.Img, productImgPath)
// 		if err != nil {
// 			ShowJsonMSG(c, code.ERROR, msg.ADD_IMAGE_ERROR)

// 		}

// 		ShowJsonDATA(c, code.SUCCESS, msg.UPDATE_SUCCESS, "")

// 	} else {
// 		// 缺少參數
// 		ShowJsonMSG(c, code.ERROR, msg.ARGS_ERROR)
// 		return
// 	}

// }

//刪除商品
func DestroyMoney(c *gin.Context) {
	var product model.Product
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	// 執行-刪除商品
	err = product.DestroyMoneyMarket(id)
	if err != nil {
		//刪除失敗
		ShowJsonMSG(c, code.ERROR, msg.DELETE_ERROR)
		return
	}

	// 刪除原本照片
	err = os.Remove(productImgPath + oldImgName)
	if err != nil {
		// 找不到圖片刪除，仍繼續
		fmt.Println(msg.CONTINUE_NOT_FOUND_IMAGE)
	}

	ShowJsonDATA(c, code.SUCCESS, msg.DELETE_SUCCESS, "")

}

// func fileRename(filename string) string {
// 	// 替換圖片檔名
// 	newFileName := GetMD5Hash(filename + time.Now().String())
// 	dotIndex := strings.LastIndex(filename, ".") //取得最後的.的索引值
// 	if dotIndex != -1 && dotIndex != 0 {
// 		newFileName += filename[dotIndex:] //取得副檔名
// 	}
// 	//輸出 檔名＋副檔名

// 	return newFileName
// }

// func addImg(c *gin.Context, file io.Reader, fileName string, filePath string) error {
// 	//  判斷資料夾
// 	if !IsExists(filePath) {
// 		// 不存在
// 		os.Mkdir(filePath, os.ModePerm)
// 		fmt.Println("創建資料夾路徑為：" + filePath)
// 	}

// 	//抓取新圖片到指定目錄
// 	out, err := os.Create(filePath + fileName)
// 	if err != nil {
// 		//沒有image資料夾
// 		return errors.New(msg.NOT_FOUND_IMAGE_FOLDER)
// 	}
// 	defer out.Close()
// 	_, err = io.Copy(out, file)
// 	if err != nil {
// 		//寫入檔案失敗
// 		return errors.New(msg.WRITE_FILE_ERROR)
// 	}
// 	return nil
// }
