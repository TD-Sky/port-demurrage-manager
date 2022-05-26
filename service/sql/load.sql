drop table if exists load;

create table load
(
    id serial primary key,
    order_number varchar(20)                    -- 订单号
        check (order_number like 'PO# %')
        not NULL,
    load_date date not NULL,                    -- 开船日期
    loads int not NULL,                         -- 出库件数
    load_ton float8 not NULL,                   -- 出库吨数
    --------------------------------------------
    constraint order_number_fk foreign key (order_number)
        references shipping_order(num)
);

create index load_date_index on load (load_date);
