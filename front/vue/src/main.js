// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import { BootstrapVue, BIcon, BIconBookmarkStarFill, BIconPersonCircle, BIconAward } from 'bootstrap-vue'
import axios from "axios";
import VueAxios from 'vue-axios'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import '../static/common.css'

Vue.use(BootstrapVue, VueAxios, axios)
Vue.config.productionTip = false
Vue.component('BIcon', BIcon)
Vue.component('BIconBookmarkStarFill', BIconBookmarkStarFill)
Vue.component('BIconPersonCircle', BIconPersonCircle)
Vue.component('BIconAward', BIconAward)


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})

