<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue'
import { Chart } from '@antv/g2'
import { type HandlerDataArray } from '@/types/Websocket'
import Emitter from '@/utils/Emitter'

const chart = new Chart({
  autoFit: true,
  height: 400,
})

const line = chart.data([])

chart
  .line()
  .encode('x', 'x')
  .encode('y', 'y')
  .encode('color', 'line')
  .encode('shape', 'smooth')
  .scale('x', { type: 'linear', domain: [0, 9], tickMethod: () => [0, 1, 2, 3, 4, 5, 6, 7, 8, 9] })
  .scale('y', { type: 'linear', domain: [0, 255], tickMethod: () => [0, 25, 50, 75, 100, 125, 150, 175, 200, 225, 255] })
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
