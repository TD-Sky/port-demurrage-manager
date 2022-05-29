package models

type PostStore struct {
    Store_date           string  `form:"入库日期" db:"store_date"`
    License_plate_number string  `form:"车牌号" db:"license_plate_number"`
    Stocks               uint32  `form:"件数" db:"stocks"`
    Store_ton            float64 `form:"吨数" db:"store_ton"`
}

type GetStore struct {
	id                   uint32
	store_date           string
	license_plate_number string
	stocks               uint32
	store_ton            float64
	duration             uint32
	fee                  uint32
}

type PutStore struct {
	store_date           *string
	license_plate_number *string
	stocks               *uint32
	store_ton            *float64
}
