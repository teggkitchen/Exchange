package models

import (
	"errors"
	msg "exchange/config"
	configDB "exchange/database"
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

// 查詢外幣查詢外幣
type MoneyMarket struct {
	MoneyId   int64     `json:"moneyId"`   // 外幣Id
	Name      string    `json:"name"`      // 外幣名稱
	Buy       float64   `json:"buy"`       // 買入
	Sell      float64   `json:"sell"`      // 賣出
	CreatedAt time.Time `json:"createdAt"` // 開始時間
	UpdatedAt time.Time `json:"updatedAt"` // 更新時間
}

// 新增外幣
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

	//當前幣別行情
	currentMarket.MoneyId = id
	currentMarket.Buy = buy
	currentMarket.Sell = sell
	currentMarket.CreatedAt = time.Now()
	currentMarket.UpdatedAt = time.Now()

	//歷史幣別行情
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

// 方法 - 比對幣別名稱重複
func (money *Money) CheckMoneyName(name string) (err error) {
	configDB.GormOpen.Debug().Table("moneys").Where("name=?", name).Scan(&money)
	if len(money.Name) > 0 {
		return errors.New(msg.DATA_REPEAT_ERROR)
	}
	return nil
}

// 修改外幣
//修改幣別行情
func (money *Money) UpdateMoneyMarket(id int64, buy float64, sell float64) (err error) {
	var currentMarket CurrentMarket
	var historicalMarket HistoricalMarket

	//當前幣別行情
	currentMarket.Buy = buy
	currentMarket.Sell = sell
	currentMarket.UpdatedAt = time.Now()

	//歷史幣別行情
	historicalMarket.MoneyId = id
	historicalMarket.Buy = buy
	historicalMarket.Sell = sell
	historicalMarket.CreatedAt = time.Now()
	historicalMarket.UpdatedAt = time.Now()

	//判斷是否有此id
	if err = configDB.GormOpen.Debug().Table("current_markets").Select([]string{"money_id"}).First(&currentMarket, id).Error; err != nil {
		return err
	}

	//修改 當前幣別行情
	if err = configDB.GormOpen.Debug().Table("current_markets").Where("money_id=?", id).Model(&currentMarket).Updates(&currentMarket).Error; err != nil {
		return err
	}

	//增加 歷史幣別行情
	result := configDB.GormOpen.Debug().Table("historical_markets").Create(&historicalMarket)
	if result.Error != nil {
		err = result.Error
		return err
	}
	return nil
}

//修改幣別行情
func (money *Money) UpdateMoneyMarketTemp(id int64, buy float64, sell float64) (redId int64, err error) {
	var currentMarket CurrentMarket
	var historicalMarket HistoricalMarket

	//當前幣別行情
	currentMarket.Buy = buy
	currentMarket.Sell = sell
	currentMarket.UpdatedAt = time.Now()

	//歷史幣別行情
	historicalMarket.MoneyId = id
	historicalMarket.Buy = buy
	historicalMarket.Sell = sell
	historicalMarket.CreatedAt = time.Now()
	historicalMarket.UpdatedAt = time.Now()

	//判斷是否有此id
	if err = configDB.GormOpen.Debug().Table("current_markets").Select([]string{"money_id"}).First(&currentMarket, id).Error; err != nil {
		return -1, err
	}

	//修改 當前幣別行情
	if err = configDB.GormOpen.Debug().Table("current_markets").Where("money_id=?", id).Model(&currentMarket).Updates(&currentMarket).Error; err != nil {
		return -1, err
	}

	//增加 歷史幣別行情
	result := configDB.GormOpen.Debug().Table("historical_markets").Create(&historicalMarket)
	if result.Error != nil {
		err = result.Error
		return -1, err
	}
	return historicalMarket.MoneyId, nil
}

// 方法 - 比對買賣是否重複
func (money *Money) IsCheckMoneyMarket(id int64, buy float64, sell float64) (isRepeat bool) {
	var currentMarket CurrentMarket
	configDB.GormOpen.Debug().Table("current_markets").Where("buy=? AND sell=?", buy, sell).Scan(&currentMarket)
	tempBuy := currentMarket.Buy
	tempSell := currentMarket.Sell
	if tempBuy != 0 && tempSell != 0 {
		return true
	}
	return false
}

// 刪除外幣
func (Money *Money) DestroyMoneyMarket(id int64) (err error) {

	//刪除 幣別名稱
	if err = configDB.GormOpen.Debug().Table("moneys").Where("id=?", id).Delete(&Money).Error; err != nil {
		return err
	}

	//刪除 當前幣別行情
	if err = configDB.GormOpen.Debug().Table("current_markets").Where("id=?", id).Delete(&CurrentMarket{}).Error; err != nil {
		return err
	}

	//刪除 歷史幣別行情（批次）
	if err = configDB.GormOpen.Debug().Table("historical_markets").Where(HistoricalMarket{MoneyId: +id}).Delete(&HistoricalMarket{}).Error; err != nil {
		return err
	}

	return nil
}

// 查詢全部
func (money *Money) QueryMoneys() (data interface{}, err error) {
	var moneys []Money
	var currentMarket []CurrentMarket
	var moneyMarkets []MoneyMarket
	result := configDB.GormOpen.Debug().Table("moneys").Find(&moneys)
	if result.Error != nil {
		err = result.Error
		return "", err
	} else if len(moneys[0:]) == 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}

	result = configDB.GormOpen.Debug().Table("current_markets").Find(&currentMarket)
	if result.Error != nil {
		err = result.Error
		return "", err
	} else if len(currentMarket[0:]) == 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}

	for _, m := range moneys {
		for _, c := range currentMarket {
			if m.Id == c.MoneyId {
				var moneyMarket MoneyMarket
				moneyMarket.Name = m.Name
				moneyMarket.MoneyId = c.MoneyId
				moneyMarket.Buy = c.Buy
				moneyMarket.Sell = c.Sell
				moneyMarket.CreatedAt = c.CreatedAt
				moneyMarket.UpdatedAt = c.UpdatedAt
				moneyMarkets = append(moneyMarkets, moneyMarket)
			}

		}

	}
	return moneyMarkets, nil
}

// 查詢單一幣
func (money *Money) QueryMoney(moneyId int64) (data interface{}, err error) {
	var moneyMarket MoneyMarket
	var currentMarket CurrentMarket

	result := configDB.GormOpen.Table("current_markets").First(&currentMarket, moneyId)
	if result.Error != nil {
		err = result.Error
		return nil, err
	} else if currentMarket.MoneyId < 0 {
		return nil, errors.New(msg.NOT_FOUND_DATA_ERROR)
	}

	configDB.GormOpen.Table("moneys").Find(&money, moneyId)
	moneyMarket.MoneyId = currentMarket.MoneyId
	moneyMarket.Name = money.Name
	moneyMarket.Buy = currentMarket.Buy
	moneyMarket.Sell = currentMarket.Sell
	moneyMarket.CreatedAt = currentMarket.CreatedAt
	moneyMarket.UpdatedAt = currentMarket.UpdatedAt

	return moneyMarket, nil
}
