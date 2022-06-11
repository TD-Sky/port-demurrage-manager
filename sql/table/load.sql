drop table if exists load;

create table load
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
