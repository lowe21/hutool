<script setup lang="ts">
import { onMounted } from 'vue'
import Emitter from '@/utils/Emitter'
import ChartComponent from '@/components/ChartComponent.vue'
import ResultComponents from "@/components/ResultComponents.vue";

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

    let result = Math.random() * 2
    if (result > 1) {
      result = -1
    } else {
      result = Math.round(result)
    }
    Emitter.emit('change-result', result)
  }, 1000)
})
</script>

<template>
  <el-container>
    <el-main>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
      <chart-component></chart-component>
    </el-main>
    <el-aside>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
      <result-components></result-components>
    </el-aside>
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
  padding: 20px;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
