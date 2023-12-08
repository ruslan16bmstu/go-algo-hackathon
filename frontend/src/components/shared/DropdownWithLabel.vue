<template>
  <DropdownMenu :overlay="false" class="hover:cursor-pointer select-none"
                @opened="isExpanded = true" @closed="isExpanded = false">
    <template #trigger>
      <div ref="closeTrigger" class="flex items-center">
        <div class="mr-1 font-bold opacity-50">{{ label }}</div>
        <div class="font-bold">{{ selected }}</div>
        <ExpandIcon class="opacity-50 transition-transform duration-200"
                    :class="{'rotate-180': isExpanded}"/>
      </div>
    </template>
    <template #body>
      <ul class="max-h-[240px] overflow-y-auto text-sm py-1">
        <li v-for="(option, i) in options" :key="i"
            @click="onSelect(i)">
          {{ option }}
        </li>
      </ul>
    </template>
  </DropdownMenu>
</template>

<script lang="ts" setup generic="T">
//@ts-ignore
import DropdownMenu from 'v-dropdown-menu'
import 'v-dropdown-menu/css'
import { computed, ref } from 'vue'
import ExpandIcon from '../../assets/expand.svg'

const props = defineProps<{
  label: string
  options: T[]
}>()


const closeTrigger = ref<null | HTMLDivElement>(null)

const closeMenu = () => {
  if (closeTrigger.value) {
    closeTrigger.value.click()
  }
}

const onSelect = (ind: number) => {
  selectedIndex.value = ind
  closeMenu()
}

const selectedIndex = ref(0)
const selected = computed(() => props.options[selectedIndex.value])
const isExpanded = ref(false)

defineEmits(['update:modelValue'])
</script>

<style scoped>
li {
  @apply px-2 py-1 hover:text-[#5142ab] transition ease-in-out;
}
</style>