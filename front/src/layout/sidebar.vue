<script lang="ts" setup>
import { ref } from 'vue'
import { Document, Menu as IconMenu, Setting, } from '@element-plus/icons-vue'

const isCollapse = ref(false)
const handleOpen = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
const handleClose = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
}
</script>

<template>
    <el-menu default-active="1" class="siderbar" :collapse="isCollapse" @open="handleOpen" @close="handleClose">

        <el-menu-item index="1">
            <el-icon>
                <icon-menu />
            </el-icon>
            <template #title>仓储</template>
        </el-menu-item>

        <el-menu-item index="2">
            <el-icon>
                <document />
            </el-icon>
            <template #title>订单</template>
        </el-menu-item>

        <el-menu-item index="3">
            <el-icon>
                <setting />
            </el-icon>
            <template #title>堆存费</template>
        </el-menu-item>

    </el-menu>
</template>

<style>
.siderbar:not(.el-menu--collapse) {
    width: 150px;
    height: 100%;
    border: none;
}
</style>
