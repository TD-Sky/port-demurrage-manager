<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { PostLoad } from '@/models/index';
import { isnt_future } from '@/utils/index';
import type { FormInstance, FormRules } from 'element-plus';

// 共享状态
const props = defineProps(["opening", "buffer"]);
const { opening, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const label_width = '70px';
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

    load_ton: [
        {
            type: "float",
            required: true,
            message: "输入正实数，小数点后有两位",
            trigger: "change",
        }
    ],

})

// 只能子组件执行，父组件收不到参数。
// 传入的是响应类型内部值
async function submit_form(form_elt: FormInstance | undefined, buf: PostLoad) {
    if (!form_elt) return
    await form_elt.validate((valid, _) => {
        if (valid) {
            emits("close_form", "post")
            emits("post_then_refresh", buf)
        }
    })
}
</script>

<template>
    <el-dialog v-model="opening['post']" title="出库单" @close="emits('close_form', 'post')">
        <el-form :rules="rules" ref="rule_form" :model="buffer">

            <el-form-item prop="load_date" label="日期" :label-width="label_width">
                <el-date-picker type="date" v-model="buffer.load_date" :disabled-date="isnt_future" />
            </el-form-item>

            <el-form-item prop="loads" label="件数" :label-width="label_width">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.loads" />
            </el-form-item>

            <el-form-item prop="load_ton" label="吨数" :label-width="label_width">
                <el-input-number :min="1.00" :controls="false" :precision="2" v-model="buffer.load_ton" />
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="submit_form(rule_form, buffer)">提交</el-button>
            </span>
        </template>
    </el-dialog>

</template>

<style scoped>
.dialog-footer {
    margin-right: 10px;
}
</style>
