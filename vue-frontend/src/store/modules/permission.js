import { constantRoutes } from '@/router'
import router from '@/router'
/**
 * Use meta.role to determine if the current user has permission-bak
 * @param roles
 * @param route
 */
// function hasPermission(roles, route) {
//   if (route.meta && route.meta.roles) {
//     return roles.some(role => route.meta.roles.includes(role))
//   } else {
//     return true
//   }
// }

/**
 * Filter asynchronous routing tables by recursion
 * @param routes asyncRoutes
 * @param roles
 */
export function filterconstantRoutes(routes) {
  const res = []
  routes.forEach(route => {
    const tmp = { ...route }
    // if (hasPermission(roles, tmp)) {
    if (tmp.children) {
      tmp.children = filterconstantRoutes(tmp.children)
    }
    res.push(tmp)
    // }
  })

  return res
}

const state = {
  routes: [],
  addRoutes: []
}

const mutations = {
  SET_ROUTES: (state, routes) => {
    state.addRoutes = routes
    state.routes = filterconstantRoutes(routes)
  }
}

const actions = {
  generateRouter({ commit }) {
    return new Promise(resolve => {
      commit('SET_ROUTES', filterconstantRoutes(constantRoutes))
      resolve(constantRoutes)
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
