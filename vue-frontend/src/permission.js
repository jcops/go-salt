import router from './router'
import store from './store'
import { Message } from 'element-ui'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import { getToken } from '@/utils/auth' // get token from cookie
// import getPageTitle from '@/utils/get-page-title'

NProgress.configure({ showSpinner: false }) // NProgress Configuration

const whiteList = ['/login', '/auth-redirect'] // no redirect whitelist
router.beforeEach((to, from, next) => {
  NProgress.start()
  if (getToken()) {
    if (to.path === '/login') {
      next({ path: '/' })
      NProgress.done() // if current page is dashboard will not trigger	afterEach hook, so manually handle it
    } else {
      if (store.getters.name === '') {
        store.dispatch('GetInfo').then(() => { // 拉取用户信息
          // generate accessible routes map based on roles
          store.dispatch('permission/generateRouter')
          // dynamically add accessible routes
          // router.addRoutes(accessRoutes)
          next({ ...to, replace: true })
          // next()
        }).catch((err) => {
          console.log('catch', err)
          store.dispatch('FedLogOut').then(() => {
            Message.error(err || 'Verification failed, please login again')
            next({ path: '/' })
          })
        })
      } else {
        // const accessRoutes = store.dispatch('permission-bak/generateRoutes')
        // router.addRoutes(accessRoutes)
        // next({ ...to, replace: true })
        next()
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      next()
    } else {
      next('/login')
      NProgress.done()
    }
  }
})
// router.beforeEach(async(to, from, next) => {
//   // start progress bar
//   NProgress.start()
//
//   // set page title
//   // document.title = getPageTitle(to.meta.title)
//
//   // determine whether the user has logged in
//   const hasToken = getToken()
//
//   if (hasToken) {
//     if (to.path === '/login') {
//       // if is logged in, redirect to the home page
//       next({ path: '/' })
//       NProgress.done()
//     } else {
//       // const { roles } = await store.dispatch('user/GetInfo')
//
// //       // generate accessible routes map based on roles
//       const accessRoutes = await store.dispatch('permission-bak/generateRoutes')
//
//       // dynamically add accessible routes
//       router.addRoutes(accessRoutes)
// //
// //       // hack method to ensure that addRoutes is complete
// //       // set the replace: true, so the navigation will not leave a history record
// //       next({ ...to, replace: true })
//     }
//   } else {
//     /* has no token*/
//
//     if (whiteList.indexOf(to.path) !== -1) {
//       // in the free login whitelist, go directly
//       next()
//     } else {
//       // other pages that do not have permission-bak to access are redirected to the login page.
//       // next(`/login?redirect=${to.path}`)
//       next('/login')
//       NProgress.done()
//     }
//   }
// })

router.afterEach(() => {
  // finish progress bar
  NProgress.done()
})
