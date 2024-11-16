<script setup lang="ts">
import { ref, computed, onUnmounted } from 'vue'
import Emitter from '@/utils/Emitter'

const result = ref(0)

const detecting = computed(() => {
  return result.value == 0
})

const passed = computed(() => {
  return result.value == 1
})

const notPass = computed(() => {
  return result.value == -1
})

Emitter.on('change-result', (data:any) => {
  result.value = data
})

onUnmounted(() => {
  Emitter.off('change-result')
})
</script>

<template>
  <el-card>
    <template #header>检测结果</template>
    <el-result v-show="detecting" icon="info" title="检测中"></el-result>
    <el-result v-show="passed" icon="success" title="合格"></el-result>
    <el-result v-show="notPass" icon="error" title="不合格"></el-result>
  </el-card>
</template>

<style scoped>
  .el-card {
    height: 400px;
  }
</style>
