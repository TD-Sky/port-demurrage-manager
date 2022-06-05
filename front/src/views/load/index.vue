<script setup lang="ts">
import { reactive, onBeforeMount, ref } from 'vue';
import { request } from '@/utils/service';
import { Edit } from '@element-plus/icons-vue';
import { GetLoad, PutLoad } from '@/models/index';
import LoadForm from './LoadForm.vue';

const loads = ref<GetLoad[]>([]);       // 出货单
const buffer = reactive(<PutLoad>{});   // 表单填写缓冲区
const display_form = ref(false);        // 是否打开表单

// 使用闭包是为了配合 onBeforeMount
const get_loads = () => {
    request({
        url: "/load",
        method: "get",
    }).then((resp) => {
        loads.value = resp.data.loads
    })
}

function put_load(data: PutLoad) {
    return request({
        url: "/load",
        method: "put",
        data,
    })
}

// 必须用闭包，
// 模板部分无法访问完整的响应类型，不能传响应类型值
const open_form = (row: GetLoad) => {
    Object.assign(buffer, row);
    display_form.value = true
}

// 子组件修改父组件状态不会生效！！！
// 必须靠 emits 触发此函数
const close_form = () => {
    display_form.value = false
}

function upload_and_refresh(data: PutLoad) {
    put_load(data).then((_) => {
        get_loads()
    })
}

// 进入网页，先读一次数据库
onBeforeMount(get_loads)
</script>

<template>
    <component>
        <el-table :data="loads" style="width: 100%" border>
            <el-table-column prop="load_date" label="日期" width="120" />
            <el-table-column prop="loads" label="件数" width="120" />
            <el-table-column prop="load_ton" label="吨数" width="120" />
            <el-table-column prop="business_number" label="业务号" width="150" />
            <el-table-column prop="lading_bill_number" label="提单号" width="180" />
            <el-table-column prop="order_number" label="订单号" width="150" />
            <el-table-column label="编辑">
                <template #default="scope">
                    <el-button size="small" @click="open_form(scope.row)" :icon="Edit" />
                    <LoadForm :display_form="display_form" :buffer="buffer" @close_form="close_form"
                        @upload_and_refresh="upload_and_refresh" />
                </template>
            </el-table-column>
        </el-table>
    </component>
</template>
