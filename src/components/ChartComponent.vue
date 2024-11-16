<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import Emitter from '@/utils/Emitter'
import { Chart } from '@antv/g2'

const chart = new Chart({
  autoFit: true,
  height: 400,
})

const line = chart.data({
  type: 'fetch',
  value: 'https://assets.antv.antgroup.com/g2/doughnut-purchases.json',
});

chart
  .line()
  .encode('x', 'year')
  .encode('y', 'count')
  .encode('shape', 'smooth')
  .scale('y', { nice: true })
  .animate('enter', { type: 'pathIn', duration: 1000 })

onMounted(() => {
  const element = document.getElementById('container')
  if (element) {
    element.append(chart.getContainer())
    chart.render()
  }
})

Emitter.on('change-data', (data) => {
  line.changeData(data)
})

onUnmounted(() => {
  chart.destroy()
  Emitter.off('change-data')
})
</script>

<template>
  <div id="container"></div>
</template>

<style scoped>
#container {
  background-color: #f8f8f8;
}
</style>
