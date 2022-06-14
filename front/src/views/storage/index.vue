<script setup lang="ts">
import { GetStorage, PutStorage, PostStorage } from '@/models/index';
import { get_stores, post_store, put_store, delete_store } from './request';
import { fees_incurred } from '@/utils/index';
import { onBeforeMount, ref, reactive } from "vue";
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';
import DelDialog from './DelDialog.vue';

const storages = ref<GetStorage[]>([])                    // 库存信息数组
const buffer = reactive(<PutStorage>{});                  // 表单填写缓冲区
const opening = reactive<{ [key: string]: boolean }>({
    "post": false,
    "put": false,
    "delete": false,
});                                                       // 会话开关
const remove_id = ref<number>();                          // 删除条目的ID

const modify_form = (storage: GetStorage) => {
    buffer.id = storage.id;
    buffer.stocks = storage.stocks;
    buffer.store_ton = storage.store_ton;
    buffer.license_plate_number = storage.license_plate_number;
    buffer.store_date = new Date(storage.store_date)

    opening["put"] = true;
}

const new_form = () => {
    opening["post"] = true;
}

const remove_item = (id: number) => {
    remove_id.value = id
    opening["delete"] = true
}

const close_form = (kind: string) => {
    opening[kind] = false
}

const put_then_refresh = (data: PutStorage) => {
    put_store(data).then((_) => {
        get_stores(storages)
    })
}

const post_then_refresh = (data: PostStorage) => {
    post_store(data).then((_) => {
        get_stores(storages)
    })
}

const delete_then_refresh = (id: number) => {
    delete_store(id).then((_) => {
        opening["delete"] = false
        get_stores(storages)
    })
}

onBeforeMount(() => {
    get_stores(storages)
})
</script>

<template>
    <Component>
        <div class="app-container">
            <el-row :gutter="50">

                <el-col :span="50">
                    <el-card class="box-card" shadow="hover">
                        <template #header>
                            <div class="card-header">
                                <span style="font-size:20px">预约入库</span>
                                <el-button size="large" :icon="Plus" @click="new_form" />
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
                                <span>
                                    ID：{{ storage.id }}
                                </span>
                                <span>
                                    <el-button type="primary" @click="modify_form(storage)" :icon="Edit"
                                        :disabled="fees_incurred(storage.duration)" />
                                    <el-button type="danger" @click="remove_item(storage.id)" :icon="Delete"
                                        :disabled="fees_incurred(storage.duration)" />
                                </span>
                            </div>
                        </template>

                        <el-descriptions column="1" border>
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

                        </el-descriptions>

                    </el-card>

                </el-col>
            </el-row>

            <PutForm :opening="opening" :buffer="buffer" @close_form="close_form"
                @put_then_refresh="put_then_refresh" />

            <PostForm :opening="opening" @close_form="close_form" @post_then_refresh="post_then_refresh" />

            <DelDialog :opening="opening" :remove_id="remove_id" @close_form="close_form"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
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
    margin-bottom: 20px;
}

.box-card {
    width: 400px;
    margin-bottom: 25px;
}
</style>
