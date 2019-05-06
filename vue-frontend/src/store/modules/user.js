// import { userlogin, logout, getInfo } from '@/api/user'
// import { getToken, setToken, removeToken } from '@/utils/auth'
// import router, { resetRouter } from '@/router'
//
// const state = {
//   token: getToken(),
//   name: '',
//   avatar: '',
//   introduction: '',
//   userid: '',
//   roles: []
// }
//
// const mutations = {
//   SET_TOKEN: (state, token) => {
//     state.token = token
//   },
//   SET_INTRODUCTION: (state, introduction) => {
//     state.introduction = introduction
//   },
//   SET_NAME: (state, name) => {
//     state.name = name
//   },
//   SET_AVATAR: (state, avatar) => {
//     state.avatar = avatar
//   },
//   SET_ROLES: (state, roles) => {
//     state.roles = roles
//   },
//   SET_USERID: (state, id) => {
//     state.userid = id
//   }
// }
//
// const actions = {
//   // user login
//   // Login({ commit }, userInfo) {
//   //   const { username, password } = userInfo
//   //   return new Promise((resolve, reject) => {
//   //     login({ username: username.trim(), password: password }).then(response => {
//   //       const { data } = response
//   //       commit('SET_TOKEN', data.token)
//   //       setToken(data.token)
//   //       resolve()
//   //     }).catch(error => {
//   //       reject(error)
//   //     })
//   //   })
//   // },
//   ULogin({ commit }, userInfo) {
//     const username = userInfo.username.trim()
//     console.log(userInfo)
//     return new Promise((resolve, reject) => {
//       userlogin(username, userInfo.password).then(response => {
//         const { data } = response
//         commit('SET_TOKEN', data.token)
//         setToken(data.token)
//       }).catch(error => {
//         reject(error)
//       })
//     })
//   },
//   // 获取用户信息
//   GetInfo({ commit, state }) {
//     return new Promise((resolve, reject) => {
//       getInfo().then(response => {
//         // commit('SET_NAME', response)
//         commit('SET_MENUS', response.menus)
//         commit('SET_NAME', response.name)
//         commit('SET_USERID', response.id)
//         // commit('SET_SUPER', response.superuser)
//         commit('SET_ROLE', response.name)
//         commit('SET_NICKNAME', response.name)
//         // console.log(response)
//
//         resolve(response)
//       }).catch(error => {
//         reject(error)
//       })
//     })
//   },
//   // // get user info
//   // getInfo({ commit, state }) {
//   //   return new Promise((resolve, reject) => {
//   //     getInfo().then(response => {
//   //       const { data } = response
//   //       console.log(data)
//   //       // if (!data) {
//   //       //   reject('Verification failed, please Login again.')
//   //       // }
//   //
//   //       // const { roles, name, avatar, introduction } = data
//   //
//   //       // roles must be a non-empty array
//   //       if (!roles || roles.length <= 0) {
//   //         reject('getInfo: roles must be a non-null array!')
//   //       }
//   //
//   //       commit('SET_ROLES', roles)
//   //       commit('SET_NAME', name)
//   //       commit('SET_AVATAR', avatar)
//   //       commit('SET_INTRODUCTION', introduction)
//   //       resolve(data)
//   //     }).catch(error => {
//   //       reject(error)
//   //     })
//   //   })
//   // },
//
//   // user logout
//   logout({ commit, state }) {
//     return new Promise((resolve, reject) => {
//       logout(state.token).then(() => {
//         commit('SET_TOKEN', '')
//         commit('SET_ROLES', [])
//         removeToken()
//         resetRouter()
//         resolve()
//       }).catch(error => {
//         reject(error)
//       })
//     })
//   },
//   // 前端 登出
//   FedLogOut({ commit }) {
//     return new Promise(resolve => {
//       commit('SET_TOKEN', '')
//       commit('CLEAR_LOCK')
//       removeToken()
//       resolve()
//     })
//   },
//
//   // remove token
//   resetToken({ commit }) {
//     return new Promise(resolve => {
//       commit('SET_TOKEN', '')
//       commit('SET_ROLES', [])
//       removeToken()
//       resolve()
//     })
//   },
//
//   // Dynamically modify permissions
//   changeRoles({ commit, dispatch }, role) {
//     return new Promise(async resolve => {
//       const token = role + '-token'
//
//       commit('SET_TOKEN', token)
//       setToken(token)
//
//       const { roles } = await dispatch('getInfo')
//
//       resetRouter()
//
//       // generate accessible routes map based on roles
//       const accessRoutes = await dispatch('permission-bak/generateRoutes', roles, { root: true })
//
//       // dynamically add accessible routes
//       router.addRoutes(accessRoutes)
//
//       resolve()
//     })
//   }
// }
//
// export default {
//   namespaced: true,
//   state,
//   mutations,
//   actions
// }

import { login, getInfo } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
// import { setStore, getStore, removeStore } from '@/utils/store'
// import { userlogin, logout, getInfo } from '@/api/user'
// import { getToken, setToken, removeToken } from '@/utils/auth'
// import router, { resetRouter } from '@/router'
const user = {
  state: {
    token: getToken(),
    name: '',
    nickname: '',
    userid: '',
    avatar: '',
    roles: []
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, name) => {
      state.name = name
    },
    SET_NICKNAME: (state, name) => {
      state.nickname = name
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    // 前端带单
    SET_MENUS: (state, menus) => {
      state.menus = menus
    },
    SET_USERID: (state, id) => {
      state.userid = id
    },
    SET_SUPER: (state, superuser) => {
      state.superuser = superuser
    },
    SET_ROLE: (state, role) => {
      state.userrole = role
    },
    SET_BROWSERHEADERTITLE: (state, action) => {
      state.browserHeaderTitle = action.browserHeaderTitle
    }
  },

  actions: {
    Login({ commit }, userInfo) {
      const { username, password } = userInfo
      return new Promise((resolve, reject) => {
        login({ username: username.trim(), password: password }).then(response => {
          const { data } = response
          commit('SET_TOKEN', data.token)
          setToken(data.token)
          resolve()
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 获取用户信息
    GetInfo({ commit, state }) {
      return new Promise((resolve, reject) => {
        getInfo().then(response => {
          // commit('SET_NAME', response)
          commit('SET_MENUS', response.menus)
          commit('SET_NAME', response.name)
          commit('SET_USERID', response.id)
          commit('SET_SUPER', response.superuser)
          commit('SET_ROLE', response.name)
          commit('SET_NICKNAME', response.nickname)
          // console.log(response)

          resolve(response)
        }).catch(error => {
          reject(error)
        })
      })
    },
    // 登出，暂时不对后端请求，前端删除JWT
    LogOut({ commit, state }) {
      commit('SET_TOKEN', '')
      commit('SET_NAME', '')
      // commit('CLEAR_LOCK')
      removeToken()
    },

    // 前端 登出
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        removeToken()
        resolve()
      })
    }
  }
}

export default user
