package dba

import (
	"github.com/jmoiron/sqlx"
	"service/models"
)

func Select_warehouses(db *sqlx.DB) []models.Warehouse {
	var warehouses []models.Warehouse

	db.Select(&warehouses, `select * from warehouse`)

	return warehouses
}

func Insert_warehouse(db *sqlx.DB, warehouse models.Warehouse) {
	db.NamedExec(
		`insert into warehouse (house_name, address, area)
            values (:house_name, :address, :area)`,
		warehouse)
}

func Update_warehouse(db *sqlx.DB, warehouse models.Warehouse) {
	db.NamedExec(
		`update warehouse set (address, area)
            = (:address, :area)
            where house_name = :house_name`,
		warehouse)
}

func Delete_warehouse(db *sqlx.DB, house_name string) {
	db.Exec(`delete from warehouse where house_name = $1`, house_name)
}
