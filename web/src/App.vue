<template>
  <div id="app">
    <HomeView></HomeView>
  </div>
</template>

<script>
import HomeView from '@/views/Home.vue'
export default {
  name: 'App',
  components: {
    HomeView
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  margin: 0px;
  padding: 0px;
}
</style>
