import router from "./router";
import store from "./store";
import {
  Notification
} from "element-ui";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

NProgress.configure({
  showSpinner: false,
});


// 判断路由是否有权限
function hasRoutePermission(roles, route) {
  if (route.meta && route.meta.roles) {
    return roles.some(role => route.meta.roles.includes(role.name))
  } else {
    return true
  }
}

router.beforeEach(async (to, from, next) => {
  NProgress.start();

  // 根据路由的title 设置每个页面的标题
  if (to.meta.title) {
    document.title = to.meta.title;
  }

  // 1、判断To是否需要认证
  // 1.1、如果不需要认证，直接跳转
  // 1.2、如果需要认证，判断是否有token
  //    1.2.1、没有token，跳转到login进行登录
  //    1.2.2、有token，获取角色，判断用户的角色是否可以跳转到To的菜单
  //    1.2.3、没有权限的话跳转到403
  //    1.2.4、有权限的话跳转
  if (to.meta && to.meta.requireAuth) {
    let token = localStorage.getItem("token")
    if (token) {
      // 先从store里查询是否已经有角色了，如果有直接判断是否可以跳转，没有的话请求角色
      const hasRoles = store.getters.roles && store.getters.roles.length > 0;
      if (hasRoles) {
        const roles = store.getters.roles
        if (hasRoutePermission(roles, to)) {
          next()
          NProgress.done()
        } else {
          next({
            path: '/403'
          })
          NProgress.done()
        }
      } else {
        try {
          await store.dispatch("DispatchInfo")
          next({
            ...to,
            replace: true
          });
          NProgress.done();
        } catch (error) {
          console.log(error)
          Notification.error({
            title: 'Error',
            message: error || '检查权限错误. '
          })
          next(`login?redirect=${to.path}`)
          NProgress.done()
        }
      }
    } else {
      Notification.error({
        title: 'Error',
        message: "Please login. "
      })
      next(`login?redirect=${to.path}`)
      NProgress.done()
    }
  } else {
    next()
    NProgress.done()
  }
});

router.afterEach(() => {
  NProgress.done();
});