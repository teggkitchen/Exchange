package models

import (
	"errors"
	msg "exchange/config"
	configDB "exchange/database"
	"fmt"
	"time"
)

type Money struct {
	Id        int64     `json:"id"`        // 外幣Id
	Name      string    `json:"name"`      // 外幣名稱
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間
}

type CurrentMarket struct {
	Id        int64     `json:"id"`        // 當前外幣Id
	MoneyId   int64     `json:"moneyId"`   // 外幣Id
	Buy       float64   `json:"buy"`       // 買入
	Sell      float64   `json:"sell"`      // 賣出
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間
}

type HistoricalMarket struct {
	Id        int64     `json:"id"`        // 當前外幣Id
	MoneyId   int64     `json:"moneyId"`   // 外幣Id
	Buy       float64   `json:"buy"`       // 買入
	Sell      float64   `json:"sell"`      // 賣出
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間
}

// 新增外幣
// 比對幣別名稱重複
func (money *Money) CheckMoneyName(name string) (err error) {
	fmt.Println(name)
	configDB.GormOpen.Debug().Table("moneys").Where("name=?", name).Scan(&money)

	fmt.Println(money.Name)
	if len(money.Name) > 0 {
		return errors.New(msg.DATA_REPEAT_ERROR)
	}
	return nil
}

// 新增幣別名稱
func (money *Money) InsertMoneyName(name string, buy float64, sell float64) (err error) {
	money.Name = name
	result := configDB.GormOpen.Table("moneys").Create(&money)
	if result.Error != nil {
		err = result.Error
		return err
	}

	if err := InsertMoneyMarket(money.Id, buy, sell); err != nil {
		return errors.New(msg.SQL_WRITE_ERROR)
	}
	return nil
}

// 新增幣別行情
func InsertMoneyMarket(id int64, buy float64, sell float64) (err error) {
	var currentMarket CurrentMarket
	var historicalMarket HistoricalMarket

	//當前行情
	currentMarket.MoneyId = id
	currentMarket.Buy = buy
	currentMarket.Sell = sell
	currentMarket.CreatedAt = time.Now()
	currentMarket.UpdatedAt = time.Now()

	//歷史行情
	historicalMarket.MoneyId = id
	historicalMarket.Buy = buy
	historicalMarket.Sell = sell
	historicalMarket.CreatedAt = time.Now()
	historicalMarket.UpdatedAt = time.Now()

	result := configDB.GormOpen.Table("current_markets").Create(&currentMarket)
	if result.Error != nil {
		err = result.Error
		return err
	}

	result = configDB.GormOpen.Table("historical_markets").Create(&historicalMarket)
	if result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}

// 修改外幣
// 比對買賣是否重複
func (money *Money) CheckMoneyMarket(id int64, buy float64, sell float64) (isRepeat bool) {
	var currentMarket CurrentMarket
	// var tempBuy float64
	// var tempSell float64

	fmt.Println(buy)
	fmt.Println(sell)
	// configDB.GormOpen.Debug().Table("current_markets").Where("buy=? AND sell=?", buy, sell).Scan(&currentMarket)
	// configDB.GormOpen.Debug().Table("current_markets").Select([]string{"buy"}).Where("buy=? AND sell=?", buy, sell).Scan(&tempBuy)
	configDB.GormOpen.Debug().Table("current_markets").Where("buy=? AND sell=?", buy, sell).Scan(&currentMarket)
	tempBuy := currentMarket.Buy
	tempSell := currentMarket.Sell
	// configDB.GormOpen.Debug().Table("current_markets").Where("sell='?'", sell).Scan(tempSell)
	// fff := currentMarket.buy
	// ppp := strconv.FormatFloat(fff, 'f', 6, 64)
	// fmt.Println(ppp)
	// fmt.Println(currentMarket.sell)

	if tempBuy != 0 && tempSell != 0 {
		return true
	}
	return false
}
func (money *Money) UpdateMoneyMarket(id int64, buy float64, sell float64) (err error) {
	var currentMarket CurrentMarket
	var historicalMarket HistoricalMarket

	//當前行情
	currentMarket.Buy = buy
	currentMarket.Sell = sell
	currentMarket.UpdatedAt = time.Now()

	//歷史行情
	historicalMarket.MoneyId = id
	historicalMarket.Buy = buy
	historicalMarket.Sell = sell
	historicalMarket.CreatedAt = time.Now()
	historicalMarket.UpdatedAt = time.Now()

	// if err = configDB.GormOpen.Debug().Table("current_markets").Select([]string{"money_id"}).First(&currentMarket, id).Error; err != nil {
	// 	return err
	// }

	if err = configDB.GormOpen.Debug().Table("current_markets").Where("money_id=?", id).Model(&currentMarket).Updates(&currentMarket).Error; err != nil {
		return err
	}

	result := configDB.GormOpen.Debug().Table("historical_markets").Create(&historicalMarket)
	if result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}

// 刪除外幣
func (Money *Money) DestroyMoneyMarket(id int64) (err error) {

	// if err = configDB.GormOpen.Table("moneys").Select([]string{"id"}).First(&Money, id).Error; err != nil {
	// 	return err
	// }

	// if err = configDB.GormOpen.Table("moneys").Delete(&Money).Error; err != nil {
	// 	return err
	// }

	if err = configDB.GormOpen.Debug().Table("moneys").Where("id=?", id).Delete(&Money).Error; err != nil {
		return err
	}
	if err = configDB.GormOpen.Debug().Table("current_markets").Where("id=?", id).Delete(&Money).Error; err != nil {
		return err
	}
	if err = configDB.GormOpen.Debug().Table("historical_markets").Where("id=?", id).Delete(&Money).Error; err != nil {
		return err
	}

	return nil
}

// 查詢全部
func (money *Money) QueryMoneys() (data interface{}, err error) {
	var moneys []Money
	result := configDB.GormOpen.Table("moneys").Find(&moneys)
	if result.Error != nil {
		err = result.Error
		return "", err
	} else if len(moneys[0:]) == 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}
	return moneys, nil
}
