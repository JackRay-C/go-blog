import Vue from "vue";
import Vuex from "vuex";
import { getFileById } from "../api/admin/file";
import {
  login,
  getInfoByAccessToken,
  getPermissionByAccessToken,
  getRolesByAccessToken
} from "../api/web/login";
import { routerMap } from "../router";
import { asyncRouterMap } from "../router/admin";


Vue.use(Vuex);

function hasRoutePermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role.name))
  } else {
    return true
  }
}


export default new Vuex.Store({
  state: {
    access_token: "",
    refresh_token: "",
    user_id: "",
    username: "",
    nickname: "",
    avatar: "",
    roles: [],
    permissions: [],
    routes: routerMap,
    addRouters: []
  },
  mutations: {
    SET_TOKEN: (state, token) => {
      state.access_token = token.access_token;
      state.refresh_token = token.refresh_token
      state.expire = token.expire
    },
    SET_USERID: (state, id) => {
      state.user_id = id;
    },
    SET_USERNAME: (state, username) => {
      state.username = username;
    },
    SET_NICKNAME: (state, nickname) => {
      state.nickname = nickname;
    },
    SET_AVATAR: (state, avatar) => {
      state.avatar = avatar;
    },
    SET_ROLES: (state, roles) => {
      state.roles = roles;
    },
    SET_EMAIL: (state, email) => {
      state.email = email;
    },
    SET_PERMISSIONS: (state, permissions) => {
      state.permissions = permissions;
    },
    SET_ROUTES: (state, routes) => {
      state.addRouters = routes;
      state.routes = routerMap.concat(routes)
    },
  },
  actions: {
    DispatchLogin({
      commit
    }, loginForm) {
      return new Promise((resovle, reject) => {
        login(loginForm)
          .then((res) => {
            if (res.code === 200) {
              localStorage.setItem("access_token", res.data.access_token);
              localStorage.setItem("refresh_token", res.data.refresh_token);
              localStorage.setItem("expire", res.data.expire);
             
              commit("SET_TOKEN", res.data);
              resovle(res);
            } else {
              reject(res.message);
            }
          })
          .catch((err) => {
            reject(err);
          });
      });
    },
    async DispatchUserInfo({
      commit
    }, access_token) {
      // 获取用户信息、权限列表、角色列表
      let userInfo = await getInfoByAccessToken(access_token)
      let permissions = await getPermissionByAccessToken(access_token)
      let roles = await getRolesByAccessToken(access_token)
      if (userInfo.code !== 200) {
        return new Promise((resolve, reject) => {
            resolve(userInfo)
        })
      }
      if (permissions.code !== 200) {
        return new Promise((resolve, reject) => {
          resolve(permissions)
        })
      }
      if (roles.code !== 200) {
        return new Promise((resolve, reject) => {
          resolve(roles)
        })
      }
      const {
        username,
        nickname,
        email,
        avatar
      } = userInfo.data
      

      let avatarResp = await getFileById(avatar)
      
      let avatar_url = "http://localhost:8000/static/uploads/2022/6/27/329272434622464.png"
      if(avatarResp && avatarResp.code === 200) {
        avatar_url = avatarResp.data.access_url
      }

      commit("SET_NICKNAME", nickname);
      commit("SET_USERNAME", username);
      commit("SET_AVATAR", avatar_url);
      commit("SET_EMAIL", email);
      commit("SET_PERMISSIONS", permissions.data)
      commit("SET_ROLES", roles.data)
      return new Promise((resovle, reject) => {
        resovle({
          "code": 200,
          "message": "success",
          "user_info": userInfo.data,
          "permissions": permissions.data,
          "roles": roles.data
        })
      })
    },

    DispatchLogout({
      commit
    }) {
      return new Promise((resolve) => {
        commit("SET_TOKEN", "");
        commit("SET_ROLES", []);
        localStorage.removeItem("token");
        localStorage.removeItem("access_token")
        localStorage.removeItem("refresh_token")
        resolve("success");
      });
    },
    DispatchGenerateRoutes({commit}, roles) {
      return new Promise((resolve, reject)=>{
        const accessRoutes = asyncRouterMap.filter(v => {
          if(roles.indexOf('admin') >= 0) return true;
          if(hasRoutePermission(roles, v)) {
            return v
          }
          return false
        });
        localStorage.setItem("addRoutes", JSON.stringify(accessRoutes))
        commit('SET_ROUTES', accessRoutes)
        
        resolve()
      })
    }
  },
  modules: {},
  getters: {
    username: (state) => state.username,
    nickname: (state) => state.nickname,
    avatar: (state) => state.avatar,
    roles: (state) => state.roles,
    permissions: (state) => state.permissions,
    routes: (state) => state.routes,
    addRouters: (state) => state.addRouters || JSON.parse(localStorage.getItem("addRoutes")),
    access_token: (state) => state.access_token || localStorage.getItem("access_token"),
    refresh_token: (state) => state.refresh_token || localStorage.getItem("refresh_token"),
    user_id: (state) => state.user_id,
  }
});