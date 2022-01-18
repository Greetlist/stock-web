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
  watch: {
    '$store.state.selectAlgo': function () {
      this.getRecommandStock()
    }
  },
  methods: {
    getRecommandStock () {
      if (this.$store.state.selectAlgo === '') {
        return
      }
      var param = {
        query_date: "",
        stock_code: this.stockCode,
        query_data_len: this.queryDataLen,
        algo_name: this.$store.state.selectAlgo
      }
      axios.post(this.server+'/api/stock/getRecommandStockPrediction', param, { header })
        .then((response) => {
          this.stockDatas = response.data.stock_datas
          if (response.data.stock_prediction_datas !== undefined) {
            this.predDatas = response.data.stock_prediction_datas
            for (var i = 0; i < response.data.stock_datas.length; i++) {
              this.stockDatas[i].prediction_record = this.predDatas[i].prediction_record
              this.stockDatas[i].show_msg = this.predDatas[i].show_msg
            }
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
