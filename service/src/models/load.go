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
	ID        uint32  `json:"id" db:"id"`
	Load_date string  `json:"load_date" db:"load_date"`
	Loads     uint32  `json:"loads" db:"loads"`
	Load_ton  float64 `json:"load_ton" db:"load_ton"`
}

type PostLoad struct {
	Load_date string  `json:"load_date" db:"load_date"`
	Loads     uint32  `json:"loads" db:"loads"`
	Load_ton  float64 `json:"load_ton" db:"load_ton"`
}

/* 序列化与反序列化的过渡结构体 */

type getLoad struct {
	ID                 int32
	Order_number       string
	Load_date          string
	Loads              int32
	Load_ton           float64
	Business_number    string
	Lading_bill_number int64
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

	if data, err := json.Marshal(load); err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return data, nil
	}
}

func (self *GetLoad) UnmarshalJSON(data []byte) error {
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
}
