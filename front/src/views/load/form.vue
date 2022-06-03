<script setup lang="ts">
import { defineProps, toRefs } from 'vue';

const props = defineProps(["display_form", "buffer"]);
const { display_form, buffer } = toRefs(props);
const formLabelWidth = '140px'

</script>

<template>
    <el-dialog v-model="display_form" title="Shipping address">
        <el-form :model="buffer">

            <el-form-item label="日期" :label-width="formLabelWidth">
                <el-input v-model="buffer.load_date" />
            </el-form-item>

            <el-form-item label="件数" :label-width="formLabelWidth">
                <el-input v-model="buffer.loads" />
            </el-form-item>

            <el-form-item label="吨数" :label-width="formLabelWidth">
                <el-input v-model="buffer.load_ton" />
            </el-form-item>

            <el-form-item label="业务号" :label-width="formLabelWidth">
                <el-input v-model="buffer.business_number" />
            </el-form-item>

            <el-form-item label="提单号" :label-width="formLabelWidth">
                <el-input v-model="buffer.lading_bill_number" />
            </el-form-item>

            <el-form-item label="订单号" :label-width="formLabelWidth">
                <el-input v-model="buffer.order_number" />
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button @click="display_form = false">Cancel</el-button>
                <el-button type="primary" @click="display_form = false">Confirm</el-button>
            </span>
        </template>
    </el-dialog>

</template>
