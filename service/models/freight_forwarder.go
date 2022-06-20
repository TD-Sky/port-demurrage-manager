package models

type FreightForwarder struct {
	Code             string `db:"code" json:"code"`
	Company_name     string `db:"company_name" json:"company_name"`
	Telephone_number string `db:"telephone_number" json:"telephone_number"`
}
