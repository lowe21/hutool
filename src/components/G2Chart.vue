<script setup lang="ts">
import { defineProps, watch, watchEffect, onMounted } from 'vue'
import { Chart } from '@antv/g2';
import { type ChartDatas } from '@/types/ChartDatas'

//let datas:ChartDatas = []
const { data } = defineProps<{data?:ChartDatas}>()

watch([() => data], (value) => {
  console.log(value)
})

watchEffect(() => {
  if (data != null) {
    console.log('数组变化了', data[0].timestamp)
  }
});

onMounted(()=> {
  const chart = new Chart({
    container: 'container',
    autoFit: true,
  })

  chart.line().data({
    type: 'custom',
    callback: () => {
      return []
    },
  }).encode('x', 'timestamp').encode('y', 'value').transform({
    type: 'sample',
    thresholds: 200,
    strategy: 'max'
  })

  chart.render()
})
</script>

<template>
  <div id="container"></div>
</template>

<style scoped>
#container {
  background-color: white;
}
</style>
