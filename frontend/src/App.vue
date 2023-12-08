<template>
  <router-view/>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { getAllStocks } from './api'

async function fetchAndSaveStocks() {
  const allStocks = await getAllStocks()
  console.log('Список всех акций:', allStocks)
}

onMounted(() => {
  // fetchAndSaveStocks()
})
</script>

<style>
</style>