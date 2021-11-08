<template>
  <el-date-picker
      v-model="selectDate"
      type="date"
      placeholder="Select Date"
      value-format="YYYY-MM-DD"
      v-on:change="getRecommandStock">
  </el-date-picker>
  <div>
    <TotalMarketGraph v-bind:totalIndexData="indexData"></TotalMarketGraph>
  </div>
</template>

<script>
import TotalMarketGraph from '@/components/TotalMarketGraph.vue'
import { getCurrentInstance } from 'vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'StockOverView',
  components: {
    TotalMarketGraph
  },
  data: function () {
    return {
      indexData: Object,
      indexRawData: '',
      indexPredData: '',
      selectDate: '',
      server: ''
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
    // const today = new Date()
    // this.selectDate = today.getFullYear() + '-' + (today.getMonth()+1) + '-' + today.getDate()
    // this.getRecommandStock()
  },
  methods: {
    getRecommandStock () {
      var param = {
        query_date: this.selectDate
      }
      axios.post(this.server+'/api/stock/getTotalMarketIndexData', param, { header })
        .then((response) => {
          var curIndexData = new Object()
          curIndexData.index_raw_data = response.data.index_raw_data
          curIndexData.index_pred_data = response.data.index_pred_data
          curIndexData.show_msg = response.data.show_msg
          this.indexData = curIndexData
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
