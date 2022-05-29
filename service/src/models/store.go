package models

type PostStore struct {
	Store_date           string  `form:"入库日期" db:"store_date"`
	License_plate_number string  `form:"车牌号" db:"license_plate_number"`
	Stocks               uint32  `form:"件数" db:"stocks"`
	Store_ton            float64 `form:"吨数" db:"store_ton"`
}

type GetStore struct {
	Id                   uint32  `json:"id" db:"id"`
	Store_date           string  `json:"store_date" db:"store_date"`
	License_plate_number string  `json:"license_plate_number" db:"license_plate_number"`
	Stocks               uint32  `json:"stocks" db:"stocks"`
	Store_ton            float64 `json:"store_ton" db:"store_ton"`
	Duration             uint32  `json:"duration" db:"duration"`
	Fee                  uint32  `json:"fee" db:"fee"`
}

// 还没有开始计费的入库条目才能更改
type PutStore struct {
	Id                   uint32  `form:"id" db:"id"`
	Store_date           string  `form:"入库日期" db:"store_date"`
	License_plate_number string  `form:"车牌号" db:"license_plate_number"`
	Stocks               uint32  `form:"件数" db:"stocks"`
	Store_ton            float64 `form:"吨数" db:"store_ton"`
}
