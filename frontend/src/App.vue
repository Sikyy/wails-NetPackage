<script setup>
import HelloWorld from './components/HelloWorld.vue'
import ViewNet from './components/ViewNet.vue'
</script>

<template>
  <el-row>
    <el-col :span="6" style="height: 100vh;">
      <ViewNet/>
    </el-col>
  </el-row>
  <!-- <img id="logo" alt="Wails logo" src="./assets/images/logo-universal.png"/>
  <HelloWorld/> -->
</template>
