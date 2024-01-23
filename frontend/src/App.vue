<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ViewNet from './components/ViewNet.vue'
import ViewNetManage from './components/ViewNetManage.vue';
</script>

<template>
  <el-row>
    <el-col :span="6" style="height: 100vh; margin: 12px;">
      <div style="margin-bottom: 12px;">
        <ViewNetManage title="开启请求"/>
      </div>
      <ViewNet/>
    </el-col>
  </el-row>
  <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/> -->
</template>
