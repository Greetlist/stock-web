<template>
  <v-chart class='chart' :option='option' ref="pieChart"/>
  <p> {{ this.distributionData }} </p>
</template>

<script>
import { use } from 'echarts/core'
import { CanvasRenderer } from 'echarts/renderers'
import { 
  PieChart,
  RadarChart
} from 'echarts/charts'
import {
  TitleComponent,
  TooltipComponent,
  LegendComponent
} from 'echarts/components'
import VChart from 'vue-echarts'

use([
  CanvasRenderer,
  TitleComponent,
  TooltipComponent,
  LegendComponent,
  PieChart,
  RadarChart,
])

export default {
  name: 'MarketDistributionPie',
  components: {
    VChart
  },
  data: function () {
    return {
      option: {
        title: {
          text: 'text',
          subtext: 'text'
        },
        legend: {
          data: ['Allocated Budget', 'Actual Spending']
        },
        radar: {
          indicator: [
            { name: 'Sales', max: 6500 },
            { name: 'Administration', max: 16000 },
            { name: 'Information Technology', max: 30000 },
            { name: 'Customer Support', max: 38000 },
            { name: 'Development', max: 52000 },
            { name: 'Marketing', max: 25000 }
          ]
        },
        series: [
          {
            name: 'Budget vs spending',
            type: 'radar',
            data: [
              {
                value: [4200, 3000, 20000, 35000, 50000, 18000],
                name: 'Allocated Budget'
              },
              {
                value: [5000, 14000, 28000, 26000, 42000, 21000],
                name: 'Actual Spending'
              }
            ]
          }
        ]
      }
    }
  },
  props: {
    distributionData: Object
  },
  methods: {
    setGraphOption () {
      this.$refs.pieChart.setOption({
        title: {
          text: 'text',
          subtext: 'text'
        },
        legend: {
          data: ['Allocated Budget', 'Actual Spending']
        },
        radar: {
          indicator: [
            { name: 'Sales', max: 6500 },
            { name: 'Administration', max: 16000 },
            { name: 'Information Technology', max: 30000 },
            { name: 'Customer Support', max: 38000 },
            { name: 'Development', max: 52000 },
            { name: 'Marketing', max: 25000 }
          ]
        },
        series: [
          {
            name: 'Budget vs spending',
            type: 'radar',
            data: [
              {
                value: [4200, 3000, 20000, 35000, 50000, 18000],
                name: 'Allocated Budget'
              },
              {
                value: [5000, 14000, 28000, 26000, 42000, 21000],
                name: 'Actual Spending'
              }
            ]
          }
        ]
      })
      console.log(this.distributionData)
    }
  },
  mounted: function() {
    this.setGraphOption()
  }
}
</script>
