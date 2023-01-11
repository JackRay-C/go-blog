import i18n from './i18n';
import router from "./router/index";
import store from "./store";
import {
  Notification
} from "element-ui";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

NProgress.configure({
  showSpinner: false,
});

function hasRoutePermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role.name))
  } else {
    return true
  }
}


router.beforeEach((to, form, next) => {
  console.log("admin")
  console.log(to)
  NProgress.start()
  if (to.meta.title) {
    document.title = to.meta.title;
  }


  // 需要认证的路由
  if (to.meta && to.meta.requireAuth) {
    if (store.getters.access_token || localStorage.getItem("access_token")) {

      let roles = store.getters.roles

      if (roles.length === 0) {
        store.dispatch('DispatchUserInfo', store.getters.access_token).then(res => {
          console.log(res)
          if (res.code && res.code === 200) {
            roles = res.roles
            if (hasRoutePermission(roles, to)) {
              NProgress.done()
              next(to)
            } else {
              NProgress.done()
              next("/403")
            }
          } else {
            console.log(res.message)
            Notification.error({
              title: res.message
            })
            next(`login?redirect=${to.path}`)
            NProgress.done()
          }

        }).catch(err => {
          Notification.error({
            title: err
          })
          next(`login?redirect=${to.path}`)
          NProgress.done()
          console.log(err)
        })
      } else {
        next()
        NProgress.done()
      }
    } else {
      Notification.error({
        title: i18n.tc("login")
      })
      next(`login?redirect=${to.path}`)
      NProgress.done()
      console.log("no access_token")
    }
  } else {
    next()
  }

})

router.afterEach(() => {
  NProgress.done()
})

// 判断路由是否是通用路由
// function routerConstant(name){
//   let r = constantRouterMap.find(v=> v.name===name)

//   if (r) {
//     return true
//   }
//   return false

//   // let ar = asyncRouterMap.find(v => {console.log(v); if( v.name === name) return true;})

//   // if (ar.meta && ar.meta.requireAuth){
//   //   return true
//   // }
//   // return false
// }


// router.beforeEach(async (to, from, next) => {
//   NProgress.start();

//   // 根据路由的title 设置每个页面的标题
//   if (to.meta.title) {
//     document.title = to.meta.title;
//   }

//   // 1、判断To是否需要认证
//   // 1.1、如果不需要认证，直接跳转
//   // 1.2、如果需要认证，判断是否有token
//   //    1.2.1、没有token，跳转到login进行登录
//   //    1.2.2、有token，获取角色，判断用户的角色是否可以跳转到To的菜单
//   //    1.2.3、没有权限的话跳转到403
//   //    1.2.4、有权限的话跳转
//   // 1、通过to.name查找路由，判断路由是否需要认证
//   console.log(router.getRoutes())

//   if(routerConstant(to.name)) {
//     next()
//   } else {
//     // 需要认证的路由
//     if (store.getters.access_token) {
//       let roles = store.getters.roles
//       if(roles.length === 0) {
//         let res = await store.dispatch('DispatchUserInfo', store.getters.access_token)
//         console.log(res)
//         roles = res.roles
//         store.dispatch("DispatchGenerateRoutes", roles).then(()=> {
//           store.getters.addRouters.forEach(v => {
//             router.addRoute(v)
//           })
//           next({...to, replace: true})
//         })
//       } else {

//       }
//     }
//   }
//   if(routerRequireAuth(to.name)) {
//     if (store.getters.access_token) {
//       let roles = store.getters.roles
//       if(roles.length === 0){
//         // todo: access_token过期通过refresh_token重新请求
//         let res = await store.dispatch('DispatchUserInfo', store.getters.access_token)
//         console.log(res)
//         roles = res.roles
//         store.dispatch("DispatchGenerateRoutes", roles).then(()=> {
//           store.getters.addRouters.forEach(v => {
//             router.addRoute(v)
//           })
//           next({...to, replace: true})
//         })
//       } else {
//         console.log("no roles")
//         next()
//       }

//     } else {
//       // 没有access_token
//       Notification.error({
//         title: 'Error',
//         message: "Please login. "
//       })
//       next(`login?redirect=${to.path}`)
//       NProgress.done()
//       console.log("no access_token")
//     }
//   } else {
//     next()
//     NProgress.done()
//   }


// });

// router.afterEach(() => {
//   NProgress.done();
// });