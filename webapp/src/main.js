// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.

// 引入一些全局样式
import './assets/css/common.css'

import Vue from 'vue'
import App from './App'
import router from './router'
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import VueParticles from 'vue-particles'

//导入全局样式表
import './assets/css/common.css'

import axios from 'axios'
axios.defaults.baseURL = 'http://127.0.0.1:8080/login'
Vue.prototype.$http = axios

Vue.config.productionTip = false
Vue.use(ElementUI);
Vue.use(VueParticles);

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
