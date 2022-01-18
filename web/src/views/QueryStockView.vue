<template>
  <el-row>
    <el-col :span="12">
      <StockSelectInput ref="stockInput" v-on:select="updateStockData"></StockSelectInput>
    </el-col>
    <el-col :span="12">
      <el-input-number v-model="queryDataLen" :min="10" :max="60" label="Query N Days"></el-input-number>
    </el-col>
  </el-row>
  <el-row>
    <el-col :span="24">
      <StockGraph v-for="(stockData, index) in stockDatas" v-bind:singleStockDataItem="stockData" :key="index"></StockGraph>
    </el-col>
  </el-row>
</template>

<script>
import StockGraph from '@/components/StockGraph.vue'
import StockSelectInput from '@/components/StockSelectInput.vue'
import { getCurrentInstance } from 'vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'StockOverView',
  data: function () {
    return {
      stockDatas: '',
      queryDataLen: 30,
      server: ''
    }
  },
  components: {
    StockGraph,
    StockSelectInput
  },
  watch: {
    '$store.state.selectAlgo': function () {
      this.updateStockData()
    }
  },
  methods: {
    updateStockData () {
      var stockCode = this.$refs.stockInput.selectedCode.split("(")[0]
      if (this.$store.state.selectAlgo === '' || stockCode === '') {
        console.log("Param is Invalid")
        return
      }
      var param = {
        stock_code: stockCode,
        query_data_len: this.queryDataLen,
        algo_name: this.$store.state.selectAlgo
      }
      axios.post(this.server+'/api/stock/queryStockData', param, { header })
        .then((response) => {
          this.stockDatas = response.data.stock_datas
          for (var i = 0; i < this.stockDatas.length; i++) {
            this.stockDatas[i].prediction_record = response.data.stock_prediction_datas[i].prediction_record
            this.stockDatas[i].show_msg = response.data.stock_prediction_datas[i].show_msg
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
  }
}
</script>
