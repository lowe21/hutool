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

chart.line()
  .encode('x', 'x')
  .encode('y', 'y')
  .encode('shape', 'smooth')
  .scale('x', { type: 'linear', domain: [0, 29], tickMethod: () => [5, 10, 15, 20, 25] })
  .scale('y', { type: 'linear', domain: [0, 10], tickMethod: () => [0, 2.5, 5, 7.5, 10] })
  .animate('enter', { type: 'pathIn', duration: 1000 })

chart.axisX().attr('title', '')
chart.axisY().attr('title', '')

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
