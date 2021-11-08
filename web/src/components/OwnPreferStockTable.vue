<template>
  <el-table
    :data="tableDatas"
    style="width: 100%"
    stripe
  >
    <el-table-column prop="stock_code" label="StockCode" :width="colWidth" />
    <el-table-column prop="exchange" label="Exchange" :width="colWidth" />
    <el-table-column prop="name" label="Name" :width="colWidth" />
    <el-table-column prop="industry" label="Industry" :width="colWidth" />
  </el-table>
</template>

<script>
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
import { getCurrentInstance } from 'vue'
export default {
  name: 'OwnPreferStockTable',
  created () {
    this.server = getCurrentInstance().appContext.config.globalProperties.$server
    this.getOwnPreferStock()
    this.calcColWidth()
  },
  data: function () {
    return {
      tableDatas: [],
      server: '',
      colWidth: 250
    }
  },
  methods: {
    getOwnPreferStock () {
      var param = {
        account: 'greetlist',
      }
      axios.post(this.server+'/api/user/getPreferStock', param, { header })
        .then((response) => {
          this.tableDatas = response.data.follow_stock_list
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    calcColWidth () {
      const screenWidth = document.body.clientWidth
      if (screenWidth < 1000) {
        this.colWidth = screenWidth / 4
      }
    }
  }
}
</script>
