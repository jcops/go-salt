import axios from 'axios'
import { Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import router from '../router'
// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
  withCredentials: true, // send cookies when cross-domain requests
  timeout: 50000 // request timeout
})

// request拦截器
service.interceptors.request.use(config => {
  if (store.getters.token) {
    config.headers['Authorization'] = 'Token ' + getToken() // 让每个请求携带自定义token 请根据实际情况自行修改
  }
  return config
}, error => {
  // Do something with request error
  console.log(error) // for debug
  Promise.reject(error)
})

// respone拦截器
service.interceptors.response.use(
  response => {
    /**
     * code为非xxx是抛错 可结合自己业务进行修改
     */
    return response.data
  },
  error => {
    if (error.response.status === 400) {
      store.dispatch('FedLogOut').then(() => {
        router.push({ path: '/#/login' })
        Message({
          message: '登陆超时, 请重新登陆',
          type: 'error',
          duration: 5 * 1000
        })
      })
    } else if (error.response.status === 403) {
      // router.push({ path: '/' })
      Message({
        message: '权限拒绝',
        type: 'error'
      })
    } else if (error.response.status === 500) {
      Message({
        message: '服务器错误',
        type: 'error'

      })
      console.log('err' + error)// for debug

      return Promise.reject(error)
    }
    // console.log('err' + error) // for debug
    return Promise.reject(error)
  }
)

export default service

// // request interceptor
// service.interceptors.request.use(
//   config => {
//     // do something before request is sent
//
//     if (store.getters.token) {
//       // let each request carry token --['X-Token'] as a custom key.
//       // please modify it according to the actual situation.
//       config.headers['X-Token'] = getToken()
//     }
//     return config
//   },
//   error => {
//     // do something with request error
//     console.log(error) // for debug
//     return Promise.reject(error)
//   }
// )
//
// // response interceptor
// service.interceptors.response.use(
//   /**
//    * If you want to get information such as headers or status
//    * Please return  response => response
//   */
//
//   /**
//    * Determine the request status by custom code
//    * Here is just an example
//    * You can also judge the status by HTTP Status Code.
//    */
//   response => {
//     const res = response.data
//
//     // if the custom code is not 20000, it is judged as an error.
//     if (res.code !== 20000) {
//       Message({
//         message: res.message || 'error',
//         type: 'error',
//         duration: 5 * 1000
//       })
//
//       // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
//       if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
//         // to re-login
//         MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
//           confirmButtonText: 'Re-Login',
//           cancelButtonText: 'Cancel',
//           type: 'warning'
//         }).then(() => {
//           store.dispatch('user/resetToken').then(() => {
//             location.reload()
//           })
//         })
//       }
//       return Promise.reject(res.message || 'error')
//     } else {
//       return res
//     }
//   },
//   error => {
//     console.log('err' + error) // for debug
//     Message({
//       message: error.message,
//       type: 'error',
//       duration: 5 * 1000
//     })
//     return Promise.reject(error)
//   }
// )
//
// export default service
