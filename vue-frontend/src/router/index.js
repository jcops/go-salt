import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/* Router Modules */
// import componentsRouter from './modules/components'
// import chartsRouter from './modules/charts'
// import tableRouter from './modules/table'
// import nestedRouter from './modules/nested'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noredirect           if `redirect:noredirect` will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission-bak requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  // {
  //   path: '/redirect',
  //   component: Layout,
  //   hidden: true,
  //   children: [
  //     {
  //       path: '/redirect/:path*',
  //       component: () => import('@/views/redirect/index')
  //     }
  //   ]
  // },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  // {
  //   path: '/auth-redirect',
  //   component: () => import('@/views/login/authRedirect'),
  //   hidden: true
  // },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '',
    component: Layout,
    redirect: 'dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: '首页',
        meta: { title: '首页', icon: 'dashboard', noCache: true, affix: true }
      }
    ]
  },
  // {
  //   path: '/documentation',
  //   component: Layout,
  //   children: [
  //     {
  //       path: 'index',
  //       component: () => import('@/views/documentation/index'),
  //       name: 'Documentation',
  //       meta: { title: 'documentation', icon: 'documentation' }
  //     }
  //   ]
  // },
  {
    path: '/usercenter',
    component: Layout,
    // redirect: '/table/complex-table',
    name: '用户管理',
    meta: {
      title: '用户管理',
      icon: 'table'
    },
    children: [
      {
        path: 'user',
        component: () => import('@/views/usercenter/user/index'),
        name: '用户信息',
        meta: { title: '用户信息', icon: 'guide', noCache: true }
      },
      {
        path: 'group',
        component: () => import('@/views/usercenter/group/index'),
        name: '角色信息',
        meta: { title: '角色信息', icon: 'guide', noCache: true }
      }
    ]
  },
  {
    path: '/salt',
    component: Layout,
    // redirect: '/table/complex-table',
    name: 'salt管理',
    meta: {
      title: 'salt管理',
      icon: 'table'
    },
    children: [
      {
        path: 'keyping',
        component: () => import('@/views/salt/keyping/index'),
        name: '节点连通性',
        meta: { title: '节点连通性', icon: 'guide', noCache: true }
      },
      {
        path: 'keylist',
        component: () => import('@/views/salt/keylist/index'),
        name: '节点配置',
        meta: { title: '节点配置', icon: 'guide', noCache: true }
      },
      {
        path: 'keyrun',
        component: () => import('@/views/salt/keyrun/index'),
        name: '执行命令',
        meta: { title: '执行命令', icon: 'guide', noCache: true }
      },
      {
        path: 'keycopy',
        component: () => import('@/views/salt/keycopy/index'),
        name: '文件拷贝',
        meta: { title: '文件拷贝', icon: 'guide', noCache: true }
      },
      {
        path: 'keydeploy',
        component: () => import('@/views/salt/keydeploy/index'),
        name: '应用部署',
        meta: { title: '应用部署', icon: 'guide', noCache: true }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */

export default new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

// const router = createRouter()
//
// // Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
// export function resetRouter() {
//   const newRouter = createRouter()
//   router.matcher = newRouter.matcher // reset router
// }
//
// export default router
