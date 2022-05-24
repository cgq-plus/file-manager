import Vue from 'vue'
import App from './App.vue'
import './plugins/element.js'

import './assets/styles/element-variables.scss'

import '@/assets/styles/index.scss' // global css

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
