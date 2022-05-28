<script  setup lang="ts">
// import barStore from '../store';
import { Fold } from '@element-plus/icons-vue';

// function toggle_sidebar() {
//     barStore.$patch({
//         collapsed: !barStore.collapsed,
//     })
// }
</script>

<template>
    <header>
        <el-button plain>
            <el-icon>
                <Fold />
            </el-icon>
        </el-button>
    </header>
</template>
