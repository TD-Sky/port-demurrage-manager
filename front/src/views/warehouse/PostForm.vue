<script setup lang="ts">
import { toRefs } from 'vue';
import { Warehouse } from '@/models/warehouse';

const FIRST_COL = "50px";

// 共享状态
const props = defineProps(["show_post", "buffer"]);
const { show_post, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const submit = (buf: Warehouse) => {
    emits("close_form")
    emits("post_then_refresh", buf)
}
</script>

<template>
    <el-dialog v-model="show_post" width="30%" title="新增场地" @close="emits('close_form')">
        <el-form :model="buffer">

            <el-form-item prop="house_name" label="名字" :label-width="FIRST_COL">
                <el-input v-model="buffer.house_name" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="address" label="地址" :label-width="FIRST_COL">
                <el-input v-model="buffer.address" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="area" label="面积" :label-width="FIRST_COL">
                <el-input-number :min="0" :controls="false" v-model.number="buffer.area" />
                <div style="margin-left: 20px;">
                    m<sup>2</sup>
                </div>
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="submit(buffer)">
                    确认
                </el-button>
            </span>
        </template>
    </el-dialog>

</template>

<style scoped>
.dialog-footer {
    margin-right: 10px;
}
</style>
