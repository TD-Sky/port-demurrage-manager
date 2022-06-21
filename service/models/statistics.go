package models

import (
	"service/utils"
	"sort"
)

type StatLoad struct {
	Company_code string
	Load_date    string
	Fee          float64
}

// 坐标面
type Plane struct {
	Load_date []string  `json:"load_date"`
	Fee       []float64 `json:"fee"`
}

type StatLoadMap = map[string]Plane

func to_StatLoadVec(loads []GetLoad) []StatLoad {
	var vec []StatLoad

	for _, load := range loads {
		vec = append(vec, StatLoad{
			Company_code: load.Company_code,
			Load_date:    utils.China_date(load.Load_date),
			Fee:          load.Fee,
		})
	}

	sort.SliceStable(vec, func(i, j int) bool {
		if vec[i].Company_code < vec[j].Company_code {
			return true
		} else if vec[i].Load_date < vec[j].Load_date {
			return true
		} else {
			return false
		}
	})

	return vec
}

func To_StatLoadMap(loads []GetLoad) StatLoadMap {
	vec := to_StatLoadVec(loads)
	load_map := make(StatLoadMap)
	length := len(vec)
	var fee float64

	for i, load := range vec {
		// 保证公司必须有数组
		plane, contains_code := load_map[load.Company_code]

		// 没有此公司，则新建点数组
		if !contains_code {
			load_map[load.Company_code] = Plane{
				Load_date: []string{},
				Fee:       []float64{},
			}
		}

		if i+1 == length ||
			vec[i+1].Company_code != load.Company_code ||
			vec[i+1].Load_date != load.Load_date {
			/*
			 * 结算时机
			 * - 最后一轮循环
			 * - 公司即将变化
			 * - 日期即将变化
			 */

			load_map[load.Company_code] = Plane{
				Load_date: append(plane.Load_date, load.Load_date),
				Fee:       append(plane.Fee, fee+load.Fee),
			}

			// 清空费用状态
			fee = 0
		} else if vec[i+1].Load_date == load.Load_date {
			// 公司、日期均为变化，累加堆存费
			fee += load.Fee
		}
	}

	return load_map
}
