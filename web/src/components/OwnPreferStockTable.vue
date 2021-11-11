<template>
  <el-table
    :data="tableDatas"
    style="width: 100%"
    stripe
    @row-click="showDialog"
  >
    <el-table-column prop="stock_code" label="StockCode" :width="colWidth" />
    <el-table-column prop="exchange" label="Exchange" :width="colWidth" />
    <el-table-column prop="name" label="Name" :width="colWidth" />
    <el-table-column prop="industry" label="Industry" :width="colWidth" />
  </el-table>
  <el-dialog
    title="Select Operation"
    v-model="selectVisible"
    :width="dialogWidth"
  >
    <el-button type="primary" v-on:click="showMarket">Market</el-button>
    <el-button type="danger" v-on:click="showDeleteDialog">Delete</el-button>
  </el-dialog>
  <el-dialog
    title="Delete Confirm"
    v-model="deleteVisible"
    :width="dialogWidth"
  >
    <el-button type="danger" v-on:click="deletePreferStock">Do Delete</el-button>
    <el-button v-on:click="closeDeleteDialog">Cancel</el-button>
  </el-dialog>
</template>

<script>
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}
import { getCurrentInstance, h } from 'vue'
import { ElNotification } from 'element-plus'
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
      colWidth: 400,
      selectVisible: false,
      deleteVisible: false,
      dialogWidth: 1000,
      selectRowStockCode: ''
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
      this.colWidth = screenWidth / 5
      this.dialogWidth = screenWidth / 2 + 100
    },
    showDialog (row, column, event) {
      console.log(row, column, event)
      this.selectVisible = true
      this.selectRowStockCode = row['stock_code']
    },
    showMarket () {
      console.log(this.selectRowStockCode)
    },
    showDeleteDialog () {
      this.deleteVisible = true
    },
    closeDeleteDialog () {
      this.deleteVisible = false
    },
    deletePreferStock() {
      var param = {
        account: 'greetlist',
        stock_code: this.selectRowStockCode
      }
      console.log(param)
      axios.post(this.server+'/api/user/deletePreferStock', param, { header })
        .then((response) => {
          var success = response.data.is_success
          if (success) {
            this.notiSuccess(response.data.error_msg)
          } else {
            this.notiFailed(response.data.error_msg)
          }
          this.getOwnPreferStock()
          this.closeAllDialog()
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    notiSuccess () {
      ElNotification({
        title: 'Result',
        message: h('i', {style: 'color: green'}, 'Add Prefer Stock Success'),
        type: 'success'
      })
    },
    notiFailed (errMsg) {
      ElNotification({
        title: 'Result',
        message: h('i', {style: 'color: red'}, errMsg),
        type: 'error'
      })
    },
    closeAllDialog() {
      this.deleteVisible = false
      this.selectVisible = false
    }
  }
}
</script>
