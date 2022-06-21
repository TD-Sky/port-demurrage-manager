<script setup lang="ts">
import { toRefs } from 'vue';
import { FreightForwarder } from '@/models/freight_forwarder';

const FIRST_COL = "80px";

// 共享状态
const props = defineProps(["show_post", "buffer"]);
const { show_post, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const submit = (buf: FreightForwarder) => {
    emits("close_form")
    emits("post_then_refresh", buf)
}
</script>

<template>
    <el-dialog v-model="show_post" width="30%" title="加入公司" @close="emits('close_form')">
        <el-form :model="buffer">

            <el-form-item prop="code" label="识别码" :label-width="FIRST_COL">
                <el-input v-model="buffer.code" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="company_name" label="名字" :label-width="FIRST_COL">
                <el-input v-model="buffer.company_name" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="telephone_number" label="联系方式" :label-width="FIRST_COL">
                <el-input v-model="buffer.telephone_number" maxlength="11" style="width: 200px" />
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
