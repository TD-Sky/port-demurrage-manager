package models

type PostStore struct {
    store_date string
    license_plate_number string
    stocks uint32
    store_ton float64
}

type GetStore struct {
    id uint32
    store_date string
    license_plate_number string
    stocks uint32
    store_ton float64
    duration uint32
    fee uint32
}

type PutStore struct {
    store_date *string
    license_plate_number *string
    stocks *uint32
    store_ton *float64
}
