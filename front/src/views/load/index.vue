<script setup lang="ts">
import { reactive, onBeforeMount, ref } from 'vue';
import { request } from '@/utils/service';
import { Edit, Plus } from '@element-plus/icons-vue';
import { GetLoad, PutLoad, PostLoad } from '@/models/index';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';

const loads = ref<GetLoad[]>([]);                   // 出货单
const buffer = reactive(<PutLoad | PostLoad>{});    // 表单填写缓冲区
const opening = reactive<{ [key: string]: boolean }>({
    post: false,
    put: false,
});

// 使用闭包是为了配合 onBeforeMount
const get_loads = () => {
    request({
        url: "/load",
        method: "get",
    }).then((resp) => {
        loads.value = resp.data.loads
    })
}

function post_load(data: PostLoad) {
    return request({
        url: "/load",
        method: "post",
        data,
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
const modify_form = (row: GetLoad) => {
    Object.assign(buffer, row);
    opening["put"] = true;
}

const new_form = () => {
    Object.assign(buffer, {
        load_date: "",
        loads: null,
        load_ton: null,
    });
    opening["post"] = true;
}

// 子组件修改父组件状态不会生效！！！
// 必须靠 emits 触发此函数
const close_form = (kind: string) => {
    opening[kind] = false
}

function post_then_refresh(data: PostLoad) {
    post_load(data).then((_) => {
        get_loads()
    })
}

function put_then_refresh(data: PutLoad) {
    put_load(data).then((_) => {
        get_loads()
    })
}

// 侧边添加按钮用的
defineProps({
    buttonTop: {
        type: Number,
        default: 300
    }
});


// 进入网页，先读一次数据库
onBeforeMount(get_loads)
</script>

<template>
    <div class="app-container">
        <el-table :data="loads" style="width:100%" border :fit="false">
            <el-table-column prop="load_date" label="日期" width="120" />
            <el-table-column prop="loads" label="件数" width="120" />
            <el-table-column prop="load_ton" label="吨数" width="120" />
            <el-table-column prop="business_number" label="业务号" width="150" />
            <el-table-column prop="lading_bill_number" label="提单号" width="180" />
            <el-table-column prop="order_number" label="订单号" width="150" />
            <el-table-column label="编辑">
                <template #default="scope">
                    <el-button @click="modify_form(scope.row)" :icon="Edit" />
                </template>
            </el-table-column>
        </el-table>
    </div>

    <div class="handle-button" :style="{ top: buttonTop + 'px' }" @click="new_form">
        <el-icon :size="24">
            <Plus />
        </el-icon>
    </div>

    <PutForm :opening="opening" :buffer="buffer" @close_form="close_form" @put_then_refresh="put_then_refresh" />
    <PostForm :opening="opening" :buffer="buffer" @close_form="close_form" @post_then_refresh="post_then_refresh" />
</template>

<style lang="scss" scoped>
.handle-button {
    width: 48px;
    height: 48px;
    background-color: #5f9ea0;
    position: absolute;
    right: 0px;
    text-align: center;
    font-size: 24px;
    border-radius: 6px 0 0 6px !important;
    z-index: 10;
    cursor: pointer;
    pointer-events: auto;
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
}
</style>
