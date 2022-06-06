<script setup lang="ts">
import { GetStorage, PutStorage, PostStorage } from '@/models/index';
import { Edit, Plus } from '@element-plus/icons-vue';
import { request } from '@/utils/service';
import { onBeforeMount, ref, reactive } from "vue";
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';

const storages = ref<GetStorage[]>([])                    // 库存信息数组
const buffer = reactive(<PostStorage | PutStorage>{});    // 表单填写缓冲区
const opening = reactive<{ [key: string]: boolean }>({
    "post": false,
    "put": false,
});

const get_stores = () => {
    request({
        url: "/store",
        method: "get",
    }).then((resp) => {
        storages.value = resp.data.stores
    })
}

function post_store(data: PostStorage) {
    return request({
        url: "/store",
        method: "post",
        data,
    })
}

function put_store(data: PutStorage) {
    return request({
        url: "/store",
        method: "put",
        data,
    })
}

const modify_form = (storage: GetStorage) => {
    Object.assign(buffer, storage);
    opening["put"] = true;
}

const new_form = () => {
    Object.assign(buffer, {
        store_date: "",
        license_plate_number: "",
        stocks: null,
        store_ton: null,
    });
    opening["post"] = true;
}

const close_form = (kind: string) => {
    opening[kind] = false
}

function put_then_refresh(data: PutStorage) {
    put_store(data).then((_) => {
        get_stores()
    })
}

function post_then_refresh(data: PostStorage) {
    post_store(data).then((_) => {
        get_stores()
    })
}

onBeforeMount(get_stores)
</script>

<template>
    <div class="app-container">
        <el-row :gutter="50">

            <el-col :span="50">
                <el-card class="box-card" shadow="hover">
                    <template #header>
                        <div class="card-header">
                            <span style="font-size:20px">预约入库</span>
                            <el-button size="large" :icon="Plus" @click="new_form()" />
                        </div>
                    </template>

                    <ul>
                        <li class="text item"> 入库日期 </li>
                        <li class="text item"> 车牌号 </li>
                        <li class="text item"> 件数 </li>
                        <li class="text item"> 吨数 </li>
                    </ul>

                </el-card>
            </el-col>

            <el-col :span="50" v-for="storage in storages">

                <el-card class="box-card" shadow="never">
                    <template #header>
                        <div class="card-header">
                            <span>ID：{{ storage.id }}</span>
                            <el-button size="large" @click="modify_form(storage)" :icon="Edit" />
                        </div>
                    </template>

                    <el-descriptions :column="1" border>
                        <el-descriptions-item label="入库日期">
                            {{ storage.store_date }}
                        </el-descriptions-item>

                        <el-descriptions-item label="车牌号">
                            {{ storage.license_plate_number }}
                        </el-descriptions-item>

                        <el-descriptions-item label="件数">
                            {{ storage.stocks }}
                        </el-descriptions-item>

                        <el-descriptions-item label="吨数">
                            {{ storage.store_ton }}
                        </el-descriptions-item>

                        <el-descriptions-item label="仓储天数">
                            {{ storage.duration }}
                        </el-descriptions-item>

                        <el-descriptions-item label="堆存费">
                            {{ storage.fee }}
                        </el-descriptions-item>
                    </el-descriptions>

                </el-card>

                <div class="empty"></div>
            </el-col>
        </el-row>

        <PutForm :opening="opening" :buffer="buffer" @close_form="close_form" @put_then_refresh="put_then_refresh" />

        <PostForm :opening="opening" :buffer="buffer" @close_form="close_form" @post_then_refresh="post_then_refresh" />

    </div>
</template>

<style>
.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.text {
    font-size: 18px;
}

.item {
    margin-bottom: 30px;
}

.box-card {
    width: 400px;
}

.empty {
    height: 25px;
}
</style>
