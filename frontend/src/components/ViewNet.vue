<template>
    <div class="demo-collapse">
    <el-collapse v-model="activeNames" @change="handleChange">
      <el-collapse-item v-for="item in list" name="1">
        <template #title>
            {{ item }}
        </template>
      </el-collapse-item>
      <el-collapse-item title="Controllability" name="4">
        <div>
          Controlled consequences: users should be granted the freedom to
          operate, including canceling, aborting or terminating current
          operation.
        </div>
      </el-collapse-item>
    </el-collapse>
  </div>
</template>

<script setup>

import { ref } from 'vue'
import { ElNotification } from 'element-plus'
import { Networklist } from '../../wailsjs/go/main/App'
let list = ref()
Networklist().then(result => {
    if (result.code !== 200){
        ElNotification({
            title: '错误',
            type: 'error'
        })
    }
    list.value = result.data
})
</script>


<style scoped>
</style>