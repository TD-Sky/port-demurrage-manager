<script setup lang="ts">
import { GetStorage, PutStorage, PostStorage } from '@/models/index';
import { get_stores, post_store, put_store, delete_store } from './request';
import { fees_incurred } from '@/utils/index';
import { ref, reactive } from "vue";
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

const build_table = () => {
    get_stores().then((resp) => {
        storages.value = resp.data.stores;
    });
}

const modify_form = (storage: GetStorage) => {
    buffer.id = storage.id;
    buffer.stocks = storage.stocks;
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
        build_table();
    })
}

const post_then_refresh = (data: PostStorage) => {
    post_store(data).then((_) => {
        build_table();
    })
}

const delete_then_refresh = (id: number) => {
    delete_store(id).then((_) => {
        opening["delete"] = false
        build_table();
    })
}

build_table();
</script>

<template>
    <Component>
        <div class="app-container">

            <div class="stores-table">
                <el-table :data="storages" border :fit="false" max-height="1000">
                    <el-table-column header-align="center" prop="store_date" label="日期" width="120" />
                    <el-table-column header-align="center" prop="license_plate_number" label="车牌号" width="120" />
                    <el-table-column header-align="center" prop="stocks" label="件数" width="120" />
                    <el-table-column header-align="center" prop="store_ton" label="吨数" width="120" />
                    <el-table-column header-align="center" prop="duration" label="仓储天数" width="120" />
                    <el-table-column align="center" label="操作" width="180">
                        <template #default="scope">
                            <el-button type="primary" @click="modify_form(scope.row)" :icon="Edit"
                                :disable="fees_incurred(scope.duration)" />
                            <el-button type="danger" @click="remove_item(scope.row.id)" :icon="Delete"
                                :disable="fees_incurred(scope.duration)" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div class="handle-button" @click="new_form">
                <el-icon :size="24">
                    <Plus />
                </el-icon>
            </div>

            <PutForm :opening="opening" :buffer="buffer" @close_form="close_form"
                @put_then_refresh="put_then_refresh" />

            <PostForm :opening="opening" @close_form="close_form" @post_then_refresh="post_then_refresh" />

            <DelDialog :opening="opening" :remove_id="remove_id" @close_form="close_form"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
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
    top: 300px;
}

.stores-table {
    width: 100%;
}
</style>
