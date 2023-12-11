<template>
  <div class="w-full">
    <div class="flex mb-6 justify-center bg-white shadow-md">
      <TheHeader class="max-w-[1064px]"/>
    </div>
    <div class="w-full flex justify-center">
      <div class="flex flex-col space-y-6 w-[1000px]">
        <slot></slot>
      </div>
    </div>
    <TheFooter/>
  </div>
</template>

<script lang="ts" setup>
import TheFooter from '../components/layout/TheFooter.vue'
import TheHeader from '../components/layout/TheHeader.vue'
</script>

<style scoped>

</style>