create or replace type cargo as
(
   load_date date,
   loads int4,
   load_ton float8
)

-- 用于调试
create or replace procedure insert_loads(sfid int8, variadic cargoes cargo[])
language plpgsql as
$$
declare
    c cargo;
begin
    insert into shipping_order (lading_bill_number) values (sfid)
        returning num as shipping_order_num;
    foreach c in array cargoes
    loop
        insert into load (order_number, load_date, loads, load_ton)
            values (shipping_order_num, c.load_date, c.loads, c.load_ton);
    end loop;
    return;
end;
$$
