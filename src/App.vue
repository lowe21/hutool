<script setup lang="ts">
import {reactive, onMounted, onUnmounted} from 'vue'
import G2Chart from '@/components/G2Chart.vue'

let timer:number
let data = reactive([{
  value: 0,
  timestamp: 0
}])

onMounted(() => {
  timer = setInterval(() => {
    data = Object.assign([{
      value: Math.random() * (1000 - 0) + 1000,
      timestamp: Date.now()
    }])
  }, 1000)
})

onUnmounted(() => {
  clearInterval(timer)
})
</script>

<template>
  <el-container>
    <el-main>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="125" height="125" />
      <G2Chart :data="data"/>
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
