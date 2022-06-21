<script setup lang="ts">
import { toRefs } from 'vue';
import { Warehouse } from '@/models/warehouse';

const FIRST_COL = "50px";

// 共享状态
const props = defineProps(["show_put", "buffer"]);
const { show_put, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "put_then_refresh"]);

const submit = (buf: Warehouse) => {
    emits("close_form")
    emits("put_then_refresh", buf)
}
</script>

<template>
    <el-dialog v-model="show_put" width="30%" title="修改场地信息" @close="emits('close_form')">
        <el-form :model="buffer">

            <el-form-item prop="address" label="地址" :label-width="FIRST_COL">
                <el-input v-model="buffer.address" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="area" label="面积" :label-width="FIRST_COL">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.area" />
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
