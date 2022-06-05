package models

type PostStore struct {
	Store_date           string  `json:"id" db:"id"`
	License_plate_number string  `json:"store_date" db:"store_date"`
	Stocks               uint32  `json:"license_plate_number" db:"license_plate_number"`
	Store_ton            float64 `json:"stocks" db:"stocks"`
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
	Id                   uint32  `json:"id" db:"id"`
	Store_date           string  `json:"store_date" db:"store_date"`
	License_plate_number string  `json:"license_plate_number" db:"license_plate_number"`
	Stocks               uint32  `json:"stocks" db:"stocks"`
	Store_ton            float64 `json:"store_ton" db:"store_ton"`
}
