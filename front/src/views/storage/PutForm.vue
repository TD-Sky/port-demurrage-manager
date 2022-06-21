<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { PutStorage } from '@/models/storage';
import type { FormInstance, FormRules } from 'element-plus';
import { isnt_future } from '@/utils/index';

// 共享状态
const props = defineProps(["opening", "buffer"]);
const { opening, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "put_then_refresh"]);

const label_width = '70px';
const rule_form = reactive(<FormInstance>{});
const rules = reactive<FormRules>({
    store_date: [
        {
            type: "date",
            required: true,
            message: "选择未来日期",
            trigger: "change",
        }
    ],

    license_plate_number: [
        {
            type: "string",
            required: true,
            message: "车牌号",
            trigger: "change",
        }
    ],

    stocks: [
        {
            type: "number",
            required: true,
            message: "输入正整数",
            trigger: "change",
        }
    ],
})

// 只能子组件执行，父组件收不到参数。
// 传入的是响应类型内部值
async function submit_form(form_elt: FormInstance | undefined, buf: PutStorage) {
    if (!form_elt) return
    await form_elt.validate((valid, _) => {
        if (valid) {
            emits("close_form", 'put')
            emits("put_then_refresh", buf)
        }
    })
}
</script>

<template>

    <el-dialog v-model="opening['put']" title="入库单" @close="emits('close_form', 'put')">
        <el-form :rules="rules" ref="rule_form" :model="buffer">

            <el-form-item prop="store_date" label="日期" :label-width="label_width">
                <el-date-picker type="date" v-model="buffer.store_date" :disabled-date="isnt_future" />
            </el-form-item>

            <el-form-item prop="license_plate_number" label="车牌号" :label-width="label_width">
                <el-input v-model="buffer.license_plate_number" maxlength="7" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="stocks" label="件数" :label-width="label_width">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.stocks" />
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="submit_form(rule_form, buffer)">
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
