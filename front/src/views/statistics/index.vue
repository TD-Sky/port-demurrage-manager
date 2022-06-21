<script lang="ts" setup>
import * as echarts from 'echarts';
import { FreightForwarder } from '@/models/freight_forwarder';
import { get_dataset } from '@/api/crud';
import { ref, reactive, watch } from "vue";
import { StatLoadMap } from '@/models/statistics';

const CHART_STYLE = reactive({
    width: "100%",
    height: "550px",
});

const chart_dom = ref();

const company_code = ref("");
const companies = ref<{
    code: string,
    name: string,
}[]>([]);

const load_map = reactive(<StatLoadMap>{});

const init_chart = (code: string) => {
    const bar_chart = echarts.init(chart_dom.value);

    const option: echarts.EChartsOption = {
        tooltip: {},
        xAxis: {
            type: "category",
            // @ts-ignore
            data: load_map[code].load_date,
        },
        yAxis: {},
        series: [
            {
                type: "bar",
                name: code,
                barMaxWidth: 50,
                // @ts-ignore
                data: load_map[code].fee,
            }
        ]
    };

    bar_chart.setOption(option);
}

get_dataset("/stat_load").then((resp) => {
    if (resp.data.load_map !== null) {
        Object.assign(load_map, resp.data.load_map);

        const cpns = Object.keys(load_map);

        // 获取所有货代公司
        get_dataset("/freight_forwarder").then((resp) => {
            if (resp.data.companies !== null) {
                companies.value = (resp.data.companies as FreightForwarder[])
                    .filter(ff => cpns.includes(ff.code))
                    .map(ff => ({
                        code: ff.code,
                        name: ff.company_name,
                    }));
            }
        });
    }
});

// 由用户选择货代公司
watch(company_code, () => init_chart(company_code.value));
</script>

<template>
    <Component>
        <div class="app-container">

            <el-select v-model="company_code" placeholder="选择货代公司">
                <!-- 渲染label，赋值value -->
                <el-option v-for="cpn in companies" :key="cpn.code" :value="cpn.code" :label="cpn.name" />
            </el-select>

            <!-- 不能用class，只能用style -->
            <div ref="chart_dom" :style="CHART_STYLE"></div>

        </div>
    </Component>
</template>
