import Vue from 'vue'
import VueApollo from 'vue-apollo'
import {apolloClient,baseClient} from './utils/apollo'
import App from './App'

import store from './store'

Vue.config.productionTip = false

Vue.prototype.$store = store

const apolloProvider = new VueApollo({
    clients: {
      api: apolloClient,   //需要添加请求头
      base: baseClient
    },
    defaultClient: baseClient  //默认请求路径，如果只有一个请求就使用这个就行
  })

App.mpType = 'app'

const app = new Vue({
    store,
    apolloProvider,
	...App
})
app.$mount()
