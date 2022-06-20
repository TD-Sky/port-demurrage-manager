<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { PutLoad } from '@/models/load';
import { isnt_future } from '@/utils/index';
import type { FormInstance, FormRules } from 'element-plus';

const FIRST_COL = "50px";

// 共享状态
const props = defineProps(["opening", "buffer"]);
const { opening, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "put_then_refresh"]);

const rule_form = reactive(<FormInstance>{});
const rules = reactive<FormRules>({
    load_date: [
        {
            type: "date",
            required: true,
            message: "选择未来日期",
            trigger: "change",
        }
    ],

    loads: [
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
async function submit_form(form_elt: FormInstance | undefined, buf: PutLoad) {
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
    <el-dialog v-model="opening['put']" title="出库单" @close="emits('close_form', 'put')">
        <el-form :rules="rules" ref="rule_form" :model="buffer">

            <el-form-item prop="load_date" label="日期" :label-width="FIRST_COL">
                <el-date-picker type="date" v-model="buffer.load_date" :disabled-date="isnt_future" />
            </el-form-item>

            <el-form-item prop="loads" label="件数" :label-width="FIRST_COL">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.loads" />
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
