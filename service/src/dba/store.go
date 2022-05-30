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

func Select_stores(db *sqlx.DB, stores *[]models.GetStore) {
	db.Select(stores,
		`select * from store
         order by store_date ASC`)
}

func Update_store(db *sqlx.DB, store models.PutStore) {
	db.NamedExec(
		`update store set (store_date, license_plate_number, stocks, store_ton)
            = (:store_date, :license_plate_number, :stocks, :store_ton)
            where id = :id`,
		store)
}
