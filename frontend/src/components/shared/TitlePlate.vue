<template>
  <div class="plate w-full pad text-xl font-bold">
    {{ title }}
  </div>
</template>

<script lang="ts" setup>
defineProps<{
  title: string
}>()
</script>

<style scoped>

</style>