<script lang="ts" setup>
import { reactive, ref } from 'vue';
import { post_loads } from './request';
import type { FormInstance, FormRules } from 'element-plus'
import { Plus, Delete } from '@element-plus/icons-vue';
import { PostLoad } from '@/models/index';
import { isnt_future } from '@/utils/index';
import { useRouter } from 'vue-router';

const router = useRouter();

const buffer = ref<PostLoad[]>([
    <PostLoad>{},
]);

const rule_form = reactive(<FormInstance>{});

const submitForm = (form_elts: FormInstance | undefined) => {

    if (!form_elts) {
        return;
    }

    Object.keys(form_elts).forEach(async (key) => {
        const form_elt = form_elts[key]     // 报错，但不能碰
        try {
            await form_elt.validate()
        } catch (exp) {
            return;
        }
    });

    post_loads(buffer.value).then((_) => {
        router.push("/load");
    });
}

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
            type: "number",
            required: true,
            message: "输入正实数，小数点后有两位",
            trigger: "change",
        }
    ],

})

const remove_load = (elt: PostLoad) => {
    const idx = buffer.value.indexOf(elt)

    if (idx !== -1) {
        buffer.value.splice(idx, 1)
    }
}

const add_load = () => {
    buffer.value.push(<PostLoad>{});
}

</script>

<template>
    <Component>
        <div class="app-container">
            <!-- ref 报错，但不能碰 -->
            <el-form v-for="(elt, idx) in buffer" :rules="rules" :ref="(ref) => rule_form[idx] = ref" :model="elt">
                <div class="load-item">
                    <el-form-item prop="load_date" label="日期" label-width="120px">
                        <el-date-picker type="date" v-model="elt.load_date" :disabled-date="isnt_future" />
                    </el-form-item>

                    <el-form-item prop="loads" label="件数" label-width="120px">
                        <el-input-number :min="1" :controls="false" v-model.number="elt.loads" />
                    </el-form-item>

                    <el-form-item prop="load_ton" label="吨数" label-width="120px">
                        <el-input-number :min="1.00" :controls="false" :precision="2" v-model="elt.load_ton" />
                    </el-form-item>

                    <el-form-item label-width="120px">
                        <el-button class="delete-button" :icon="Delete" @click.prevent="remove_load(elt)" />
                    </el-form-item>

                </div>
            </el-form>

            <div style="margin-left:120px">
                <el-button @click="add_load" :icon="Plus" />
                <el-button type="primary" @click="submitForm(rule_form)" :disabled="buffer.length == 0">
                    提交
                </el-button>
            </div>

        </div>
    </Component>
</template>

<style lang="scss" scoped>
.delete-button {
    background-color: #FF3030;
    color: #FFFFFF;
}

.load-item {
    margin-bottom: 40px;
}
</style>
