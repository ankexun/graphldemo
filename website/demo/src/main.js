import Vue from 'vue'
import VueApollo from 'vue-apollo'
import {apolloClient,baseClient} from './utils/apollo'
import App from './App.vue'

Vue.config.productionTip = false

const apolloProvider = new VueApollo({
    clients: {
      api: apolloClient,   //需要添加请求头
      base: baseClient
    },
    defaultClient: baseClient  //默认请求路径，如果只有一个请求就使用这个就行
  })
  
// Vue.use(VueApollo)
new Vue({
  apolloProvider,
  render: h => h(App),
}).$mount('#app')