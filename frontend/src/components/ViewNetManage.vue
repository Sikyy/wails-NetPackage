<template>
    <main>
        <el-button type="primary" @click="toggleTraffic">
            {{ title }}
            <!-- {{ trafficStatus ? '开始监听' : '停止监听' }} -->
        </el-button>
    </main>
</template>

<script setup>
import { ref } from 'vue'
import { CaptureTraffic, StopCaptureTraffic } from '../../wailsjs/go/main/App'
defineProps(['title'])

const trafficStatus = ref(false)

function toggleTraffic() {
  if (trafficStatus.value) {
    // 如果当前是开启状态，则执行停止监听方法
    StopCaptureTraffic().then(() => {
      console.log("stop")
      trafficStatus.value = false
    })
  } else {
    // 如果当前是关闭状态，则执行开始监听方法
    CaptureTraffic().then(() => {
      console.log("start")
      trafficStatus.value = true
    })
  }
}
</script>

<style scoped>
</style>
