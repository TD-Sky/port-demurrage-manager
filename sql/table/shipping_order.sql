drop table if exists shipping_order;

create table shipping_order
(
    num serial4 primary key,                -- 订单号
    business_number serial4,                -- 业务号
    lading_bill_number int8 not NULL        -- 提单号
);
