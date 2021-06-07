<template>
  <div>
    <HeadNavBar></HeadNavBar>
    <StockGraph></StockGraph>
  </div>
</template>

<script>
import HeadNavBar from '@/components/NavBar.vue'
import StockGraph from '@/components/StockGraph.vue'
export default {
  name: 'HomeView',
  components: {
    HeadNavBar,
    StockGraph
  }
}
</script>
