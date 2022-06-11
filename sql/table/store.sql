drop table if exists store;

create table store
(
    id serial primary key,
    store_date date not NULL,                   -- 入库日期
    license_plate_number char(7) not NULL,      -- 车牌号
    stocks int not NULL,                        -- 入库件数
    store_ton float8 not NULL,                  -- 入库吨数
    duration int default 0 not NULL,            -- 存放天数，每日0点更新
    fee float8 default 0 not NULL               -- 这批货物产生的费用
);
