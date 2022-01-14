import Vue from 'vue'
import Vuex from 'vuex'
import { login, getUserInfo } from '../api/web/login'
import {webRoutes, asyncRoutes, filterRouter} from '../router'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: '',
    username: '',
    nickname: '',
    avatar: '',
    roles: [],
    permissions: [],
    routes: [],
    addRoutes: [],
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_USERNAME: (state, username) => {
      state.username = username
    },
    SET_NICKNAME: (state, nickname) => {
      state.nickname = nickname
    },
    SET_AVATAR: (state, avatar ) => {
      state.avatar = avatar
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles
    },
    SET_EMAIL: (state, email) => {
      state.email = email
    },
    SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions
    },
    SET_ROUTES: (state, routes) => {
      state.addRoutes = routes
      state.routes = webRoutes.concat(routes)
    }
  },
  actions: {
    DispatchLogin({commit}, loginForm) {
      return new Promise((resovle, reject) => {
        login(loginForm).then(res => {
          if(res.code === 200) {
            commit('SET_TOKEN', res.data.token)
            localStorage.setItem("token", res.data.token)
            resovle(res)
          } else {
            reject(res)
          }
        }).catch(err => {
          reject(err)
        })
      })
    },
    DispatchInfo({commit, }) {
      return new Promise((resovle, reject) => {
        getUserInfo().then(res => {
    
          const {data} = res 
          if (!data) {
            reject(res.message)
          }

          const {username, nickname, email, avatar, roles, permissions } = data 

          if (!roles || roles.length <= 0) {
            reject('getInfo: roles must be a non-null array')
          }

          commit('SET_NICKNAME', nickname)
          commit('SET_USERNAME', username)
          commit('SET_AVATAR', avatar)
          commit('SET_EMAIL', email)
          commit('SET_ROLES', roles)
          commit('SET_PERMISSIONS', permissions)

          resovle(data)

        }).catch(error => {
          reject(error)
        })
      })
    },
    DispatchGenerateRoutes({commit}, roles) {
      return new Promise(resolve => {
        let accessedRoutes
        roles.forEach(role => {
          if(role.name === 'Admin') {
            accessedRoutes = asyncRoutes || []
          } else {
            accessedRoutes = filterRouter(asyncRoutes, roles)
          }
        });
        commit('SET_ROUTES', accessedRoutes)
        resolve(accessedRoutes)
      })
    },
    DispatchLogout({commit}) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        localStorage.removeItem("token")
        resolve()
      })
    }
  },
  modules: {},
  getters: {
    token: state => state.token || localStorage.getItem("token"),
    username: state=> state.username
  }
})

