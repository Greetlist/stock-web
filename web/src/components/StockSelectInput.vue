<template>
  <el-autocomplete
    v-model="selectedCode"
    :fetch-suggestions="getSuggestionStockList"
    placeholder="Input StockCode"
    clearable>
  </el-autocomplete>
</template>

<script>
import { getCurrentInstance } from 'vue'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
export default {
  name: 'StockSelectInput',
  data: function () {
    return {
      stockList: [],
      selectedCode: ''
    }
  },
  created () {
    const server = getCurrentInstance().appContext.config.globalProperties.$server
    axios.get(server+'/api/stock/getAllStockCode', header)
      .then((response) => {
        // 可选item只能按照文档上面构造,要有个value，憨逼玩意
        // stockList = ['000001', ...] 这样就显示不了
        var len = response.data.stock_info_list.length
        for (var i = 0; i < len; i++) {
          var itemValue = response.data.stock_info_list[i].stock_code + '(' + response.data.stock_info_list[i].stock_name + ')'
          var newItem = { value: itemValue }
          this.stockList.push(newItem)
        }
      }).catch(function (error) {
        console.log(error)
      })
  },
  methods: {
    getSuggestionStockList(queryString, cb) {
      var curStockList = this.stockList
      var filterList = queryString ? curStockList.filter(this.filterFunc(queryString)) : curStockList
      cb(filterList)
    },
    filterFunc(queryString) {
      return (codeItem) => {
        return (codeItem.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
      }
    },
  }
}
</script>
