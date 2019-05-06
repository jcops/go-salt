import Vue from 'vue'

import Cookies from 'js-cookie'

import 'normalize.css/normalize.css' // a modern alternative to CSS resets

import Element from 'element-ui'
import './styles/element-variables.scss'

import '@/styles/index.scss' // global css

import App from './App'
import store from './store'
import router from './router'

// import i18n from './lang' // internationalization
import './icons' // icon
import './permission' // permission-bak control
import './utils/error-log' // error log

import * as filters from './filters'
// import zh from './lang/zh' // global filters
import elementZhLocale from 'element-ui/lib/locale/lang/zh-CN'// element-ui lang
/**
 * If you don't want to use mock-server
 * you want to use mockjs for request interception
 * you can execute:
 *
 * import { mockXHR } from '../mock'
 * mockXHR()
 */

Vue.use(Element, {
  size: Cookies.get('size') || 'medium', // set element-ui default size
  // i18n: (key, value) => i18n.t(key, value)
  elementZhLocale
})

// register global utility filters
Object.keys(filters).forEach(key => {
  Vue.filter(key, filters[key])
})

Vue.config.productionTip = false

new Vue({
  el: '#app',
  router,
  store,
  // i18n,
  render: h => h(App)
})
