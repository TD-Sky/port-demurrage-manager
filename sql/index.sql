-- 依据日期顺序处理出入库
create index load_date_index on load (load_date);
create index store_date_index on store (store_date);
