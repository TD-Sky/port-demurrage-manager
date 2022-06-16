<script lang="ts" setup>
import * as echarts from 'echarts';
import { GetLoad } from '@/models/index';
import { get_loads } from '@/views/load/request';
import { ref, reactive } from "vue";

const chart_dom = ref();
const chart_style = reactive({
    width: "100%",
    height: "550px",
});

const init_chart = (dims: number[], data_src: any[]) => {
    const bar_chart = echarts.init(chart_dom.value);
    const series: echarts.SeriesOption[] = [];

    for (let i = 0; i < dims.length - 1; i++) {
        series.push({
            type: "bar",
        });
    }

    const option: echarts.EChartsOption = {
        // 类型不对报错，不用管
        dataset: {
            dimensions: dims,
            source: data_src,
        },
        xAxis: {
            type: 'category',
        },
        yAxis: {},
        series: series,
    };

    bar_chart.setOption(option);
}

get_loads().then((resp) => {
    const loads: GetLoad[] = resp.data.loads;
    const chart_map: Map<string, number[]> = new Map();

    for (const load of loads) {
        const fees = chart_map.get(load.load_date);

        if (fees === undefined) {
            chart_map.set(load.load_date, [load.fee]);
        } else {
            fees.push(load.fee);
        }
    }

    const dims: number[] = [];

    const max_len = Array.from(chart_map.values())
        .map(fees => fees.length)
        .reduce((prev, current) => {
            return current >= prev ? current : prev;
        });

    for (let i = 0; i <= max_len; i++) {
        dims.push(i);
    }

    const data_src: any[] = [];
    chart_map.forEach((fees, date) => {
        data_src.push([date, fees].flat());
    });

    init_chart(dims, data_src);
});
</script>

<template>
    <Component>
        <div class="app-container">
            <!-- 不能用class，只能用style -->
            <div ref="chart_dom" :style="chart_style">
            </div>
        </div>
    </Component>
</template>
