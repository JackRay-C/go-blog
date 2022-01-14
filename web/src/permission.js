import router from "./router";
import store from "./store";
import { Message } from "element-ui";
import NProgress from "nprogress";
import "nprogress/nprogress.css";

NProgress.configure({
  showSpinner: false,
});

router.beforeEach(async (to, from, next) => {
  NProgress.start();

  // 根据路由的title 设置每个页面的标题
  if (to.meta.title) {
    document.title = to.meta.title;
  }

  // 判断To是否需要认证
  let token = localStorage.getItem("token");
  if (token) {
    const hasRoles = store.getters.roles && store.getters.roles.length > 0;
    if (hasRoles) {
      next();
      NProgress.done();
    } else {
      try {
        const { roles } = await store.dispatch("DispatchInfo");

        const accessRoutes = await store.dispatch(
          "DispatchGenerateRoutes",
          roles
        );
  
        accessRoutes.forEach((route) => {
          router.addRoute(route);
        });
  
        next({...to, replace: true});
        NProgress.done();
      } catch (error) {
        await store.dispatch('DispatchLogout')
        Message.error(error || 'Error')
        next(`login?redirect=${to.path}`)
        NProgress.done()
      }
    }
  } else {
    // 没有token判断是否需要认证
    if (to.meta && to.meta.requireAuth === true) {
      next(`/login?redirect=${to.path}`);
      NProgress.done();
    }
    next();
    NProgress.done();
  }
});

router.afterEach(() => {
  NProgress.done();
});
