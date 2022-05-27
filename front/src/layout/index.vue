<script setup lang="ts">
import Sidebar from './sidebar.vue'
import Header from './header.vue'
</script>

<template>
    <div class="common-layout">
        <el-container>
            <el-aside>
                <Sidebar />
            </el-aside>
            <el-container>
                <el-header>
                    <Header />
                </el-header>
                <el-main>Main</el-main>
            </el-container>
        </el-container>
    </div>
</template>
