import { createApp } from 'vue'
import App from './App.vue'
import store from './store'
import router from './router'
import 'element-plus/dist/index.css'
import {
  ElButton,
  ElMessageBox,
  ElContainer,
  ElAside,
  ElMenu,
  ElSubMenu,
  ElMenuItem,
  ElMenuItemGroup,
  ElHeader,
  ElMain,
  ElRow,
  ElCol,
  ElAutocomplete,
  ElDatePicker,
  ElPagination,
  ElInputNumber
} from 'element-plus'

// Icons
import {
  Fold,
  Expand,
  BrushFilled
} from '@element-plus/icons'

const app = createApp(App)

app.use(ElButton)
app.use(ElMessageBox)
app.use(ElContainer)
app.use(ElAside)
app.use(ElMenu)
app.use(ElSubMenu)
app.use(ElMenuItem)
app.use(ElMenuItemGroup)
app.use(ElHeader)
app.use(ElMain)
app.use(ElRow)
app.use(ElCol)
app.use(ElAutocomplete)
app.use(Fold)
app.use(Expand)
app.use(BrushFilled)
app.use(ElDatePicker)
app.use(ElPagination)
app.use(ElInputNumber)

// for global var
app.config.globalProperties.$server = 'http://121.5.100.186:8082'

app.use(store).use(router).mount('#app')
