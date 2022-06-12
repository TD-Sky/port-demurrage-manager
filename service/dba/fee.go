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
		var total_loads int32
		var total_load_ton float64

		err := db.QueryRow(
			`select sum(loads), sum(load_ton) from load
                where load_date = CURRENT_DATE`).
			Scan(&total_loads, &total_load_ton)

		// 今日无出货业务，终止函数
		if err != nil {
			return
		}

		rows, _ := db.Query(`select id, stocks, store_ton
                                from store
                                where store_date <= CURRENT_DATE
                                order by
                                store_date ASC, stocks DESC`)
		defer rows.Close()

		for rows.Next() {
			var id int32
			var store_ton float64
			var stocks int32

			rows.Scan(&id, &stocks, &store_ton)

			total_loads -= stocks
			total_load_ton -= store_ton

			if total_load_ton >= 0 {
				db.Exec(`delete from store where id = $1`, id)
			} else {
				db.Exec(`update store set (stocks, store_ton)
                            = ($1, $2) where id = $3`,
					-total_loads, -total_load_ton, id)

				break // 直至完全出货
			}
		}

		// 出货了，删除出货单
		db.Exec(`delete from load where load_date = CURRENT_DATE`)
	}
}
