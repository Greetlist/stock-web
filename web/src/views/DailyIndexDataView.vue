<template>
  <el-date-picker
      v-model="selectDate"
      type="date"
      placeholder="Select Date"
      value-format="YYYY-MM-DD"
      v-on:change="getRecommandIndex">
  </el-date-picker>
  <div>
    <TotalMarketGraph v-for="(indexData, index) in indexDataSlice" v-bind:totalIndexData="indexData" :key="index"></TotalMarketGraph>
  </div>
  <el-row justify="space-around">
    <el-pagination
        layout="prev, pager, next"
        :page-count="totalPageCount"
        :hide-on-single-page="true"
        @current-change="changeSlice"
        >
    </el-pagination>
  </el-row>
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
      indexDataSlice: [],
      indexDataList: '',
      selectDate: '',
      singlePageDataNum: 3,
      totalPageCount: 1,
      server: ''
    }
  },
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
    // const today = new Date()
    // this.selectDate = today.getFullYear() + '-' + (today.getMonth()+1) + '-' + today.getDate()
    // this.getRecommandIndex()
  },
  watch: {
    '$store.state.selectAlgo': function () {
      this.getRecommandIndex()
    }
  },
  methods: {
    getRecommandIndex () {
      if (this.selectDate === '' || this.$store.state.selectAlgo === '') {
        console.log("Param is Invalid")
        return
      }
      var param = {
        query_date: this.selectDate,
        algo_name: this.$store.state.selectAlgo
      }
      axios.post(this.server+'/api/stock/getTotalMarketIndexData', param, { header })
        .then((response) => {
          this.indexDataList = response.data.index_data_list
          this.totalPageCount = Math.ceil(response.data.index_data_list.length / this.singlePageDataNum)
          this.indexDataSlice = this.indexDataList.slice(0, this.singlePageDataNum)
          console.log(this.indexDataSlice)
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    changeSlice(currentPage) {
      var startIndex = (currentPage-1) * this.singlePageDataNum
      var endIndex = currentPage * this.singlePageDataNum
      this.indexDataSlice = this.indexDataList.slice(startIndex, endIndex)
    }
  }
}
</script>
