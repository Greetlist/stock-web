<template>
  <el-autocomplete
    v-model="selectedCode"
    :fetch-suggestions="getSuggestionStockList"
    placeholder="Input StockCode"
    v-on="selfListener">
  </el-autocomplete>
</template>

<script>
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
  created: function () {
    this.getAllStockCode()
  },
  methods: {
    getAllStockCode: function () {
      axios.get('http://121.5.100.186:8081/api/stock/getAllStockCode', header)
        .then(function (response) {
          // 可选item只能按照文档上面构造,要有个value，憨逼玩意
          // stockList = ['000001', ...] 这样就显示不了
          var len = response.data.stock_list.length
          for (var i = 0; i < len; i++) {
            var newItem = { value: response.data.stock_list[i] }
            this.stockList.push(newItem)
          }
        }.bind(this))
        .catch(function (error) {
          console.log(error)
        })
    },
    getSuggestionStockList: function (queryString, cb) {
      var curStockList = this.stockList
      var filterList = queryString ? curStockList.filter(this.filterFunc(queryString)) : curStockList
      cb(filterList)
    },
    filterFunc: function (queryString) {
      return (codeItem) => {
        return (codeItem.value.toLowerCase().indexOf(queryString.toLowerCase()) === 0)
      }
    }
  },
  computed: {
    selfListener: function () {
      return Object.assign({},
        this.$listeners
      )
    }
  }
}
</script>
