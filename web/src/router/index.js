import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'



Vue.use(VueRouter)

//---------- 解决重复跳转当前路由报错的问题 -------------------
const originalPush = VueRouter.prototype.push

VueRouter.prototype.push = function push(location) {
    return originalPush.call(this, location).catch(err => err)
}
//-------------------------------------------------------

const routes = [{
        path: '/',
        name: 'Index',
        meta: {
            title: '主页',
            layout: 'default'
        },
        component: Home
    },
    {
        path: '/tag',
        name: 'Tag',
        meta: {
            title: '标签页面',
            layout: 'default'
        },
        component: () => import('../views/Tag.vue')
    },
    {
        path: '/tag/:id',
        name: 'TagDetail',
        meta: {
            title: '标签详情',
            layout: 'default'
        },
        component: () => import('../views/TagDetail.vue')
    },
    {
        path: '/subject',
        name: 'Subject',
        meta: {
            title: '专题',
            layout: 'default'
        },
        component: () => import('../views/Subject.vue')
    },
    {
        path: '/about',
        name: 'About',
        meta: {
            title: '关于我',
            layout: 'default'
        },
        component: () => import('../views/About.vue')
    },
    {
        path: '/detail/:id',
        name: 'Detail',
        meta: {
            title: '文章详情',
            layout: 'default'
        },
        component: () => import('../views/Detail1.vue')
    },
    {
        path: '/subject/:id',
        name: 'SubjectDetail',
        meta: {
            title: '专题文章',
            layout: 'default'
        },
        component: () => import('../views/SubjectDetail.vue')
    },
    {
        path: '/admin/dashboard',
        name: 'AdminDashBoard',
        meta: {
            title: '管理台',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Dashboard.vue')
    },
    {
        path: '/admin/posts',
        name: 'AdminPosts',
        meta: {
            title: '所有博客',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/pages',
        name: 'AdminPages',
        meta: {
            title: '页面管理',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Pages.vue')
    },
    {
        path: '/admin/setting',
        name: 'AdminSetting',
        meta: {
            title: '设置',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Settings.vue')
    },
    {
        path: '/admin/dicts',
        name: 'AdminDicts',
        meta: {
            title: '设置',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Dicts.vue')
    },
    {
        path: '/admin/subject',
        name: 'AdminSubject',
        meta: {
            title: '专题',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Subjects.vue')
    },
    {
        path: '/admin/subject/:id/setting/',
        name: 'AdminSubjectSetting',
        meta: {
            title: '专题设置',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    {
        path: '/admin/subject/:id/post',
        name: 'AdminSubjectPost',
        meta: {
            title: '专题文章',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/subject/new',
        name: 'AdminNewSubject',
        meta: {
            title: '新建专题',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    {
        path: '/admin/draft',
        name: 'AdminDraft',
        meta: {
            title: '专题',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/edit',
        name: 'AdminNew',
        meta: {
            title: '新建文章',
            requireAuth: true
        },
        component: () => import('../views/admin/Edit.vue')
    },
    {
        path: '/admin/edit/:id',
        name: 'AdminEdit',
        meta: {
            title: '编辑文章',
            requireAuth: true
        },
        component: () => import('../views/admin/Edit.vue')
    },
    {
        path: '/admin/tag',
        name: 'AdminTag',
        meta: {
            title: '标签管理',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Tags.vue')
    },
    {
        path: '/admin/users',
        name: 'AdminUser',
        meta: {
            title: '用户管理',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Users.vue')
    },
    {
        path: '/login',
        name: 'Login',
        meta: {
            title: '登录',
            layout: 'default'
        },
        component: () => import('../views/Login.vue')
    },

    {
        path: '/*',
        name: '404',
        meta: {
            title: '404 Not Found'
        },
        component: () => import('../views/404.vue')
    },
]

const router = new VueRouter({
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
    routes
})

router.beforeEach((to, from, next) => {
    // 根据路由的title 设置每个页面的标题
    if (to.meta.title) {
        document.title = to.meta.title
    }

    // 如果下一个路由是login，放行
    if (to.path === '/login') {
        next()
    } else {
        // 否则判断路由元数据有没有requireAuth 属性
        if (to.meta.requireAuth) {
            // 如果有该属性，但是localStorage里面没有发现token缓存，则跳转到登陆界面
            let token = localStorage.getItem('token')
            if (token === null || token === '' || token === undefined) {
                next({
                    path: '/login',
                    query: {
                        redirect: to.fullPath
                    }
                })
            } else {
                // 否则token存在，放行
                next()

            }
        } else {
            // 如果路由没有该属性，直接放行
            next()
        }
    }
})




export default router