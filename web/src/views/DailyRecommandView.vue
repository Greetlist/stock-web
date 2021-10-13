<template>
  <el-date-picker
      v-model="selectDate"
      type="date"
      placeholder="Select Date"
      value-format="YYYY-MM-DD"
      v-on:change="getRecommandStock">
  </el-date-picker>
  <div>
    <StockGraph v-for="(stockData, index) in stockDatasSlice" v-bind:singleStockDataItem="stockData" :key="index"></StockGraph>
  </div>
  <el-row justify="space-around">
    <el-pagination
        layout="prev, pager, next"
        :page-count="totalPageCount"
        :hide-on-single-page="true"
        @current-change="changeStockSlice"
        >
    </el-pagination>
  </el-row>
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
      predDatas: '',
      stockDatasSlice: '',
      selectDate: '',
      totalPageCount: 1,
      singlePageDataNum: 2,
      server: ''
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
    const today = new Date()
    this.selectDate = today.getFullYear() + '-' + (today.getMonth()+1) + '-' + today.getDate()
    this.getRecommandStock()
  },
  methods: {
    getRecommandStock () {
      var param = {
        query_date: this.selectDate
      }
      axios.post(this.server+'/api/stock/getRecommandStockPrediction', param, { header })
        .then((response) => {
          this.stockDatas = response.data.stock_datas
          this.totalPageCount = Math.ceil(response.data.stock_datas.length / this.singlePageDataNum)
          this.stockDatasSlice = this.stockDatas.slice(0, this.singlePageDataNum)
          if (response.data.stock_prediction_datas !== undefined) {
            this.predDatas = response.data.stock_prediction_datas
            for (var i = 0; i < this.singlePageDataNum; i++) {
              this.stockDatasSlice[i].prediction_record = this.predDatas[i].prediction_record
            }
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    changeStockSlice(currentPage) {
      var startIndex = (currentPage-1) * this.singlePageDataNum
      var endIndex = currentPage * this.singlePageDataNum
      this.stockDatasSlice = this.stockDatas.slice(startIndex, endIndex)
      if (this.predDatas != '') {
        for (var i = 0; i < this.singlePageDataNum; i++) {
          this.stockDatasSlice[i].prediction_record = this.predDatas[i+startIndex].prediction_record
        }
      }
    }
  }
}
</script>
