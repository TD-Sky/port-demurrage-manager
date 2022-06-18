package dba

import "github.com/jmoiron/sqlx"

func Schema(db *sqlx.DB) {
	db.MustExec(
		`create table if not exists store
        (
            id serial primary key,
            store_date date not NULL,                   -- 入库日期
            license_plate_number char(7) not NULL,      -- 车牌号
            stocks int not NULL,                        -- 入库件数
            store_ton float8 not NULL,                  -- 入库吨数
            duration int default 0 not NULL             -- 存放天数，每日0点更新
        );

        create table if not exists shipping_order
        (
            num serial4 primary key,                -- 订单号
            business_number serial4,                -- 业务号
            lading_bill_number int8 not NULL        -- 提单号
        );

        create table if not exists load
        (
            id serial4 primary key,
            order_number int4 not NULL,     -- 订单号
            load_date date not NULL,        -- 开船日期
            loads int4 not NULL,            -- 出库件数
            load_ton float8 not NULL,       -- 出库吨数
            --------------------------------------------
            constraint order_number_fk foreign key (order_number)
                references shipping_order(num)
        );

        create table if not exists login
        (
            username text primary key,      -- 用户名
            passwd text not NULL            -- 密码
        );`)
}
