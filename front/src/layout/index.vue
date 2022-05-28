<script setup lang="ts">
import Sidebar from './sidebar.vue'
import Topbar from './topbar.vue'
</script>

<template>
    <div class="common-layout">
        <el-container>
            <el-aside>
                <Sidebar />
            </el-aside>
            <el-container>
                <el-header>
                    <Topbar />
                </el-header>
                <el-main>
                    <slot />
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>
