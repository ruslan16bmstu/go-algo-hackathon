import { ref } from 'vue'

export const useLoading = () => {
  const isLoading = ref(false)
  
  const load = async <T>(func: () => T) => {
    let res = null
    
    if (isLoading.value) {
      return
    }
    
    isLoading.value = true
    
    try {
      res = await func()
    } catch (e) {
      console.log('Error while loading:', e)
    } finally {
      isLoading.value = false
    }
    
    return res
  }
  
  return {isLoading, load}
}