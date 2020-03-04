/**
 * Created by kevin on 2019-10-11.
 * http配置
 */

import store from '@/store/index'
import * as types from '@/store/mutation-types'
import { getCurrentPageUrl } from '@/utils/index'

var Fly=require("flyio/dist/npm/wx")
var fly=new Fly

// fly配置
fly.config.timeout=10000
// 这个url一定要找到nginx的相应location ^~ 部分
//#ifdef MP-WEIXIN
fly.config.baseURL = 'https://m.zxoa.com.cn/agent/web/'
//#endif
//#ifdef H5
fly.config.baseURL = '/agent/web'
//#endif

// http request 拦截器
fly.interceptors.request.use(
    (request) => {
      if (store.state.token) {
        request.headers.Authorization = store.state.token
        // request.headers.Authorization = `isLogin`
        // 以下语句只在微信小程序时模拟Cookie编译
        //#ifdef MP-WEIXIN
        request.headers.Cookie = uni.getStorageSync("sessionid")
      } //#endif
      return request
    },
    err => {
      return Promise.reject(err)
    },
  )
  
// http response 拦截器
fly.interceptors.response.use(
    response => {
        return response
    },
    error => {
        if (error.response) {
            switch (error.response.status) {
                case 401:
                // 401 清除token信息并跳转到登录页面
                store.commit(types.LOGOUT)
                
                // 只有在当前路由不是登录页面才跳转
                getCurrentPageUrl().pop() != '/pages/user/login' &&
                uni.redirectTo({
                  url: 'pages/user/login'
                })
            }
        }
        // console.log(JSON.stringify(error));//console : Error: Request failed with status code 402
        return Promise.reject(error.response.data)
    },
)

export default fly