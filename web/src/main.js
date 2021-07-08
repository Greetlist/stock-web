import Vue from 'vue'
import App from './App.vue'
import router from './router'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import {
  Dialog,
  Button,
  MessageBox,
  Container,
  Aside,
  Menu,
  Submenu,
  MenuItem,
  MenuItemGroup,
  Header,
  Main,
  Row,
  Col,
  Autocomplete,
  Scrollbar
} from 'element-ui'

Vue.use(Dialog)
Vue.use(Button)
Vue.use(MessageBox)
Vue.use(Container)
Vue.use(Aside)
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItem)
Vue.use(MenuItemGroup)
Vue.use(Header)
Vue.use(Main)
Vue.use(Row)
Vue.use(Col)
Vue.use(Autocomplete)
Vue.use(Scrollbar)

Vue.config.productionTip = false

Vue.prototype.$confirm = MessageBox.confirm

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
