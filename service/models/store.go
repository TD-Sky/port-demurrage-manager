package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"service/utils"
	"time"
)

/* 获取 */

type GetStore struct {
	ID                   uint32    `db:"id"`
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               uint32    `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
	Duration             uint32    `db:"duration"`
	Fee                  uint32    `db:"fee"`
}

func (self GetStore) MarshalJSON() ([]byte, error) {
	store := gin.H{
		"id":                   self.ID,
		"stocks":               self.Stocks,
		"store_ton":            self.Store_ton,
		"store_date":           utils.CST(&self.Store_date).Format("2006-01-02"),
		"license_plate_number": self.License_plate_number,
		"duration":             self.Duration,
		"fee":                  self.Fee,
	}

	return json.Marshal(store)
}

/* 新增 */

type PostStore struct {
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               uint32    `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
}

type postStore struct {
	Store_date           time.Time `json:"store_date"`
	License_plate_number string    `json:"license_plate_number"`
	Stocks               uint32    `json:"stocks"`
	Store_ton            float64   `json:"store_ton"`
}

func (self *PostStore) UnmarshalJSON(data []byte) error {
	var store postStore

	json.Unmarshal(data, &store)

	self.Store_date = utils.CST(&store.Store_date)
	self.Stocks = store.Stocks
	self.Store_ton = store.Store_ton
	self.License_plate_number = store.License_plate_number

	return nil
}

/* 修改，仅允许操作未计费条目 */

type PutStore struct {
	ID                   uint32    `db:"id"`
	Store_date           time.Time `db:"store_date"`
	License_plate_number string    `db:"license_plate_number"`
	Stocks               uint32    `db:"stocks"`
	Store_ton            float64   `db:"store_ton"`
}

type putStore struct {
	ID                   uint32    `json:"id"`
	Store_date           time.Time `json:"store_date"`
	License_plate_number string    `json:"license_plate_number"`
	Stocks               uint32    `json:"stocks"`
	Store_ton            float64   `json:"store_ton"`
}

func (self *PutStore) UnmarshalJSON(data []byte) error {
	var store putStore

	json.Unmarshal(data, &store)

	self.ID = store.ID
	self.Store_date = utils.CST(&store.Store_date)
	self.Stocks = store.Stocks
	self.Store_ton = store.Store_ton
	self.License_plate_number = store.License_plate_number

	return nil
}
