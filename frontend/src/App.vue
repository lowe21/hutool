<script setup lang="ts">
import { onMounted } from 'vue'
import connection from '@/utils/Websocket'
import ChartComponent from '@/components/ChartComponent.vue'
import Chart2Component from '@/components/Chart2Component.vue'
import ResultComponents from "@/components/ResultComponents.vue"

onMounted(() => {
  connection()
})
</script>

<template>
  <el-container>
    <el-header>
      <img alt="Vue logo" class="logo" src="./assets/logo.svg" width="75" height="75" />
    </el-header>
    <el-container>
      <el-main>
        <section>
          <chart-component></chart-component>
        </section>
        <section>
          <chart2-component></chart2-component>
        </section>
      </el-main>
      <el-aside>
        <result-components></result-components>
      </el-aside>
    </el-container>
  </el-container>
</template>

<style scoped>
.el-container {
  height: 100vh;
  line-height: 1.5;
}

.el-aside {
  padding: 20px;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}
</style>
