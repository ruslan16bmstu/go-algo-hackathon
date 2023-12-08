<template>
  <svg :width="width" :height="height">
    <polyline v-for="(line, index) in linesData" :key="index"
              :points="convertPointsToSVGPolyline(line.points, width, height)"
              :stroke="line.color" :stroke-width="line.width" fill="none">
    </polyline>
  </svg>
</template>

<script setup lang="ts">
import { ref, onMounted, defineProps, computed } from 'vue'
import { findMinMax, Point, scalePointsToCanvas } from '../../utils'

interface Line {
  points: Point[]
  color: string
  width: number
}

const props = defineProps<{
  height: number
  width: number
  linesData: Line[]
}>()

const pointsStr = ref<string>('')

const minMax = computed(() => {
  let allPoints: Point[] = []
  props.linesData.forEach(({points}) => {allPoints = allPoints.concat(points)})
  return findMinMax(allPoints)
})

const convertPointsToSVGPolyline = (points: Point[], width: number, height: number) => {
  return scalePointsToCanvas(points, width, height, minMax.value).map(p => {
    return `${p.x},${p.y}`
  }).join(' ')
}

onMounted(() => {
})

</script>
