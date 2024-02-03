<template>
    <el-table :data="tableData" stripe style="width: 100%" max-height="1245">
      <el-table-column prop="id" label="会话ID" width="100" />
      <el-table-column prop="date" label="日期" width="100" />
      <el-table-column prop="clientname" label="客户端" width="100" />
      <el-table-column prop="status" label="状态" width="100" />
      <el-table-column prop="upload" label="上传" width="100" />
      <el-table-column prop="download" label="下载" width="100" />
      <el-table-column prop="still_time" label="时长" width="100" />
      <el-table-column prop="method" label="方法" width="100" />
      <el-table-column prop="host" label="地址" />
  </el-table>
</template>
  
  <script lang="ts" setup>
  import { ref, onMounted } from 'vue';
  import { EventsOn } from '../../wailsjs/runtime/runtime';
  
  const tableData = ref([
    {
      id: '',
      date: '',
      clientname: '',
      status: '',
      upload: '',
      download: '',
      still_time: '',
      method: '',
      host: '',
    },
  ]);
  
  onMounted(() => {
    GetTraffic();
  });
  
  function GetTraffic() {
  EventsOn('captureTraffic', (data) => {
    console.log("接收到流量数据:", data);

    // 检查数据的ID是否已经存在于tableData中
    const existingIndex = tableData.value.findIndex(item => item.id === data.ID);

    // 如果 method 不是 "UDP"，才进行添加或更新操作
    if (data.Method !== "UDP" && data.Method !== "TCP"  && data.Host !== ":443") {//&& data.Method !== "TCP"
      if (existingIndex !== -1) {
        // 如果存在，更新现有数据
        tableData.value[existingIndex] = {
          id: data.ID,
          date: data.Time_s,
          clientname: data.ClientName,
          status: data.Status,
          upload: data.SessionUpTraffic,
          download: data.SessionDownTraffic,
          still_time: data.Length_of_time + "ms",
          method: data.Method,
          host: data.Host,
        };
      } else {
        // 如果不存在，添加新数据
        tableData.value.push({
          id: data.ID,
          date: data.Time_s,
          clientname: data.ClientName,
          status: data.Status,
          upload: data.SessionUpTraffic,
          download: data.SessionDownTraffic,
          still_time: data.Length_of_time + "ms",
          method: data.Method,
          host: data.Host,
        });
      }
    }
  });
}
</script>