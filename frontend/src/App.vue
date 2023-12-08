<template>
  <StocksPage/>
</template>

<script lang="ts" setup>
import { onMounted } from 'vue'
import { getAllStocks } from './api'
import StocksPage from './components/pages/StocksPage.vue'

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