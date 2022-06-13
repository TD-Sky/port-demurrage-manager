package dba

import (
	"github.com/jmoiron/sqlx"
	"service/models"
)

func Insert_store(db *sqlx.DB, store models.PostStore) {
	db.NamedExec(
		`insert into store (store_date, license_plate_number, stocks, store_ton)
             values (:store_date, :license_plate_number, :stocks, :store_ton)`,
		store)
}

func Select_stores(db *sqlx.DB) []models.GetStore {
	var stores []models.GetStore

	db.Select(&stores,
		`select * from store
         order by
         store_date ASC, stocks DESC`)

	return stores
}

func Update_store(db *sqlx.DB, store models.PutStore) {
	db.NamedExec(
		`update store set (store_date, license_plate_number, stocks, store_ton)
            = (:store_date, :license_plate_number, :stocks, :store_ton)
            where id = :id`,
		store)
}

func Delete_store(db *sqlx.DB, id int32) {
	db.Exec(`delete from store where id = $1`, id)
}

func Group_stores_by_day(db *sqlx.DB) []models.DayStore {
	var stores []models.DayStore

	db.Select(stores, `select
                        store_date,
                        sum(stocks) as stocks,
                        sum(store_ton) as store_ton,
                        duration,
                        sum(fee) as fee
                       from store
                       group by store_date, duration
                       order by store_date`)
	return stores
}
