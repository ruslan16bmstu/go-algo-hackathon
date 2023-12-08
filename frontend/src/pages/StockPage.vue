<template>
  <PageLayout>
    <div>{{ route.params.stockId }}</div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { useRoute } from 'vue-router'
import PageLayout from '../layouts/PageLayout.vue'

const route = useRoute()

</script>

<style scoped>

</style>