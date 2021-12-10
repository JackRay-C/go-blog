import Vue from 'vue'
import Vuex from 'vuex'
import { login } from '../api/login'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    token: '',
    username: '',
    nickname: ''
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
    DispatchLogout({commit}) {
      return new Promise(resolve => {
        commit('SET_TOKEN', '')
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
