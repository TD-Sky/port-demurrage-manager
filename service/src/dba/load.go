package dba

import (
	"service/models"

	"github.com/jmoiron/sqlx"
)

func Select_loads(db *sqlx.DB, loads *[]models.GetLoad) {
	db.Select(loads,
		`select id, order_number, load_date, loads, load_ton, business_number, lading_bill_number
         from shipping_order
         left join load
         on num = order_number`)
}
