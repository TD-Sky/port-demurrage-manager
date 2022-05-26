drop table if exists shipping_order;

create table shipping_order
(
    num varchar(20)                             -- 订单号
        primary key
        check (like 'PO# %'),
    business_number char(10) not NULL           -- 业务号
        not NULL check (like 'LLH%'),
    lading_bill_number varchar(20) not NULL     -- 提单号
);
