<script lang="ts" setup>
import { toRefs } from 'vue';

// 共享状态
const props = defineProps(["opening", "remove_code"]);
const { opening, remove_code } = toRefs(props);

const emits = defineEmits(["close_form", "delete_then_refresh"]);

</script>

<template>
    <el-dialog v-model="opening['delete']" title="移除公司" width="30%">
        <span>确定删除吗？</span>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="emits('close_form', 'delete')">
                    取消
                </el-button>
                <el-button type="primary" @click="emits('delete_then_refresh', remove_code)">
                    确定
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>

<style scoped>
.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
