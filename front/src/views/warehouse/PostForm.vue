<script setup lang="ts">
import { reactive, toRefs } from 'vue';
import { Warehouse } from '@/models/warehouse';
import type { FormInstance, FormRules } from 'element-plus';

// 共享状态
const props = defineProps(["show_post"]);
const { show_post } = toRefs(props);

const buffer = reactive(<Warehouse>{});

// 将触发的父组件函数
const emits = defineEmits(["close_form", "post_then_refresh"]);

const label_width = '50px';
const rule_form = reactive(<FormInstance>{});
const rules = reactive<FormRules>({
    house_name: [
        {
            type: "string",
            required: true,
            message: "唯一名字",
            trigger: "change",
        }
    ],

    address: [
        {
            type: "string",
            required: true,
            message: "地址",
            trigger: "change",
        }
    ],

    area: [
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
async function submit_form(form_elt: FormInstance | undefined, buf: Warehouse) {
    if (!form_elt) return
    await form_elt.validate((valid, _) => {
        if (valid) {
            emits("close_form")
            emits("post_then_refresh", buf)
        }
    })
}
</script>

<template>
    <el-dialog v-model="show_post" width="30%" title="新增场地" @close="emits('close_form')">
        <el-form :rules="rules" ref="rule_form" :model="buffer">

            <el-form-item prop="house_name" label="名字" :label-width="label_width">
                <el-input v-model="buffer.house_name" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="address" label="地址" :label-width="label_width">
                <el-input v-model="buffer.address" style="width: 200px" />
            </el-form-item>

            <el-form-item prop="area" label="面积" :label-width="label_width">
                <el-input-number :min="1" :controls="false" v-model.number="buffer.area" />
                <div style="margin-left: 20px;">
                    m<sup>2</sup>
                </div>
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
