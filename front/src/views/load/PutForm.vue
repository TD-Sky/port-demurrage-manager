<script setup lang="ts">
import { toRefs } from 'vue';
import { PutLoad } from '@/models/load';
import { isnt_future } from '@/utils/index';

const FIRST_COL = "50px";

// 共享状态
const props = defineProps(["show_put", "buffer"]);
const { show_put, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "put_then_refresh"]);

const submit = (buf: PutLoad) => {
    emits("close_form")
    emits("put_then_refresh", buf)
}
</script>

<template>
    <el-dialog v-model="show_put" title="出库单" @close="emits('close_form')">
        <el-form :model="buffer">

            <el-form-item prop="load_date" label="日期" :label-width="FIRST_COL">
                <el-date-picker type="date" v-model="buffer.load_date" :disabled-date="isnt_future" />
            </el-form-item>

            <el-form-item prop="loads" label="件数" :label-width="FIRST_COL">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.loads" />
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="submit(buffer)">
                    提交
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
