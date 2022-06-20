package dba

import (
	"github.com/bwmarrin/snowflake"
	"github.com/jmoiron/sqlx"
	"service/models"
)

func Select_loads(db *sqlx.DB) []models.GetLoad {
	var loads []models.GetLoad

	db.Select(&loads,
		`select id, order_number, load_date, loads,
                load_ton, business_number,company_code, lading_bill_number
            from shipping_order
            left join load
            on num = order_number
            order by
            load_date ASC,
            loads DESC`)

	return loads
}

func Update_load(db *sqlx.DB, load models.PutLoad) {
	db.NamedExec(
		`update load set (load_date, loads, load_ton)
            = (:load_date, :loads, :load_ton)
            where id = :id`,
		load)
}

func Insert_shipping_order(db *sqlx.DB, shipping_order models.PostShippingOrder, sfid snowflake.ID) {
	var order_number int32

	db.QueryRow(
		`insert into shipping_order (company_code, lading_bill_number)
            values ($1, $2) returning num`, shipping_order.Company_code, sfid).
		Scan(&order_number)

	for i := 0; i < len(shipping_order.Loads); i++ {
		shipping_order.Loads[i].Order_number = order_number
	}

	db.NamedExec(
		`insert into load (order_number, load_date, loads, load_ton)
            values (:order_number, :load_date, :loads, :load_ton)`,
		shipping_order.Loads)
}

func Delete_load(db *sqlx.DB, id int32) {
	db.Exec(
		`delete from load where id = $1
            returning order_number`,
		id)

	// 然后触发器 trigger_clean_shipping_order 会清理空出货订单
}
