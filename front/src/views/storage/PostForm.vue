<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { PostStorage } from '@/models/storage';
import { isnt_future } from '@/utils/index';

const FIRST_COL = "70px";

// 共享状态
const props = defineProps(["show_post"]);
const { show_post } = toRefs(props);

const buffer = reactive(<PostStorage>{});

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const submit = (buf: PostStorage) => {
    emits("close_form")
    emits("post_then_refresh", buf)
}
</script>

<template>
    <el-dialog v-model="show_post" title="入库单" @close="emits('close_form')">
        <el-form :model="buffer">

            <el-form-item prop="store_date" label="日期" :label-width="FIRST_COL">
                <el-date-picker type="date" v-model="buffer.store_date" :disabled-date="isnt_future" />
            </el-form-item>

            <el-form-item prop="license_plate_number" label="车牌号" :label-width="FIRST_COL">
                <el-input v-model="buffer.license_plate_number" maxlength="7" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="stocks" label="件数" :label-width="FIRST_COL">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.stocks" />
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
