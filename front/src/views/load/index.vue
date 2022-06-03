<script setup lang="ts">
import { reactive, onBeforeMount, ref } from 'vue';
import { request } from '@/utils/service';
import { Edit } from '@element-plus/icons-vue'
import { Load } from '@/models/index';
import LoadForm from './form.vue';

const loads = ref<Load[]>([]);
const display_form = ref(false);
const buffer = reactive(<Load>{});

const edit_load = (row: Load) => {
    display_form.value = true;
    Object.assign(buffer, row)
}

onBeforeMount(() => {
    request({
        url: "/load",
        method: "get",
    }).then((resp) => {
        loads.value = resp.data.loads
    })
})
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
                    <el-button size="small" @click="edit_load(scope.row)" :icon="Edit" />
                </template>
            </el-table-column>

        </el-table>

        <LoadForm :display_form="display_form" :buffer="buffer" />
    </component>

</template>

<style scoped>
.el-button--text {
    margin-right: 15px;
}

.el-select {
    width: 300px;
}

.el-input {
    width: 300px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
