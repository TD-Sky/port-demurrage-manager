package models

import (
	"encoding/json"
	"fmt"
	"service/utils"
	"time"
)

/* 真实结构体 */

type GetLoad struct {
	ID                 int32     `db:"id"`
	Order_number       int32     `db:"order_number"`
	Load_date          time.Time `db:"load_date"`
	Loads              int32     `db:"loads"`
	Load_ton           float64   `db:"load_ton"`
	Business_number    int32     `db:"business_number"`
	Lading_bill_number int64     `db:"lading_bill_number"`
}

type PutLoad struct {
	ID        int32     `db:"id"`
	Load_date time.Time `db:"load_date"`
	Loads     int32     `db:"loads"`
	Load_ton  float64   `db:"load_ton"`
}

type PostLoad struct {
	Order_number int32     `db:"order_number"`
	Load_date    time.Time `db:"load_date"`
	Loads        int32     `db:"loads"`
	Load_ton     float64   `db:"load_ton"`
}

// 用于 gin 的参数绑定
type PostLoads struct {
	Loads []PostLoad `json:"loads"`
}

/* 序列化与反序列化的过渡结构体 */

type getLoad struct {
	ID                 int32   `json:"id"`
	Order_number       string  `json:"order_number"`
	Load_date          string  `json:"load_date"`
	Loads              int32   `json:"loads"`
	Load_ton           float64 `json:"load_ton"`
	Business_number    string  `json:"business_number"`
	Lading_bill_number int64   `json:"lading_bill_number"`
}

type putLoad struct {
	ID        int32   `json:"id"`
	Load_date string  `json:"load_date"`
	Loads     int32   `json:"loads"`
	Load_ton  float64 `json:"load_ton"`
}

type postLoad struct {
	Load_date string  `json:"load_date"`
	Loads     int32   `json:"loads"`
	Load_ton  float64 `json:"load_ton"`
}

func (self GetLoad) MarshalJSON() ([]byte, error) {
	load := getLoad{
		ID:                 self.ID,
		Order_number:       fmt.Sprintf("PO# %07d", self.Order_number),
		Loads:              self.Loads,
		Load_ton:           self.Load_ton,
		Load_date:          utils.Date_string(&self.Load_date),
		Business_number:    fmt.Sprintf("LLH%07d", self.Business_number),
		Lading_bill_number: self.Lading_bill_number,
	}

	return json.Marshal(load)
}

/* func (self *GetLoad) UnmarshalJSON(data []byte) error {
	var load getLoad

	json.Unmarshal(data, &load)

	self.ID = load.ID
	self.Order_number = utils.Parse_order_number(load.Order_number)
	self.Load_date = utils.Parse_date(load.Load_date)
	self.Loads = load.Loads
	self.Load_ton = load.Load_ton
	self.Business_number = utils.Parse_business_number(load.Business_number)
	self.Lading_bill_number = load.Lading_bill_number

	return nil
} */

/* func (self PutLoad) MarshalJSON() ([]byte, error) {
	load := putLoad{
		ID:        self.ID,
		Loads:     self.Loads,
		Load_ton:  self.Load_ton,
		Load_date: utils.Date_string(&self.Load_date),
	}

	return json.Marshal(load)
} */

func (self *PutLoad) UnmarshalJSON(data []byte) error {
	var load putLoad

	json.Unmarshal(data, &load)

	self.ID = load.ID
	self.Load_date = utils.Parse_date(load.Load_date)
	self.Loads = load.Loads
	self.Load_ton = load.Load_ton

	return nil
}

/* func (self PostLoad) MarshalJSON() ([]byte, error) {
	load := postLoad{
		Loads:     self.Loads,
		Load_ton:  self.Load_ton,
		Load_date: utils.Date_string(&self.Load_date),
	}

	return json.Marshal(load)
} */

func (self *PostLoad) UnmarshalJSON(data []byte) error {
	var load postLoad

	json.Unmarshal(data, &load)

	self.Order_number = 0
	self.Load_date = utils.Parse_date(load.Load_date)
	self.Loads = load.Loads
	self.Load_ton = load.Load_ton

	return nil
}
