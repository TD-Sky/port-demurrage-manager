package dba

import (
	"github.com/jmoiron/sqlx"
)

func Demurrage_updater(db *sqlx.DB) func() {
	return func() {
		// 先增加库存时间
		db.Exec(`update store set duration
                    = duration + 1
                    where store_date <= CURRENT_DATE`)

		// 再计费
		db.Exec(`update store set fee
                    = fee + 0.5 * store_ton
                    where duration > 14`)
	}
}

// 前提：库存量满足出货量
func Shipment(db *sqlx.DB) func() {
	return func() {
		var total_load_ton float64

		err := db.QueryRow(
			`select sum(load_ton) from load
                where load_date = CURRENT_DATE`).
			Scan(&total_load_ton)

		if err != nil {
			return
		}

		rows, _ := db.Query(`select id, store_ton from store
                                where store_date <= CURRENT_DATE`)
		defer rows.Close()

		for rows.Next() {
			var id int32
			var store_ton float64

			rows.Scan(&id, &store_ton)

			total_load_ton -= store_ton

			if total_load_ton >= 0 {
				db.Exec(`delete from store where id = $1`, id)
			} else {
				db.Exec(`update store set store_ton
                            = store_ton + $1
                            where id = $2`, store_ton, id)
				break // 直至完全出货
			}
		}
	}
}
