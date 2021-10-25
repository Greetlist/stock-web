<template>
  <v-chart class='chart' :option='option' ref="stockGraph"/>
</template>

<script>
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { 
  CandlestickChart,
  LineChart,
  BarChart 
} from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  GridComponent,
  DataZoomComponent,
  BrushComponent
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
  LineChart,
  BarChart,
  BrushComponent
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
            width: '1px',
            height: '1px'
          }
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross'
          },
          borderWidth: 1,
          borderColor: '#ccc',
          padding: 10,
          textStyle: {
            color: '#000'
          },
          position: [10, 10]
        },
        axisPointer: {
          link: { xAxisIndex: 'all' },
          label: {
            backgroundColor: '#777'
          }
        },
        legend: {
          bottom: '1%',
          left: 'center'
        },
        brush: {
          xAxisIndex: 'all',
          brushLink: 'all',
          outOfBrush: {
            colorAlpha: 0.1
          }
        },
        grid: [
          {
            left: 'center',
            width: '90%',
            height: '35%',
            top: '5%'
          },
          {
            left: 'center',
            width: '90%',
            height: '10%',
            bottom: '25%'
          }
        ],
        xAxis: [
          {
            type: 'category',
            gridIndex: 0,
            data: this.dateList,
            scale: true,
            boundaryGap: false,
            axisLine: { onZero: false },
            splitLine: { show: false },
            splitNumber: 20
          },
          {
            type: 'category',
            gridIndex: 1,
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
            gridIndex: 0,
            scale: true
          },
          {
            scale: true,
            gridIndex: 1,
            splitNumber: 2,
            axisLabel: {show: false},
            axisLine: {show: false},
            axisTick: {show: false},
            splitLine: {show: false}
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
      stockName: '',
      dateList: [],
      dataList: [],
      volumeList: [],
      moneyList: [],
      ma13List: [],
      ma34List: [],
      ma55List: [],
      brushStartDate: '',
      brushEndDate: ''
    }
  },
  props: {
    singleStockDataItem: Object
  },
  mounted: function () {
    this.clearAllData()
    this.setAllData()
    this.setGraphOption()
    this.dispatchGraphAction()
  },
  watch: {
    singleStockDataItem: function () {
      this.clearAllData()
      this.setAllData()
      this.setGraphOption()
      this.dispatchGraphAction()
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
      this.ma55List = [],
      this.brushStartDate = '',
      this.brushEndDate = ''
    },
    setAllData () {
      // 0 --> day 1 --> week 2 --> month
      var select_period = 0
      this.stockCode = this.singleStockDataItem.stock_info.stock_code
      this.stockName = this.singleStockDataItem.stock_info.stock_name
      this.dateList = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.date
      this.volumeList = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.volume
      this.moneyList = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.money
      this.ma13List = this.singleStockDataItem.stock_raw_datas[select_period].factor_raw_data.ma13
      this.ma34List = this.singleStockDataItem.stock_raw_datas[select_period].factor_raw_data.ma34
      this.ma55List = this.singleStockDataItem.stock_raw_datas[select_period].factor_raw_data.ma55
      for (var i = 0; i < this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.date.length; i++) {
        var open = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.open[i]
        var close = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.close[i]
        var low = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.low[i]
        var high = this.singleStockDataItem.stock_raw_datas[select_period].market_raw_data.high[i]
        this.dataList.push([open, close, low, high])
      }

      // append prediction data if exists
      if (this.singleStockDataItem.prediction_record !== undefined) {
        this.dateList = this.dateList.concat(this.singleStockDataItem.prediction_record.date)
        var predictionLen = this.singleStockDataItem.prediction_record.date.length
        for (var j = 0; j < predictionLen; j++) {
          var predOpen = this.singleStockDataItem.prediction_record.open[j]
          var predClose = this.singleStockDataItem.prediction_record.close[j]
          var predLow = this.singleStockDataItem.prediction_record.low[j]
          var predHigh = this.singleStockDataItem.prediction_record.high[j]
          this.dataList.push([predOpen, predClose, predLow, predHigh])
        }
      }

      var dateListLen = this.dataList.length
      this.brushEndDate = this.dateList[dateListLen-1]
      this.brushStartDate = this.dateList[dateListLen-predictionLen]
    },
    setGraphOption () {
      this.$refs.stockGraph.setOption({
        xAxis: [
          {
            type: 'category',
            data: this.dateList
          },
          {
            type: 'category',
            data: this.dateList
          }
        ],
        series: [
          {
            name: 'FinanceData',
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
          },
          {
            name: 'Volume',
            type: 'bar',
            data: this.volumeList,
            xAxisIndex: 1,
            yAxisIndex: 1
          }
        ],
        title: {
          text: this.stockCode,
          subtext: this.stockName
        }
      })
    },
    dispatchGraphAction () {
      this.$refs.stockGraph.dispatchAction({
        type: 'brush',
        areas: [
          {
            brushType: 'lineX',
            coordRange: [this.brushStartDate, this.brushEndDate],
            xAxisIndex: 0
          }
        ]
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
