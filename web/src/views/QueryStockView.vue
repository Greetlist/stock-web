<template>
  <div>
    <StockSelectInput ref="stockInput" v-on:select="updateStockData"></StockSelectInput>
    <StockGraph ref="stockGraph"></StockGraph>
  </div>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
import StockSelectInput from '@/components/StockSelectInput.vue'
export default {
  name: 'StockOverView',
  data: function () {
    return {
      selectCode: ''
    }
  },
  components: {
    StockGraph,
    StockSelectInput
  },
  methods: {
    updateStockData: function () {
      var stockCode = this.$refs.stockInput.selectedCode
      this.$refs.stockGraph.getQueryStockData(stockCode)
    }
  }
}
</script>
