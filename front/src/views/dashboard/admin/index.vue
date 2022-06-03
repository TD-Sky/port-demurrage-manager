<script setup lang="ts">
import { Storage } from '@/models/index';
import { request } from '@/utils/service';
import { onBeforeMount, ref } from "vue";

const storages = ref<Storage[]>([])

onBeforeMount(() => {
    request({
        url: "/store",
        method: "get",
    }).then((resp) => {
        storages.value = resp.data.stores
    })
})
</script>

<template>
    <div class="app-container">
        <el-row :gutter="50">
            <el-col :span="50" v-for="storage in storages">
                <el-card class="box-card">
                    <template #header>
                        <div class="card-header">
                            <span>ID：{{ storage.id }}</span>
                            <el-button class="button" text>Operation button</el-button>
                        </div>
                    </template>

                    <div class="text item">
                        入库日期：{{ storage.store_date }}
                    </div>

                    <div class="text item">
                        车牌号: {{ storage.license_plate_number }}
                    </div>

                    <div class="text item">
                        件数：{{ storage.stocks }}
                    </div>

                    <div class="text item">
                        吨数：{{ storage.store_ton }}
                    </div>

                    <div class="text item">
                        仓储天数：{{ storage.duration }}
                    </div>

                    <div class="text item">
                        堆存费：{{ storage.fee }}
                    </div>
                </el-card>
                <div class="empty"></div>
            </el-col>
        </el-row>
    </div>
</template>

<style>
.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.text {
    font-size: 14px;
}

.item {
    margin-bottom: 18px;
}

.box-card {
    width: 400px;
}

.empty {
    height: 25px;
}
</style>
