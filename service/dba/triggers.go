package dba

import "github.com/jmoiron/sqlx"

func Triggers(db *sqlx.DB) {
	clean_shipping_order :=
		`-- 钩子函数
            create or replace function clean_shipping_order() returns trigger as
            $$
            begin
            if (select count(*) from load where order_number = OLD.order_number) = 0 then
                delete from shipping_order where num = OLD.order_number;
            end if;
            return NULL;
            end;
            $$
            language plpgsql;



            -- 每删除一行出库数据，都要检查订单是否还有出库项；
            -- 没有则删除出货单
            create or replace trigger trigger_clean_shipping_order after delete
            on load
            for each row
            execute function clean_shipping_order();`

	db.MustExec(clean_shipping_order)
}
