<template>
    <el-table :data="tableData" stripe style="width: 100%">
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
      id: '1',
      date: '17:00:00',
      clientname: 'Telegram',
      status: '已完成',
      upload: '100M',
      download: '100M',
      still_time: '200ms',
      method: 'GET',
      host: '192.168.50.1',
    },
  ]);
  
  onMounted(() => {
    GetTraffic();
  });
  
  function GetTraffic() {
    EventsOn('captureTraffic', (data) => {
      console.log("Received traffic data:", data);
  
      // 将接收到的数据添加到 tableData 的末尾{
      tableData.value.push({
        id :data.ID,
        date :data.Time_s,
        clientname: "",
        status: data.Status,
        upload: data.SessionUpTraffic,
        download: data.SessionDownTraffic,
        still_time: data.Length_of_time,
        method: data.Method,
        host: data.Host,
      });
    });
  }
  </script>