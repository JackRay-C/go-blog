
import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Tag from "../views/Tag.vue"
import TagDetail from "../views/TagDetail.vue"



Vue.use(VueRouter)

// 解决重复跳转当前路由报错的问题
const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}


export const webRoutes = [
    {
        path: '/',
        name: 'Index',
        meta: {
            title: '主页',
            layout: 'default',
            sidebar: false,
        },
        component: Home
    },
    {
        path: '/tag',
        name: 'Tag',
        meta: {
            title: '标签页面',
            layout: 'default',
            sidebar: false,
        },
        component: Tag
    },
    {
        path: '/tag/:id',
        name: 'TagDetail',
        meta: {
            title: '标签详情',
            layout: 'default',
            sidebar: false,
        },
        component: TagDetail
    },
    {
        path: '/subject',
        name: 'Subject',
        meta: {
            title: '专题',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/Subject.vue')
    },
    {
        path: '/about',
        name: 'About',
        meta: {
            title: '关于我',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/About.vue')
    },
    {
        path: '/detail/:id',
        name: 'Detail',
        meta: {
            title: '文章详情',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/Detail.vue')
    },
    {
        path: '/subject/:id',
        name: 'SubjectDetail',
        meta: {
            title: '专题文章',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/SubjectDetail.vue')
    },
    {
        path: '/login',
        name: 'Login',
        meta: {
            title: '登录',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/Login.vue')
    },
    {
        path: '/404',
        name: '404',
        meta: {
            title: '404 Not Found'
        },
        component: () => import('../views/404.vue')
    },
]

export const asyncRoutes = [
    {
        path: '/admin/dashboard',
        name: 'AdminDashBoard',
        meta: {
            title: '管理台',
            layout: 'admin',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Dashboard.vue')
    },
    {
        path: '/admin/posts',
        name: 'AdminPosts',
        meta: {
            title: '所有博客',
            layout: 'admin',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/pages',
        name: 'AdminPages',
        meta: {
            title: '页面管理',
            layout: 'admin',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Pages.vue')
    },
    {
        path: '/admin/setting',
        name: 'AdminSetting',
        meta: {
            title: '设置',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
        },
        component: () => import('../views/admin/Settings.vue')
    },
    {
        path: '/admin/dicts',
        name: 'AdminDicts',
        meta: {
            title: '设置',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
            roles: ['Admin']
        },
        component: () => import('../views/admin/Dicts.vue')
    },
    {
        path: '/admin/subject',
        name: 'AdminSubject',
        meta: {
            title: '专题',
            layout: 'admin',
            requireAuth: true,
            sidebar: true,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/Subjects.vue')
    },
    {
        path: '/admin/subject/:id/setting/',
        name: 'AdminSubjectSetting',
        meta: {
            title: '专题设置',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    {
        path: '/admin/subject/:id/post',
        name: 'AdminSubjectPost',
        meta: {
            title: '专题文章',
            layout: 'admin',
            sidebar: false,
            requireAuth: true,
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/subject/new',
        name: 'AdminNewSubject',
        meta: {
            title: '新建专题',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    {
        path: '/admin/draft',
        name: 'AdminDraft',
        meta: {
            title: '草稿',
            layout: 'admin',
            sidebar: true,
            requireAuth: true,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/edit',
        name: 'AdminNew',
        meta: {
            title: '新建文章',
            requireAuth: true,
            sidebar: false,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/Edit.vue')
    },
    {
        path: '/admin/edit/:id',
        name: 'AdminEdit',
        meta: {
            title: '编辑文章',
            requireAuth: true,
            sidebar: false,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/Edit.vue')
    },
    {
        path: '/admin/tag',
        name: 'AdminTag',
        meta: {
            title: '标签管理',
            layout: 'admin',
            sidebar: true,
            requireAuth: true,
            roles: ['Admin', 'Editor']
        },
        component: () => import('../views/admin/Tags.vue')
    },
    {
        path: '/admin/users',
        name: 'AdminUser',
        meta: {
            title: '用户管理',
            layout: 'admin',
            sidebar: true,
            requireAuth: true,
            roles: ['Admin']
        },
        component: () => import('../views/admin/Users.vue')
    },
    { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new VueRouter({
    mode: 'history',
    linkExactActiveClass: 'active',
    base: process.env.BASE_URL,
    scrollBehavior(to, from, savedPosition) {
        if (savedPosition && to.meta.keepAlive) {
            return savedPosition
        }
        return {
            x: 0,
            y: 0
        }
    },
    routes: webRoutes
})

const router = createRouter()

export function resetRouter() {
    const newRouter = createRouter()
    router.matcher = newRouter.matcher
}


// 判断路由是否有权限
export function hasRouter(roles, route) {
    if (route.meta && route.meta.roles) {
        return roles.some(role => route.meta.roles.includes(role))
    } else {
        return true
    }
}


// 过滤出角色拥有的路由
export function filterRouter(routes, roles) {
    const accessRoutes = []

    routes.forEach(route => {
        if (hasRouter(roles, route)) {
            accessRoutes.push(route)
        }
    })
    return accessRoutes
}


export default router