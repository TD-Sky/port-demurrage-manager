<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { FreightForwarder } from '@/models/freight_forwarder';
import type { FormInstance, FormRules } from 'element-plus';

// 共享状态
const props = defineProps(["opening", "buffer"]);
const { opening, buffer } = toRefs(props);

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const label_width = '80px';
const rule_form = reactive(<FormInstance>{});
const rules = reactive<FormRules>({
    code: [
        {
            type: "string",
            required: true,
            message: "识别码",
            trigger: "change",
        }
    ],

    company_name: [
        {
            type: "string",
            required: true,
            message: "公司名",
            trigger: "change",
        }
    ],

    telephone_number: [
        {
            type: "string",
            required: true,
            message: "联系方式",
            trigger: "change",
        }
    ],
})

// 只能子组件执行，父组件收不到参数。
// 传入的是响应类型内部值
async function submit_form(form_elt: FormInstance | undefined, buf: FreightForwarder) {
    if (!form_elt) return
    await form_elt.validate((valid, _) => {
        if (valid) {
            emits("close_form", 'post');
            emits("post_then_refresh", buf);
        }
    })
}
</script>

<template>
    <el-dialog v-model="opening['post']" width="30%" title="加入公司" @close="emits('close_form', 'post')">
        <el-form :rules="rules" ref="rule_form" :model="buffer">

            <el-form-item prop="code" label="识别码" :label-width="label_width">
                <el-input v-model="buffer.code" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="company_name" label="名字" :label-width="label_width">
                <el-input v-model="buffer.company_name" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="telephone_number" label="联系方式" :label-width="label_width">
                <el-input v-model="buffer.telephone_number" maxlength="11" style="width: 200px" />
            </el-form-item>

        </el-form>

        <template #footer>
            <span class="dialog-footer">
                <el-button type="primary" @click="submit_form(rule_form, buffer)">确认</el-button>
            </span>
        </template>
    </el-dialog>

</template>

<style scoped>
.dialog-footer {
    margin-right: 10px;
}
</style>
