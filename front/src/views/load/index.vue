<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import { request } from '@/utils/service';
import { Edit } from '@element-plus/icons-vue'

type Load = {
    id: number,
    order_number: string,
    load_date: string,
    loads: number,
    load_ton: number,
    business_number: string,
    lading_bill_number: string,
}

const loads = ref<Load[]>([]);

onBeforeMount(() => {
    request({
        url: "/load",
        method: "get",
    }).then((resp) => {
        loads.value = resp.data.loads
    })
})

const handleEdit = (index: number, row: Load) => {
    console.log(index, row)
}
</script>

<template>
    <component>
        <el-table :data="loads" style="width: 100%">

            <el-table-column prop="load_date" label="日期" width="180" />
            <el-table-column prop="loads" label="件数" width="180" />
            <el-table-column prop="load_ton" label="吨数" width="180" />
            <el-table-column prop="business_number" label="业务号" width="180" />
            <el-table-column prop="lading_bill_number" label="提单号" width="180" />
            <el-table-column prop="order_number" label="订单号" width="180" />

            <el-table-column label="编辑">
                <template #default="scope">
                    <el-button size="small" @click="handleEdit(scope.$index, scope.row)" :icon="Edit" />
                </template>
            </el-table-column>

        </el-table>
    </component>
</template>
