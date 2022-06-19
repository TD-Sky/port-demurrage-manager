<script setup lang="ts">
import { Warehouse } from '@/models/warehouse';
import { get_warehouses, post_warehouse, put_warehouse, delete_warehouse } from './request';
import { ref, reactive } from "vue";
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';
import DelDialog from './DelDialog.vue';

const warehouses = ref<Warehouse[]>([]);                // 表格数据
const buffer = reactive(<Warehouse>{});                 // 表单填写缓冲区
const remove_name = ref<string>();                      // 删除条目的ID

const opening = reactive<{ [key: string]: boolean }>({
    "post": false,
    "put": false,
    "delete": false,
});                                                       // 会话开关

const build_table = () => {
    get_warehouses().then((resp) => {
        if (resp.data.warehouses !== null) {
            warehouses.value = resp.data.warehouses;
        } else {
            warehouses.value = [];
        }
    });
}

const modify_form = (warehouse: Warehouse) => {
    buffer.house_name = warehouse.house_name;
    buffer.address = warehouse.address;
    buffer.area = warehouse.area;

    opening["put"] = true;
}

const new_form = () => {
    opening["post"] = true;
}

const remove_item = (house_name: string) => {
    remove_name.value = house_name;
    opening["delete"] = true;
}

const close_form = (kind: string) => {
    opening[kind] = false
}

const put_then_refresh = (data: Warehouse) => {
    put_warehouse(data).then((_) => {
        build_table();
    })
}

const post_then_refresh = (data: Warehouse) => {
    post_warehouse(data).then((_) => {
        build_table();
    })
}

const delete_then_refresh = (house_name: string) => {
    delete_warehouse(house_name).then((_) => {
        opening["delete"] = false;
        build_table();
    })
}

build_table();
</script>

<template>
    <Component>
        <div class="app-container">

            <div class="warehouse-table">
                <el-table :data="warehouses" border :fit="false" max-height="1000">
                    <el-table-column header-align="center" prop="house_name" label="名字" width="120" />
                    <el-table-column header-align="center" prop="address" label="地址" width="120" />
                    <el-table-column header-align="center" prop="area" label="面积" width="120" />
                    <el-table-column align="center" label="操作" width="180">
                        <template #default="scope">
                            <el-button type="primary" @click="modify_form(scope.row)" :icon="Edit" />
                            <el-button type="danger" @click="remove_item(scope.row.house_name)" :icon="Delete" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div class="plus-button" @click="new_form">
                <el-icon :size="24">
                    <Plus />
                </el-icon>
            </div>

            <PutForm :opening="opening" :buffer="buffer" @close_form="close_form"
                @put_then_refresh="put_then_refresh" />

            <PostForm :opening="opening" @close_form="close_form" @post_then_refresh="post_then_refresh" />

            <DelDialog :opening="opening" :remove_name="remove_name" @close_form="close_form"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
</template>

<style lang="scss" scoped>
.warehouse-table {
    width: 100%;
}
</style>
