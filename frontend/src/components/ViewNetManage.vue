<template>
    <main>
        <el-button type="primary" @click="dialogVisible = true">
            {{ title }}
        </el-button>
        <el-dialog
            v-model="dialogVisible"
            :title="title"
            width="60%"
        >
            <span>This is a message</span>
            <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogVisible = false">Cancel</el-button>
                <el-button type="primary" @click="dialogVisible = false">
                Confirm
                </el-button>
            </span>
            </template>
        </el-dialog>
    </main>
</template>

<script setup>
import { ref } from 'vue'
defineProps(['title'])
const dialogVisible = ref(false)
</script>


<style scoped>
</style>