package models

type Warehouse struct {
	House_name string `db:"house_name" json:"house_name"`
	Address    string `db:"address" json:"address"`
	Area       int32  `db:"area" json:"area"`
}
