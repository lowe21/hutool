<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { Chart } from '@antv/g2'
import { type HandlerDataArray } from '@/types/Websocket'
import Emitter from '@/utils/Emitter'

const chart = new Chart({
  autoFit: true,
  height: 400,
})

const line = chart.data([
  {x: 0, y: 65},
  {x: 1, y: 40},
  {x: 2, y: 95},
  {x: 3, y: 60},
  {x: 4, y: 80},
  {x: 5, y: 55},
  {x: 6, y: 70},
  {x: 7, y: 29},
  {x: 8, y: 40},
  {x: 9, y: 85},
]);

chart
  .line()
  .encode('x', 'x')
  .encode('y', 'y')
  .encode('shape', 'smooth')
  .scale('x', { type: 'linear', domain: [0, 9], tickMethod: () => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] })
  .scale('y', { type: 'linear', domain: [0, 100], tickMethod: () => [0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100] })
  .animate('enter', { type: 'pathIn', duration: 1000 })

onMounted(() => {
  const element = document.getElementById('container')
  if (element) {
    element.append(chart.getContainer())
    chart.render()
  }
})

Emitter.on('change-data', (data: HandlerDataArray) => {
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
