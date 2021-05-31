import Vue from 'vue'
// import App from './App.vue'

import VueRouter from 'vue-router'

Vue.use(VueRouter)

const Foo = { template: '<div>foo</div>' }
const Bar = { template: '<div>bar</div>' }

const router = new VueRouter({
  mode: 'hash',
  routes: [
    { path: '/foo', component: Foo },
    { path: '/bar', component: Bar }
  ]
})

Vue.config.productionTip = false

new Vue({
  router
  // render: h => h(App)
}).$mount('#app')
