<template>
  <!--<el-button v-if="isLogout()">Login In</el-button> -->
  <OwnPreferStockTable></OwnPreferStockTable>
</template>
<script>
import OwnPreferStockTable from '@/components/OwnPreferStockTable.vue'
import { getCurrentInstance } from 'vue'
export default {
  name: 'StockHoldView',
  components: {
    OwnPreferStockTable
  },
  methods: {
    isLogout () {
      return getCurrentInstance().appContext.config.globalProperties.$session_id === ''
    }
  }
}
</script>
