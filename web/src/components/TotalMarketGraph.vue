<template>
  <v-chart class='chart' :option='option' ref="indexGraph"/>
  <div v-if="this.totalIndexData.show_msg !== 'undefined'" style="white-space: pre-line;">{{ this.totalIndexData.show_msg }}</div>
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
  name: 'TotalMarketGraph',
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
            width: '80%',
            height: '35%',
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
          }
        ],
        yAxis: [
          {
            type: 'value',
            gridIndex: 0,
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
      dateList: [],
      dataList: [],
      brushStartDate: '',
      brushEndDate: ''
    }
  },
  props: {
    totalIndexData: Object
  },
  mounted: function () {
    this.clearAllData()
    this.setAllData()
    this.setGraphOption()
    this.dispatchGraphAction()
  },
  watch: {
    totalIndexData: function () {
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
      this.brushStartDate = '',
      this.brushEndDate = ''
    },
    setAllData () {
      this.dateList = this.totalIndexData.index_raw_data.date
      for (var i = 0; i < this.totalIndexData.index_raw_data.date.length; i++) {
        var open = this.totalIndexData.index_raw_data.open[i]
        var close = this.totalIndexData.index_raw_data.close[i]
        var low = this.totalIndexData.index_raw_data.low[i]
        var high = this.totalIndexData.index_raw_data.high[i]
        this.dataList.push([open, close, low, high])
      }

      // append prediction data if exists
      if (this.totalIndexData.index_pred_data !== undefined) {
        this.dateList = this.dateList.concat(this.totalIndexData.index_pred_data.date)
        var predictionLen = this.totalIndexData.index_pred_data.date.length
        for (var j = 0; j < predictionLen; j++) {
          var predOpen = this.totalIndexData.index_pred_data.open[j]
          var predClose = this.totalIndexData.index_pred_data.close[j]
          var predLow = this.totalIndexData.index_pred_data.low[j]
          var predHigh = this.totalIndexData.index_pred_data.high[j]
          this.dataList.push([predOpen, predClose, predLow, predHigh])
        }
      }

      var dateListLen = this.dataList.length
      this.brushEndDate = this.dateList[dateListLen-1]
      this.brushStartDate = this.dateList[dateListLen-predictionLen]
    },
    setGraphOption () {
      this.$refs.indexGraph.setOption({
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
        ],
        title: {
          text: this.totalIndexData.index_name
        }
      })
    },
    dispatchGraphAction () {
      this.$refs.indexGraph.dispatchAction({
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
