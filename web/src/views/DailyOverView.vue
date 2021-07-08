<template>
  <div>
    <StockGraph v-for="(stockData, index) in stockDatas" v-bind:singleStockDataItem="stockData" :key="index"></StockGraph>
  </div>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'StockOverView',
  components: {
    StockGraph
  },
  data: function () {
    return {
      stockDatas: ''
    }
  },
  created: function () {
    var vm = this
    axios.get('http://121.5.100.186:8081/api/stock/getDailyCalcStockData', { header })
      .then(function (response) {
        vm.stockDatas = response.data.stock_datas
      })
      .catch(function (error) {
        console.log(error)
      })
  }
}
</script>
