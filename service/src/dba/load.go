package dba

import (
	"github.com/jmoiron/sqlx"
	"service/models"
)

func Select_loads(db *sqlx.DB, loads *[]models.GetLoad) {
	db.Select(loads,
        `select id, order_number, load_date, loads, load_ton, business_number, lading_bill_number
            from shipping_order
            left join load
            on num = order_number
            order by load_date ASC`)
}

func Update_load(db *sqlx.DB, load models.PutLoad) {
	db.NamedExec(
		`update load set (load_date, loads, load_ton)
            = (:load_date, :loads, :load_ton)
            where id = :id`,
		load)
}
