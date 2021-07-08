<template>
  <div>
    <StockSelectInput ref="stockInput" v-on:select="updateStockData"></StockSelectInput>
    <StockGraph v-for="(stockData, index) in stockDatas" v-bind:singleStockDataItem="stockData" :key="index"></StockGraph>
  </div>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
import StockSelectInput from '@/components/StockSelectInput.vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'StockOverView',
  data: function () {
    return {
      stockDatas: ''
    }
  },
  components: {
    StockGraph,
    StockSelectInput
  },
  methods: {
    updateStockData: function () {
      var vm = this
      var stockCode = this.$refs.stockInput.selectedCode
      var param = {
        stock_list: [stockCode]
      }
      axios.post('http://121.5.100.186:8081/api/stock/getQueryStockData', param, { header })
        .then(function (response) {
          vm.stockDatas = response.data.stock_datas
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
