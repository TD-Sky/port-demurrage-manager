package dba

import (
	"github.com/jmoiron/sqlx"
	"service/models"
)

func Select_freight_forwarders(db *sqlx.DB) []models.FreightForwarder {
	var companies []models.FreightForwarder

	db.Select(&companies, `select * from freight_forwarder`)

	return companies
}

func Insert_freight_forwarder(db *sqlx.DB, company models.FreightForwarder) {
	db.NamedExec(
		`insert into freight_forwarder (code, company_name, telephone_number)
            values (:code, :company_name, :telephone_number)`,
		company)
}

func Update_freight_forwarder(db *sqlx.DB, company models.FreightForwarder) {
	db.NamedExec(
		`update freight_forwarder set (company_name, telephone_number)
            = (:company_name, :telephone_number)
            where code = :code`,
		company)
}

func Delete_freight_forwarder(db *sqlx.DB, code string) error {
	_, err := db.Exec(`delete from freight_forwarder where code = $1`, code)

	return err // 仍有外键约束，禁止删除
}
