<script setup>
import ViewNet from './components/ViewNet.vue'
import ViewNetManage from './components/ViewNetManage.vue';
import ViewRequest from './components/ViewRequest.vue';
</script>

<template>
  <el-row>
    <el-col :span="6">
      <div style="margin-bottom: 12px;">
        <ViewNetManage title="开启请求"/>
      </div>
      <ViewNet/>
    </el-col>
    <!-- 右侧列，使用 offset 属性进行偏移 -->
    <el-col :span="18">
      <ViewRequest/>
    </el-col>
  </el-row>
</template>
