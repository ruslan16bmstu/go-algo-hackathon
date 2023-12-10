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
            <DropdownWithLabel v-model="selectedIndustryIndex" label="Отрасль:"
                               :options="industries.map(({name}) => name)"
                               class="text-sm"/>
          </div>
        </div>
        <div class="border-t">
          <SharesPredict :stock-predictions="filteredData"/>
        </div>
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { storeToRefs } from 'pinia'
import { computed, onMounted, onUnmounted, ref, watch } from 'vue'
import DropdownWithLabel from '../components/shared/DropdownWithLabel.vue'
import TitlePlate from '../components/shared/TitlePlate.vue'
import SharesPredict from '../components/trading/StockPredictions.vue'
import PageLayout from '../layouts/PageLayout.vue'
import { useStockStore } from '../stores/stock'
import 'nprogress/nprogress.css'
import NProgress from 'nprogress'

const selectedIndustryIndex = ref(0)

const industries = [
  {"id": 0, "index": "", name: "Все"},
  {"id": 1, "index": "MOEXEU", name: "Электроэнергетика"},
  {"id": 2, "index": "MOEXTL", name: "Телекоммуникации"},
  {"id": 3, "index": "MOEXMM", name: "Металлы и добыча"},
  {"id": 4, "index": "MOEXCN", name: "Потребительский сектор"},
  {"id": 5, "index": "MOEXRE", name: "Строительные компании"},
  {"id": 6, "index": "MOEXOG", name: "Нефть и газ"},
  {"id": 7, "index": "MOEXFN", name: "Финансы"},
  {"id": 8, "index": "MOEXCH", name: "Химия и нефтехимия"},
  {"id": 9, "index": "MOEXIT", name: "Информационные технологии"},
  {"id": 10, "index": "MOEXTN", name: "Транспорт"},
]

const selectedIndustry = computed(() => industries[selectedIndustryIndex.value])

const stockStore = useStockStore()
const {isLoading} = storeToRefs(stockStore)

watch(isLoading, () => {
  if (isLoading.value) {
    NProgress.start()
  } else {
    NProgress.done()
  }
})

onMounted(() => {
  if (data.value.length < 2) {
    stockStore.loadData()
  }
})

onUnmounted(() => {
  NProgress.done()
})

const {data} = storeToRefs(useStockStore())

const filteredData = computed(() => {
  if (selectedIndustryIndex.value == 0) {
    return data.value
  }

  return data.value.filter(d => d.industry == selectedIndustry.value.id)
})
</script>

<style scoped>
</style>