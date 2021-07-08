<template>
  <v-chart class='chart' :option='option' ref="stockGraph"/>
</template>

<script>
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { CandlestickChart, LineChart } from 'echarts/charts'
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
  CandlestickChart,
  LineChart
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
          left: 'right'
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          }
        },
        legend: {
        },
        xAxis: [
          {
            type: 'category',
            data: this.dateList,
            scale: true,
            axisLine: { onZero: false },
            splitLine: { show: false },
            splitNumber: 20
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
            start: 80,
            end: 100
          }
        ]
      },
      stockCode: '',
      dateList: [],
      dataList: [],
      volumeList: [],
      moneyList: [],
      ma13List: [],
      ma34List: [],
      ma55List: []
    }
  },
  props: {
    singleStockDataItem: Object
  },
  mounted: function () {
    this.setGraphOption()
  },
  watch: {
    singleStockDataItem: function () {
      this.setGraphOption()
    }
  },
  methods: {
    setGraphOption: function () {
      var vm = this
      vm.stockCode = vm.singleStockDataItem.stock_code
      for (var i = 0; i < vm.singleStockDataItem.records.length; i++) {
        var record = vm.singleStockDataItem.records[i]
        vm.dataList.push([record.open, record.close, record.low, record.high])
        vm.dateList.push(record.date)
        vm.volumeList.push(record.volume)
        vm.moneyList.push(record.money)
        vm.ma13List.push(record.ma13)
        vm.ma34List.push(record.ma34)
        vm.ma55List.push(record.ma55)
      }
      vm.$refs.stockGraph.setOption({
        xAxis: {
          data: vm.dateList
        },
        series: [
          {
            type: 'candlestick',
            data: vm.dataList,
            itemStyle: {
              color: '#DC143C',
              color0: '#32CD32',
              borderColor: '#FF0000',
              borderColor0: '#00FF00'
            }
          },
          {
            name: 'ma13',
            type: 'line',
            data: vm.ma13List,
            smooth: true,
            lineStyle: {
              color: '#000080',
              opacity: 0.7
            }
          },
          {
            name: 'ma34',
            type: 'line',
            data: vm.ma34List,
            smooth: true,
            lineStyle: {
              color: '#00EE00',
              opacity: 0.7
            }
          },
          {
            name: 'ma55',
            type: 'line',
            data: vm.ma55List,
            smooth: true,
            lineStyle: {
              color: '#CD0000',
              opacity: 0.7
            }
          }
        ],
        title: {
          text: vm.stockCode
        }
      })
    }
  }
}
</script>

<style scoped>
.chart {
  height: 800px;
  weight: 70%
}
</style>
