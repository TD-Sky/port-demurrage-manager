package dba

import (
	"github.com/bwmarrin/snowflake"
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

func Insert_loads(db *sqlx.DB, loads []models.PostLoad, sfid snowflake.ID) {
	var order_number int32

	db.QueryRow(
		`insert into shipping_order (lading_bill_number)
            values ($1) returning num`, sfid).
		Scan(&order_number)

	for i := 0; i < len(loads); i++ {
		loads[i].Order_number = order_number
	}

	db.NamedExec(
		`insert into load (order_number, load_date, loads, load_ton)
            values (:order_number, :load_date, :loads, :load_ton)`,
		loads)
}

func Delete_load(db *sqlx.DB, id int32) {
	db.Exec(
		`delete from load where id = $1
            returning order_number`,
		id)

    // 然后触发器 trigger_clean_shipping_order 会清理空出货订单
}
