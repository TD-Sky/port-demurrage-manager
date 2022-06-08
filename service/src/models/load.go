package models

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"service/utils"
	"time"
)

/* 获取 */

type GetLoad struct {
	ID                 int32     `db:"id"`
	Order_number       int32     `db:"order_number"`
	Load_date          time.Time `db:"load_date"`
	Loads              int32     `db:"loads"`
	Load_ton           float64   `db:"load_ton"`
	Business_number    int32     `db:"business_number"`
	Lading_bill_number int64     `db:"lading_bill_number"`
}

func (self GetLoad) MarshalJSON() ([]byte, error) {
	load := gin.H{
		"id":                 self.ID,
		"order_number":       fmt.Sprintf("PO# %07d", self.Order_number),
		"loads":              self.Loads,
		"load_ton":           self.Load_ton,
		"load_date":          utils.CST(&self.Load_date).Format("2006-01-02"),
		"business_number":    fmt.Sprintf("LLH%07d", self.Business_number),
		"lading_bill_number": self.Lading_bill_number,
	}

	return json.Marshal(load)
}

/* 新增 */

type PostLoad struct {
	Order_number int32     `db:"order_number"`
	Load_date    time.Time `db:"load_date"`
	Loads        int32     `db:"loads"`
	Load_ton     float64   `db:"load_ton"`
}

type postLoad struct {
	Load_date time.Time `json:"load_date"`  // 日期类型的作用不可替代
	Loads     int32     `json:"loads"`
	Load_ton  float64   `json:"load_ton"`
}

func (self *PostLoad) UnmarshalJSON(data []byte) error {
	var load postLoad

	json.Unmarshal(data, &load)

	self.Order_number = 0
	self.Load_date = utils.CST(&load.Load_date)
	self.Loads = load.Loads
	self.Load_ton = load.Load_ton

	return nil
}

/* 修改 */

type PutLoad struct {
	ID        int32     `db:"id"`
	Load_date time.Time `db:"load_date"`
	Loads     int32     `db:"loads"`
	Load_ton  float64   `db:"load_ton"`
}

type putLoad struct {
	ID        int32     `json:"id"`
	Load_date time.Time `json:"load_date"`  // 日期类型的作用不可替代
	Loads     int32     `json:"loads"`
	Load_ton  float64   `json:"load_ton"`
}

func (self *PutLoad) UnmarshalJSON(data []byte) error {
	var load putLoad

	json.Unmarshal(data, &load)

	self.ID = load.ID
	self.Load_date = utils.CST(&load.Load_date)
	self.Loads = load.Loads
	self.Load_ton = load.Load_ton

	return nil
}
