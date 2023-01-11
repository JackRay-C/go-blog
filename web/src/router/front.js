const frontRouterMap = [
    {
        path: '/',
        name: 'index',
        meta: {
            title: '主页',
            layout: 'default',
            sidebar: false,
        },
        component: ()=> import('../views/Home.vue')
    },
    {
        path: '/tag',
        name: 'tag',
        meta: {
            title: '标签页面',
            layout: 'default',
            sidebar: false,
        },
        component: ()=> import('../views/Tag.vue')
    },
    {
        path: '/tag/:id',
        name: 'tagDetail',
        meta: {
            title: '标签详情',
            layout: 'default',
            sidebar: false,
        },
        component: ()=> import('../views/TagDetail.vue')
    },
    {
        path: '/subject',
        name: 'subject',
        meta: {
            title: '专题',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/Subject.vue')
    },
    {
        path: '/about',
        name: 'about',
        meta: {
            title: '关于我',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/About.vue')
    },
    {
        path: '/detail/:id',
        name: 'detail',
        meta: {
            title: '文章详情',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/Detail.vue')
    },
    {
        path: '/subject/:id',
        name: 'subjectDetail',
        meta: {
            title: '专题文章',
            layout: 'default',
            sidebar: false,
        },
        component: () => import('../views/SubjectDetail.vue')
    },
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '登录',
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
    {
        path: '/403',
        name: '403',
        meta: {
            title: '403 Forbidden'
        },
        component: () => import('../views/403.vue')
    },
]

export default frontRouterMap