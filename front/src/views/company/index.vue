<script setup lang="ts">
import { FreightForwarder } from '@/models/freight_forwarder';
import { get_dataset, post_data, put_data, delete_data } from '@/api/crud';
import { ref, reactive } from "vue";
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import PutForm from './PutForm.vue';
import PostForm from './PostForm.vue';
import DelDialog from './DelDialog.vue';

const companies = ref<FreightForwarder[]>([]);          // 表格数据
const buffer = reactive(<FreightForwarder>{});          // 表单填写缓冲区
const remove_code = ref<string>();                      // 删除条目的ID

// 显示对话框的函数
const show_post = ref(false);
const show_put = ref(false);
const show_delete = ref(false);

const build_table = () => {
    get_dataset("/freight_forwarder").then((resp) => {
        if (resp.data.companies !== null) {
            companies.value = resp.data.companies;
        } else {
            companies.value = [];
        }
    });
}

const modify_form = (company: FreightForwarder) => {
    buffer.code = company.code;
    buffer.company_name = company.company_name;
    buffer.telephone_number = company.telephone_number;

    show_put.value = true;
}

const remove_item = (house_name: string) => {
    remove_code.value = house_name;

    show_delete.value = true;
}

const put_then_refresh = (data: FreightForwarder) => {
    put_data("/freight_forwarder", data).then((_) => {
        build_table();
    })
}

const post_then_refresh = (data: FreightForwarder) => {
    post_data("/freight_forwarder", data).then((_) => {
        build_table();
    })
}

const delete_then_refresh = (code: string) => {
    delete_data("/freight_forwarder", code).then((_) => {
        show_delete.value = false;
        build_table();
    })
}

build_table();
</script>

<template>
    <Component>
        <div class="app-container">

            <div class="company-table">
                <el-table :data="companies" border :fit="false" max-height="1000">
                    <el-table-column header-align="center" prop="code" label="识别码" width="200" />
                    <el-table-column header-align="center" prop="company_name" label="名字" width="300" />
                    <el-table-column header-align="center" prop="telephone_number" label="联系方式" width="150" />
                    <el-table-column align="center" label="操作" width="180">
                        <template #default="scope">
                            <el-button type="primary" @click="modify_form(scope.row)" :icon="Edit" />
                            <el-button type="danger" @click="remove_item(scope.row.code)" :icon="Delete" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div class="plus-button" @click="show_post = true">
                <el-icon :size="24">
                    <Plus />
                </el-icon>
            </div>

            <PostForm :show_post="show_post" @close_form="show_post = false" @post_then_refresh="post_then_refresh" />

            <PutForm :show_put="show_put" :buffer="buffer" @close_form="show_put = false"
                @put_then_refresh="put_then_refresh" />

            <DelDialog :show_delete="show_delete" :remove_code="remove_code" @close_form="show_delete = false"
                @delete_then_refresh="delete_then_refresh" />
        </div>
    </Component>
</template>

<style lang="scss" scoped>
.company-table {
    width: 100%;
}
</style>
