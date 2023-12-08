<template>
  <PageLayout>
    <TitlePlate title="Прогнозы"/>
    <div>
      <div class="plate">
        <div>
          <div class="font-bold pad">Топ на покупку / продажу</div>
        </div>
        <div class="border-t">
          <div class="pad">
            <DropdownWithLabel v-model="selectedIndustryIndex" label="Отрасль:" :options="industries"
                               class="text-sm"/>
          </div>
        </div>
        <div class="border-t">
          <SharesPredict :stock-predictions="stockWithPredictionArr"/>
        </div>
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { computed, onMounted, ref } from 'vue'
import DropdownWithLabel from '../components/shared/DropdownWithLabel.vue'
import TitlePlate from '../components/shared/TitlePlate.vue'
import SharesPredict from '../components/trading/StockPredictions.vue'
import type { StockWithPrediction } from '../components/trading/types'
import PageLayout from '../layouts/PageLayout.vue'
import { useStockStore } from '../stores/stock'

const selectedIndustryIndex = ref(0)

const industries = [
  'Все',
  'Металлы и добыча',
  'Недвижимость',
  'Нефть и газ',
  'Потребительский сектор',
  'Промышленность',
  'Телекоммуникации',
  'Транспорт',
  'Финансы',
  'Химия и нефтехимия',
  'Электроэнергетика',
  'IT-сектор',
]

const selectedIndustry = computed(() => industries[selectedIndustryIndex.value])

const stockStore = useStockStore()

onMounted(() => {
  stockStore.loadData()
})

const {data} = storeToRefs(useStockStore())

const maxAbsDelta = 5
const stockWithPredictionArr = computed(() => {
  return data.value.map(d => {
    const delta = +(Math.random() * 2 * maxAbsDelta - maxAbsDelta).toFixed(3)
    return {...d, delta, points: delta > 0 ? predictedExampleProfit : predictedExampleLoose} as StockWithPrediction
  })
})

const predicExampleYProfit = [
  553,
  551,
  558,
  558,
  558.6,
  558.6,
  564,
  563.6,
  563.2,
  561.8,
  561.4,
  562.2,
  562.2,
  564,
  566.4,
  565.4,
  566.4,
  565,
  567,
  563.8,
  560.8,
  559,
  559.4,
  563,
  562,
  565,
  565.8,
  565.4,
  562.2,
  563.6,
  563,
  563,
  561.6,
  561.8,
  562,
  562.6,
  561.8,
  560,
  561.2,
  561.2,
  561.2,
  561.6,
  563.8,
  564.6,
  564.6,
  564.6,
  565.6,
  567,
  567.2,
  568.8,
  567.6,
  567.4,
  567,
  567.2,
  567.4,
  564.4,
  564.8,
  563,
  561.4,
  562.8,
  561.4,
  562.4,
  564,
  565.4,
  565,
  564,
  565,
  565.2,
  565,
  564.2,
  563.6,
  564.6,
  565.4,
  566.4,
  567.4,
  569,
  569.6,
  568.8,
  568.6,
  569.2,
  569,
  568.8,
  568.6,
  567.6,
  567.2,
  567.4,
  567.4,
  567,
  567.4,
  567.2,
  566.6,
  566.6,
  566.8,
  567.2,
  568.2,
  567.2,
  569.2,
  569.8,
  568.6,
  568.4,
  567.6,
  566.4,
  567.4,
  566.2,
  566.6,
  568.6
]

const predictedExampleProfit = predicExampleYProfit.map((y, x) => ({x, y}))
const predictedExampleLoose = predicExampleYProfit.reverse().map((y, x) => ({x, y}))

</script>

<style scoped>
</style>