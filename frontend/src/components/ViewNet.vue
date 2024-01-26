<template>
    <div class="demo-collapse">
    <el-collapse v-model="activeNames" @change="handleChange">
      <el-collapse-item v-for="item in list" name="1">
        <template #title>
            {{ item }}
        </template>
      </el-collapse-item>
    </el-collapse>
    </div>
    <!-- <div>
        <el-table :data="tableData" stripe height="250" style="width: 100%">
          <el-table-column prop="date" label="日期" width="100" />
          <el-table-column prop="clientname" label="客户端" width="100" />
          <el-table-column prop="status" label="状态" width="100" />
          <el-table-column prop="upload" label="上传" width="100" />
          <el-table-column prop="download" label="下载" width="100" />
          <el-table-column prop="still_time" label="时长" width="100" />
          <el-table-column prop="method" label="方法" width="100" />
          <el-table-column prop="host" label="地址" />
        </el-table>
    </div> -->
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