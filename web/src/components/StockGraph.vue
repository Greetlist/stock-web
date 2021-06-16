<template>
  <v-chart class='chart' :option='option' />
</template>

<script>
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { CandlestickChart } from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent
} from 'echarts/components'
import VChart from 'vue-echarts'
const axios = require('axios').default
const header = {
  'Access-Control-Allow-Origin': '*'
}

use([
  CanvasRenderer,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
  CandlestickChart
])
export default {
  name: 'StockGraph',
  components: {
    VChart
  },
  data: function () {
    return {
      option: {
        title: {
          text: 'Code',
          left: 'center'
        },
        tooltip: {
        },
        legend: {
        },
        xAxis: [
          {
            type: 'category',
            data: this.dateList,
            scale: true,
            axisLine: { onZero: false },
            splitLine: { show: true }
          }
        ],
        yAxis: [
          {
            type: 'value',
            scale: true
          }
        ],
        dataZoom: [
          {
            type: 'inside',
            xAxisIndex: [0, 1],
            start: 98,
            end: 100
          }
        ],
        series: [{
          type: 'candlestick',
          data: this.dataList,
          itemStyle: {
            color: '#c23531',
            color0: '#314656',
            borderColor: '#c23531',
            borderColor0: '#314656'
          }
        }]
      },
      stockCode: '',
      dateList: [],
      dataList: [],
      volumeList: [],
      moneyList: []
    }
  },
  created: function () {
  },
  methods: {
    getQueryStockData: function () {
      var param = {
        stock_list: [this.stockCode]
      }
      axios.post('http://121.5.100.186:8081/api/stock/getQueryStockData', param, { header })
        .then(function (response) {
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    }
  }
}
</script>

<style scoped>
.chart {
  height: 400px;
  weight: 70%
}
</style>
