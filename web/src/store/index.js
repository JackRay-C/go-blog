import Vue from 'vue'
import Vuex from 'vuex'
import { login, getUserInfo } from '../api/web/login'
import {webRoutes} from '../router'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: '',
    username: '',
    nickname: '',
    avatar: '',
    roles: [],
    permissions: [],
    routes: []
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
      state.routes = routes
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
          commit('SET_ROUTES', webRoutes)

          resovle(data)

        }).catch(error => {
          reject(error)
        })
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
    token: state => state.token,
    username: state=> state.username,
    nickname: state => state.nickname,
    avatar: state => state.avatar,
    roles: state => state.roles,
    permissions: state => state.permissions,
    routes: state => state.routes,
  }
})

