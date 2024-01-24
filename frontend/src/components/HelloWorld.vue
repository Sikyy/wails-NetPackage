<script setup>
import {reactive, ref} from 'vue'
import {CaptureTraffic} from '../../wailsjs/go/main/App'
import {StopCaptureTraffic} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

let code = ref()
let list = ref()
let startbut = ref()
let stopbut = ref()


function greet() {
  CaptureTraffic().then(() => {
    console.log("start")
    startbut.value = "开始监听"
  })
}

function stop() {
  StopCaptureTraffic().then(() => {
    console.log("stop")
    stopbut.value = "停止监听"
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
      <button class="btn" @click="stop">Stop</button>
      <div>code : {{ code }}</div>
      <div>list : {{ list }}</div>
      <div>start：{{ startbut }}</div>
      <div>stop：{{ stopbut }}</div>
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
