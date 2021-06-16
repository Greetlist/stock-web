<template>
  <div>
    <StockSelectInput></StockSelectInput>
    <StockGraph v-if="show"></StockGraph>
  </div>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
import StockSelectInput from '@/components/StockSelectInput.vue'
export default {
  name: 'StockOverView',
  data: function () {
    return {
      show: true,
      selectCode: ''
    }
  },
  components: {
    StockGraph,
    StockSelectInput
  }
}
</script>
