<template>
  <el-date-picker
      v-model="selectDate"
      type="date"
      placeholder="Select Date"
      value-format="YYYY-MM-DD"
      v-on:change="getMarketDistribution">
  </el-date-picker>
  <div>
    <MarketDistributionPie v-for="(distributionData, index) in marketDistributionData" v-bind:distributionData="distributionData" :key="index"></MarketDistributionPie>
  </div>
</template>

<script>
import MarketDistributionPie from '@/components/MarketDistributionPie.vue'
import { getCurrentInstance } from 'vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'MarketDistributionView',
  components: {
    MarketDistributionPie
  },
  data: function () {
    return {
      marketDistributionData: [],
      selectDate: '',
      server: ''
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
  },
  methods: {
    getMarketDistribution () {
      var param = {
        query_date: this.selectDate
      }
      axios.post(this.server+'/api/stock/getMarketDistribution', param, { header })
        .then((response) => {
          var dataList = response.data.distribution_data_list
          for (var i = 0; i < dataList.length; i++) {
            var curItem = [
              {value: dataList[i].large, name: 'Large'},
              {value: dataList[i].mid, name: 'Mid'},
              {value: dataList[i].small, name: 'Small'}
            ]
            this.marketDistributionData.push(curItem)
          }
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>
