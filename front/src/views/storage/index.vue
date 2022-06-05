<script setup lang="ts">
import { GetStorage, PutStorage } from '@/models/index';
import { Edit } from '@element-plus/icons-vue';
import { request } from '@/utils/service';
import { onBeforeMount, ref, reactive } from "vue";
import StoreForm from './StoreForm.vue';

const storages = ref<GetStorage[]>([])         // 库存信息数组
const buffer = reactive(<PutStorage>{});    // 表单填写缓冲区
const display_form = ref(false);            // 是否打开表单

const get_stores = () => {
    request({
        url: "/store",
        method: "get",
    }).then((resp) => {
        storages.value = resp.data.stores
    })
}

function put_store(data: PutStorage) {
    return request({
        url: "/store",
        method: "put",
        data,
    })
}

const open_form = (storage: GetStorage) => {
    Object.assign(buffer, storage);
    display_form.value = true
}

const close_form = () => {
    display_form.value = false
}

function upload_and_refresh(data: PutStorage) {
    put_store(data).then((_) => {
        get_stores()
    })
}

onBeforeMount(get_stores)
</script>

<template>
    <div class="app-container">
        <el-row :gutter="50">
            <el-col :span="50" v-for="storage in storages">

                <el-card class="box-card">
                    <template #header>
                        <div class="card-header">
                            <span>ID：{{ storage.id }}</span>
                            <el-button size="small" @click="open_form(storage)" :icon="Edit" />
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

        <StoreForm :display_form="display_form" :buffer="buffer" @close_form="close_form"
            @upload_and_refresh="upload_and_refresh" />
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
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    width: 400px;
}

.empty {
    height: 25px;
}
</style>
