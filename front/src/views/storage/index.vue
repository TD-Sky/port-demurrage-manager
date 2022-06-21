<script setup lang="ts">
import { GetStorage, PutStorage, PostStorage } from '@/models/storage';
import { get_dataset, post_data, put_data, delete_data } from '@/api/crud';
import { fees_incurred } from '@/utils/index';
import { ref, reactive } from "vue";
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';
import DelDialog from './DelDialog.vue';

const storages = ref<GetStorage[]>([])                    // 库存信息数组
const buffer = reactive(<PutStorage | PostStorage>{});                  // 表单填写缓冲区
const remove_id = ref<number>();                          // 删除条目的ID

// 显示对话框的函数
const show_post = ref(false);
const show_put = ref(false);
const show_delete = ref(false);

const build_table = () => {
    get_dataset("/store").then((resp) => {
        if (resp.data.storages !== null) {
            storages.value = resp.data.stores;
        } else {
            storages.value = [];
        }
    });
}

const new_form = () => {
    // @ts-ignore
    buffer.stocks = null;
    buffer.license_plate_number = "";
    // @ts-ignore
    buffer.store_date = null;

    show_post.value = true;
}

const modify_form = (storage: GetStorage) => {
    (buffer as PutStorage).id = storage.id;
    buffer.stocks = storage.stocks;
    buffer.license_plate_number = storage.license_plate_number;
    buffer.store_date = new Date(storage.store_date)

    show_put.value = true;
}

const remove_item = (id: number) => {
    remove_id.value = id

    show_delete.value = true;
}

const put_then_refresh = (data: PutStorage) => {
    put_data("/store", data).then((_) => {
        build_table();
    })
}

const post_then_refresh = (data: PostStorage) => {
    post_data("/store", data).then((_) => {
        build_table();
    })
}

const delete_then_refresh = (id: number) => {
    delete_data("store", id).then((_) => {
        show_delete.value = false;
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

            <div class="plus-button" @click="new_form">
                <el-icon :size="24">
                    <Plus />
                </el-icon>
            </div>

            <PostForm :show_post="show_post" :buffer="buffer" @close_form="show_post = false"
                @post_then_refresh="post_then_refresh" />

            <PutForm :show_put="show_put" :buffer="buffer" @close_form="show_put = false"
                @put_then_refresh="put_then_refresh" />

            <DelDialog :show_delete="show_delete" :remove_id="remove_id" @close_form="show_delete = false"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
</template>

<style lang="scss" scoped>
.stores-table {
    width: 100%;
}
</style>
