package models

type GetLoad struct {
	ID                 uint32  `json:"id" db:"id"`
	Order_number       string  `json:"order_number" db:"order_number"`
	Load_date          string  `json:"load_date" db:"load_date"`
	Loads              uint32  `json:"loads" db:"loads"`
	Load_ton           float64 `json:"load_ton" db:"load_ton"`
	Business_number    string  `json:"business_number" db:"business_number"`
	Lading_bill_number string  `json:"lading_bill_number" db:"lading_bill_number"`
}

type PutLoad struct {
	ID                 uint32  `json:"id" db:"id"`
	Load_date          string  `json:"load_date" db:"load_date"`
	Loads              uint32  `json:"loads" db:"loads"`
	Load_ton           float64 `json:"load_ton" db:"load_ton"`
}
