package handlers

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"math"
	"service/dba"
	"service/models"
	"service/utils"
)

func predict(db *sqlx.DB, loads []models.GetLoad) ([]float64, error) {
	stores := dba.Group_stores_by_day(db)
	var fees []float64

	for _, load := range loads {
		var fee float64 = 0

		for _, store := range stores {
			// 出库日期与入库日期的差，非正即错
			time_interval := utils.Sub_day(load.Load_date,
				store.Store_date)

			// 若库存不充足，才会遍历到超过当前出库日期的库存
			if time_interval <= 0 {
				return nil, errors.New(utils.China_date(store.Store_date))
			} else {
				// 按更小的量出货
				fee += utils.Fee(time_interval,
					math.Min(store.Store_ton, load.Load_ton))

				if load.Loads >= store.Stocks {
					// 消耗完了一日的库存，则移除库存信息
					load.Load_ton -= store.Store_ton
					load.Loads -= store.Stocks

					stores = stores[1:]

					// 恰好完全出货，中止内循环
					if load.Loads == 0 {
						break
					}
				} else if load.Loads < store.Stocks {
					// 库存量超过出货量，则库存量减去出货量，并中止
					stores[0].Stocks -= load.Loads
					stores[0].Store_ton -= load.Load_ton

					break
				}
			}
		}

		fees = append(fees, fee)
	}

	return fees, nil
}
