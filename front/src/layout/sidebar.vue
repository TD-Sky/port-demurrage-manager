<script lang="ts" setup>
// import barStore from '../store';
// import { storeToRefs } from 'pinia';
import { Document, Menu as IconMenu, Setting, } from '@element-plus/icons-vue';

//let { collapsed } = storeToRefs(barStore);

</script>

<template>
    <!-- <el-menu default-active="1" class="siderbar" :collapse="collapsed" @open="handleOpen" @close="handleClose"> -->
    <el-menu default-active="this.$route.path" class="siderbar" router>

        <el-menu-item index="/store">
            <el-icon>
                <icon-menu />
            </el-icon>
            <span slot="title">仓储</span>
        </el-menu-item>

        <el-menu-item index="/load">
            <el-icon>
                <document />
            </el-icon>
            <span slot="title">订单</span>
        </el-menu-item>

        <el-menu-item index="/fee">
            <el-icon>
                <setting />
            </el-icon>
            <span slot="title">堆存费</span>
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
