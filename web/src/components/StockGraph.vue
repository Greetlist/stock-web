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
            data: this.mockDate(),
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
          data: this.mockData()
        }]
      },
      dailyStockData,
      queryStockData
    }
  },
  created: function () {
    const header = {
      'Access-Control-Allow-Origin': '*'
    }
    axios.get('http://121.5.100.186:8081/test', header)
      .then(function (response) {
        console.log(response)
      })
      .catch(function (error) {
        console.log(error)
      })
      .then(function () {
        console.log('Finish Post Request')
      })
  },
  methods: {
    mockDate: function () {
      var dateList = []
      var year = '2020'
      for (var i = 1; i <= 12; i++) {
        for (var j = 1; j < 31; j++) {
          var curStr = year + '-' + i.toString() + '-' + j.toString()
          dateList.push(curStr)
        }
      }
      console.log(dateList)
      return dateList
    },
    mockData: function () {
      var dataList = []
      for (var i = 1; i <= 12; i++) {
        for (var j = 1; j < 31; j++) {
          var curData = [10, 15, 10, 15]
          dataList.push(curData)
        }
      }
      return dataList
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
