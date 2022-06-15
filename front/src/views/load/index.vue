<script setup lang="ts">
import { reactive, onBeforeMount, ref, computed } from 'vue';
import { get_loads, put_load, delete_load } from './request';
import { Edit, Plus, Delete } from '@element-plus/icons-vue';
import { GetLoad, PutLoad } from '@/models/index';
import { useRouter } from 'vue-router';
import PutForm from './PutForm.vue';
import DelDialog from './DelDialog.vue';

const router = useRouter();

/* 状态 */
const loads = ref<GetLoad[]>([]);                           // 出货单
const buffer = reactive(<PutLoad>{});                       // 表单修改缓冲区
const remove_id = ref<number>();                            // 删除条目的ID
const opening = reactive<{ [key: string]: boolean }>({
    "put": false,
    "delete": false,
});                                                         // 会话开关

const build_table = () => {
    get_loads().then((resp) => {
        loads.value = resp.data.loads;
    });
}

// 必须用闭包，
// 模板部分无法访问完整的响应类型，不能传响应类型值
const modify_form = (row: GetLoad) => {
    buffer.id = row.id;
    buffer.loads = row.loads;
    buffer.load_date = new Date(row.load_date);

    opening["put"] = true
}

const new_form = () => {
    router.push("/load/post")
}

const remove_item = (id: number) => {
    remove_id.value = id
    opening["delete"] = true
}

// 子组件修改父组件状态不会生效！！！
// 必须靠 emits 触发此函数
const close_form = (kind: string) => {
    opening[kind] = false
}

function put_then_refresh(data: PutLoad) {
    put_load(data).then((_) => {
        build_table();
    })
}

function delete_then_refresh(id: number) {
    delete_load(id).then((_) => {
        opening["delete"] = false
        build_table();
    })
}

const total_fee = computed(() => {
    // 迭代器写法会卡刷新
    let sum = 0

    for (const item of loads.value) {
        sum += item.fee
    }

    return sum;
});

// 进入网页，先读一次数据库
onBeforeMount(() => {
    build_table();
})
</script>

<template>
    <Component>
        <div class="app-container">

            <div class="loads_table">
                <el-table :data="loads" border :fit="false" max-height="1000">
                    <el-table-column header-align="center" prop="load_date" label="日期" width="120" />
                    <el-table-column header-align="center" prop="loads" label="件数" width="120" />
                    <el-table-column header-align="center" prop="load_ton" label="吨数" width="120" />
                    <el-table-column align="center" prop="business_number" label="业务号" width="150" />
                    <el-table-column align="center" prop="lading_bill_number" label="提单号" width="180" />
                    <el-table-column align="center" prop="order_number" label="订单号" width="150" />
                    <el-table-column header-align="center" prop="fee" label="堆存费" width="120" />
                    <el-table-column align="center" label="操作" width="180">
                        <template #default="scope">
                            <el-button type="primary" @click="modify_form(scope.row)" :icon="Edit" />
                            <el-button type="danger" @click="remove_item(scope.row.id)" :icon="Delete" />
                        </template>
                    </el-table-column>
                </el-table>
            </div>

            <div class="total_fee">
                <el-descriptions border>
                    <el-descriptions-item label="预付堆存费" width="840px">
                        {{ total_fee }}
                    </el-descriptions-item>
                </el-descriptions>
            </div>
        </div>

        <div class="handle-button" @click="new_form">
            <el-icon :size="24">
                <Plus />
            </el-icon>
        </div>

        <PutForm :opening="opening" :buffer="buffer" @close_form="close_form" @put_then_refresh="put_then_refresh" />

        <DelDialog :opening="opening" :remove_id="remove_id" @close_form="close_form"
            @delete_then_refresh="delete_then_refresh" />

    </Component>
</template>

<style lang="scss" scoped>
.handle-button {
    width: 48px;
    height: 48px;
    background-color: #5f9ea0;
    position: absolute;
    right: 0px;
    text-align: center;
    font-size: 24px;
    border-radius: 6px 0 0 6px !important;
    z-index: 10;
    cursor: pointer;
    pointer-events: auto;
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: center;
    top: 300px;
}

.loads_table {
    width: 85.4%;
}

.total_fee {
    width: 85.3%;
}
</style>
