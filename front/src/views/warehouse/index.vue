<script setup lang="ts">
import { Warehouse } from '@/models/warehouse';
import { get_dataset, post_data, put_data, delete_data } from '@/api/crud';
import { ref, reactive } from "vue";
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';
import DelDialog from './DelDialog.vue';

const warehouses = ref<Warehouse[]>([]);                // 表格数据
const buffer = reactive(<Warehouse>{});                 // 表单填写缓冲区
const remove_name = ref<string>();                      // 删除条目的ID

// 显示对话框的函数
const show_post = ref(false);
const show_put = ref(false);
const show_delete = ref(false);

const build_table = () => {
    get_dataset("/warehouse").then((resp) => {
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

    show_put.value = true;
}

const remove_item = (house_name: string) => {
    remove_name.value = house_name;

    show_delete.value = true;
}

const put_then_refresh = (data: Warehouse) => {
    put_data("/warehouse", data).then((_) => {
        build_table();
    })
}

const post_then_refresh = (data: Warehouse) => {
    post_data("/warehouse", data).then((_) => {
        build_table();
    })
}

const delete_then_refresh = (house_name: string) => {
    delete_data("/warehouse", house_name).then((_) => {
        show_delete.value = false;
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

            <!-- 侧栏添加条目按钮 -->
            <div class="plus-button" @click="show_post = true">
                <el-icon :size="24">
                    <Plus />
                </el-icon>
            </div>

            <PostForm :show_post="show_post" @close_form="show_post = false" @post_then_refresh="post_then_refresh" />

            <PutForm :show_put="show_put" :buffer="buffer" @close_form="show_put = true"
                @put_then_refresh="put_then_refresh" />

            <DelDialog :show_delete="show_delete" :remove_name="remove_name" @close_form="show_delete = false"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
</template>

<style lang="scss" scoped>
.warehouse-table {
    width: 100%;
}
</style>
