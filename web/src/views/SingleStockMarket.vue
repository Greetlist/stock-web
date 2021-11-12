<template>
  <StockGraph v-for="(stockData, index) in stockDatas" v-bind:singleStockDataItem="stockData" :key="index"></StockGraph>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
import { getCurrentInstance } from 'vue'
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
      stockDatas: '',
      stockCode: '',
      queryDataLen: 90,
      server: ''
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
    this.stockCode = this.$route.params.stockCode
    console.log(this.stockCode)
    this.getRecommandStock()
  },
  methods: {
    getRecommandStock () {
      var param = {
        stock_list: [this.stockCode],
        query_data_len: this.queryDataLen
      }
      axios.post(this.server+'/api/stock/queryStockData', param, { header })
        .then((response) => {
          this.stockDatas = response.data.stock_datas
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
