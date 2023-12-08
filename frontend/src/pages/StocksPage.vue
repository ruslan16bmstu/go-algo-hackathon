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
          <SharesPredict/>
        </div>
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import DropdownWithLabel from '../components/shared/DropdownWithLabel.vue'
import TitlePlate from '../components/shared/TitlePlate.vue'
import SharesPredict from '../components/trading/StockPredict.vue'
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

</script>

<style scoped>
</style>