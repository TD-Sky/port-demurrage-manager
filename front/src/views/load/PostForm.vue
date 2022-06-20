<script lang="ts" setup>
import { PostShippingOrder, PostLoad } from '@/models/load';
import { FreightForwarder } from '@/models/freight_forwarder';
import { reactive, ref } from 'vue';
import { isnt_future } from '@/utils/index';
import { Plus, Delete } from '@element-plus/icons-vue';
import { get_dataset, post_data } from '@/api/crud';
import { useRouter } from 'vue-router';

const ROUTER = useRouter();
const FIRST_COL = "50px";

// 表单所绑定缓冲区
const buffer = reactive<PostShippingOrder>(
    {
        company_code: "",
        loads: [
            <PostLoad>{}, // 没有字段信息的
        ],
    }
);

// 下拉框内可选择的公司
const companies = ref<{
    code: string,
    name: string,
}[]>([]);

// 添加新出货条目
const push_load = () => {
    buffer.loads.push(<PostLoad>{});
}

// 移除指定位置的出货条目
const remove_load = (index: number) => {
    buffer.loads.splice(index, 1);
}

// 提交表单
const submit_order = () => {
    // Reactive 可以直接传
    post_data("/load", buffer).then((_) => {
        ROUTER.push("/load");
    });
}

// 获取所有货代公司
get_dataset("/freight_forwarder").then((resp) => {
    if (resp.data.companies !== null) {
        companies.value = (resp.data.companies as FreightForwarder[])
            .map(ff => ({
                code: ff.code,
                name: ff.company_name,
            }));
    } else {
        companies.value = [];
    }
});
</script>

<template>
    <Component>
        <div class="app-container">

            <el-form :model="buffer" inline>
                <!-- 每个div是一行 -->

                <div>
                    <el-form-item label="公司" :label-width="FIRST_COL">

                        <el-select v-model="buffer.company_code" placeholder="xxxxx">
                            <!-- 渲染label，赋值value -->
                            <el-option v-for="cpn in companies" :key="cpn.code" :value="cpn.code" :label="cpn.name" />
                        </el-select>

                    </el-form-item>
                </div>

                <div v-for="(load, index) in buffer.loads">
                    <el-form-item label="日期" :label-width="FIRST_COL">
                        <el-date-picker type="date" v-model="load.load_date" :disabled-date="isnt_future" />
                    </el-form-item>

                    <el-form-item label="数量">
                        <el-input-number :min="1" :controls="false" v-model.number="load.loads" />
                    </el-form-item>

                    <el-form-item>
                        <el-button type="danger" :icon="Delete" @click.prevent="remove_load(index)" />
                    </el-form-item>
                </div>

                <div>
                    <!-- 以空内容利用对齐 -->
                    <el-form-item label=" " :label-width="FIRST_COL">
                        <el-button :icon="Plus" @click="push_load" />
                    </el-form-item>

                    <el-form-item label=" " :label-width="FIRST_COL">
                        <el-button type="primary" @click="submit_order" :disabled="buffer.loads.length == 0">
                            提交
                        </el-button>
                    </el-form-item>
                </div>

            </el-form>

        </div>
    </Component>
</template>
