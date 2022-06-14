package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"service/utils"
	"time"
)

/* 获取 */

type GetStore struct {
	ID                   int32     `db:"id"`
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               int32     `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
	Duration             int32     `db:"duration"`
}

func (self GetStore) MarshalJSON() ([]byte, error) {
	store := gin.H{
		"id":                   self.ID,
		"stocks":               self.Stocks,
		"store_ton":            self.Store_ton,
		"store_date":           utils.China_date(self.Store_date),
		"license_plate_number": self.License_plate_number,
		"duration":             self.Duration,
	}

	return json.Marshal(store)
}

/* 新增 */

type PostStore struct {
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               int32     `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
}

type postStore struct {
	Store_date           time.Time `json:"store_date"`
	License_plate_number string    `json:"license_plate_number"`
	Stocks               int32     `json:"stocks"`
}

func (self *PostStore) UnmarshalJSON(data []byte) error {
	var store postStore

	json.Unmarshal(data, &store)

	self.Store_date = store.Store_date.In(utils.CN)
	self.Stocks = store.Stocks
	self.Store_ton = 0.05 * float64(store.Stocks)
	self.License_plate_number = store.License_plate_number

	return nil
}

/* 修改，仅允许操作未计费条目 */

type PutStore struct {
	ID                   int32     `db:"id"`
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               int32     `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
}

type putStore struct {
	ID                   int32     `json:"id"`
	Store_date           time.Time `json:"store_date"`
	License_plate_number string    `json:"license_plate_number"`
	Stocks               int32     `json:"stocks"`
}

func (self *PutStore) UnmarshalJSON(data []byte) error {
	var store putStore

	json.Unmarshal(data, &store)

	self.ID = store.ID
	self.Store_date = store.Store_date.In(utils.CN)
	self.Stocks = store.Stocks
	self.Store_ton = 0.05 * float64(store.Stocks)
	self.License_plate_number = store.License_plate_number

	return nil
}

/* 按日划分，用于预测的计算 */

type DayStore struct {
	Store_date time.Time `db:"store_date"`
	Stocks     int32     `db:"stocks"`
	Store_ton  float64   `db:"store_ton"`
	Duration   int32     `db:"duration"`
}
