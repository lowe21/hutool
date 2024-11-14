<script setup lang="ts">
import { onMounted } from 'vue'
import ChartComponent from '@/components/ChartComponent.vue'
import Emitter from '@/utils/Emitter'

onMounted(() => {
  setInterval(() => {
    const data = []
    for (let i = 1996; i <= 2024; i++) {
      data.push({
        year: i,
        count: Math.floor(Math.random() * 10000),
      })
    }
    Emitter.emit('change-data', data)
  }, 1000)
})
</script>

<template>
  <el-container>
    <el-main>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
      <chart-component></chart-component>
    </el-main>
    <el-aside></el-aside>
  </el-container>
</template>

<style scoped>
.el-container {
  height: 100vh;
  line-height: 1.5;
}

.el-main {
  background-color: green;
}

.el-aside {
  background-color: purple;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
