package dba

import "github.com/jmoiron/sqlx"

func Schema(db *sqlx.DB) {
	db.MustExec(`
        create table if not exists warehouse
        (
            house_name text primary key,                -- 场地名称
            address text not NULL,                      -- 地址
            area int4 not NULL                          -- 面积(平方米)
        );

        create table if not exists store
        (
            id serial primary key,
            store_date date not NULL,                   -- 入库日期
            license_plate_number char(7) not NULL,      -- 车牌号
            stocks int not NULL,                        -- 入库件数
            store_ton float8 not NULL,                  -- 入库吨数
            duration int default 0 not NULL             -- 存放天数，每日0点更新
        );

        create table if not exists freight_forwarder
        (
            code text primary key,                      -- 识别码
            company_name text not NULL,                 -- 名字
            telephone_number char(11) not NULL          -- 联系方式
        );

        create table if not exists shipping_order
        (
            num serial4 primary key,                    -- 订单号
            business_number serial4,                    -- 业务号
            company_code text,                          -- 公司识别码
            lading_bill_number int8 not NULL,           -- 提单号
            --------------------------------------------
            constraint company_code_fk foreign key (company_code)
                references freight_forwarder(code)
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
        );
        `)
}
