<template>
  <div class="w-full">
    <div class="flex mb-6 justify-center bg-white shadow-md">
      <TheHeader class="max-w-[1064px]"/>
    </div>
    <div class="w-full flex justify-center">
      <div class="flex flex-col space-y-6 w-[1000px]">
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
      </div>
    </div>
    <TheFooter/>
  </div>
</template>

<script lang="ts" setup>
import { computed, ref } from 'vue'
import TheFooter from '../components/layout/TheFooter.vue'
import TheHeader from '../components/layout/TheHeader.vue'
import DropdownWithLabel from '../components/shared/DropdownWithLabel.vue'
import TitlePlate from '../components/shared/TitlePlate.vue'
import SharesPredict from '../components/trading/SharesPredict.vue'

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

</script>

<style scoped>
</style>