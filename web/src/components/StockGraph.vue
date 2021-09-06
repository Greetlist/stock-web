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
          left: 'right',
          top: 'bottom',
          textStyle: {
            width: '1vh',
            height: '1vh' 
          }
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
            start: 0,
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
    this.clearAllData()
    this.setGraphOption()
  },
  watch: {
    singleStockDataItem: function () {
      this.clearAllData()
      this.setGraphOption()
    }
  },
  methods: {
    clearAllData () {
      this.dateList = [],
      this.dataList = [],
      this.volumeList = [],
      this.moneyList = [],
      this.ma13List = [],
      this.ma34List = [],
      this.ma55List = []
    },
    setGraphOption () {
      this.stockCode = this.singleStockDataItem.stock_code
      for (var i = 0; i < this.singleStockDataItem.records.length; i++) {
        var record = this.singleStockDataItem.records[i]
        this.dataList.push([record.open, record.close, record.low, record.high])
        this.dateList.push(record.date)
        this.volumeList.push(record.volume)
        this.moneyList.push(record.money)
        this.ma13List.push(record.ma13)
        this.ma34List.push(record.ma34)
        this.ma55List.push(record.ma55)
      }
      this.$refs.stockGraph.setOption({
        xAxis: {
          data: this.dateList
        },
        series: [
          {
            type: 'candlestick',
            data: this.dataList,
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
            data: this.ma13List,
            smooth: true,
            lineStyle: {
              color: '#000080',
              opacity: 0.7
            }
          },
          {
            name: 'ma34',
            type: 'line',
            data: this.ma34List,
            smooth: true,
            lineStyle: {
              color: '#00EE00',
              opacity: 0.7
            }
          },
          {
            name: 'ma55',
            type: 'line',
            data: this.ma55List,
            smooth: true,
            lineStyle: {
              color: '#CD0000',
              opacity: 0.7
            }
          }
        ],
        title: {
          text: this.stockCode
        }
      })
    }
  }
}
</script>

<style scoped>
.chart {
  height: 40vh;
}
</style>
