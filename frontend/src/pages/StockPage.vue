<template>
  <PageLayout>
    <div class="plate pad">
      <div>
        <div class="flex items-center">
          <div class="text-xl mr-2">{{ route.params.stockId }}</div>
          <div class="flex items-center mr-1">
            <ArrowIcon/>
            <div class="font-bold text-3xl">212,66</div>
          </div>
          <div class="text-xs flex flex-col">
            <div>+0,77%</div>
            <div>RUB</div>
          </div>
        </div>
      </div>

      <div class="mt-10 flex justify-center">
        <CandlesWithPredictionChart class="w-[800px] h-[500px]" />
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import CandlesWithPredictionChart from '../components/trading/CandlesWithPredictionChart.vue'
import PageLayout from '../layouts/PageLayout.vue'
import ArrowIcon from '../assets/arrow-drop-down.svg'

const route = useRoute()

</script>

<style scoped>

</style>