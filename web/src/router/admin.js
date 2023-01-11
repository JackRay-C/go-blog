
const adminRouterMap = [
    {
        path: '/admin/dashboard',
        name: 'dashboard',
        meta: {
            title: '管理台',
            layout: 'admin',
            icon: 'dashboard',
            iconClass: 'dashboard_svg__a',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Dashboard.vue')
    },
    {
        path: '/admin/posts',
        name: 'posts',
        meta: {
            title: '所有博客',
            layout: 'admin',
            icon: 'edit',
            iconClass: 'edit_svg__a',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Posts.vue')
    },
    {
        path: '/admin/drafts',
        name: 'drafts',
        meta: {
            title: '草稿箱',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Drafts.vue')
    },
    {
        path: '/admin/published',
        name: 'published',
        meta: {
            title: '已发表',
            layout: 'admin',
            requireAuth: true
        },
        component: () => import('../views/admin/Published.vue')
    },
    {
        path: '/admin/profile',
        name: 'profile',
        meta: {
            title: '个人设置',
            layout: 'admin',
            sidebar: false,
            requireAuth: true
        },
        component: () => import('../views/admin/Profile.vue')
    },
    {
        path: '/admin/pages',
        name: 'pages',
        meta: {
            title: '页面管理',
            layout: 'admin',
            icon: 'pages',
            iconClass: 'page_svg__a',
            requireAuth: true,
            sidebar: true,
            roles: ["Admin"]
        },
        component: () => import('../views/admin/Pages.vue')
        // component: ()=> import('../views/admin/Posts1.vue')
    },
    {
        path: '/admin/setting',
        name: 'settings',
        meta: {
            title: '设置',
            layout: 'admin',
            icon: 'settings',
            iconClass: '',
            requireAuth: true,
            sidebar: false,
            roles: ["Admin"]
        },
        component: () => import('../views/admin/Settings.vue')
    },
    {
        path: '/admin/dicts',
        name: 'dicts',
        meta: {
            title: '字典',
            layout: 'admin',
            requireAuth: true,
            icon: 'dicts',
            iconClass: '',
            sidebar: false,
            roles: ["Admin"]
        },
        component: () => import('../views/admin/Dicts.vue')
    },
    {
        path: '/admin/subjects',
        name: 'subjects',
        meta: {
            title: '专题',
            layout: 'admin',
            icon: 'subjects',
            iconClass: 'subjects_svg__a',
            requireAuth: true,
            sidebar: true,
        },
        component: () => import('../views/admin/Subjects.vue')
    },
    {
        path: '/admin/subjects/:id/setting/',
        name: 'AdminSubjectSetting',
        meta: {
            title: '专题设置',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
            roles: ["Admin", "Editor"]
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    {
        path: '/admin/subjects/:id/post',
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
        path: '/admin/subjects/new',
        name: 'AdminNewSubject',
        meta: {
            title: '新建专题',
            layout: 'admin',
            requireAuth: true,
            sidebar: false,
            roles: ["Admin", "Editor"]
        },
        component: () => import('../views/admin/NewSubject.vue')
    },
    // {
    //     path: '/admin/drafts',
    //     name: 'AdminDraft',
    //     meta: {
    //         title: '草稿',
    //         layout: 'admin',
    //         sidebar: false,
    //         requireAuth: true,
    //         roles: ["Admin", "Editor"]
    //     },
    //     component: () => import('../views/admin/Drafts.vue')
    // },
    {
        path: '/admin/edit/:id',
        name: 'AdminEdit',
        meta: {
            title: '编辑文章',
            requireAuth: true,
            sidebar: false,
            roles: ["Admin", "Editor"]
        },
        component: () => import('../views/admin/Editor.vue')
    },
    {
        path:"/admin/new/:id",
        name: 'AdminNew',
        meta: {
            title: '新建文章',
            requireAuth: true,
            sidebar: false,
            roles: ["Admin", "Editor"]
        },
        component: ()=> import('../views/admin/Editor.vue')
    },
    // {
    //     path: '/admin/edit/:id',
    //     name: 'AdminEdit',
    //     meta: {
    //         title: '编辑文章',
    //         requireAuth: true,
    //         sidebar: false,
    //         roles: ["Admin", "Editor"]
    //     },
    //     component: () => import('../views/admin/Edit.vue')
    // },
    {
        path: '/admin/tags',
        name: 'tags',
        meta: {
            title: '标签管理',
            layout: 'admin',
            icon: 'tags',
            iconClass: 'tag_svg__a',
            sidebar: true,
            requireAuth: true,
            roles: ["Admin", "Editor"]
        },
        component: () => import('../views/admin/Tags.vue')
    },
    {
        path: '/admin/accounts',
        name: 'accounts',
        meta: {
            title: '账号管理',
            layout: 'admin',
            icon: 'members',
            iconClass: 'members_svg__cls-1',
            sidebar: true,
            requireAuth: true,
            roles: ["Admin"]
        },
        component: () => import('../views/admin/Accounts.vue')
    },
    
    
]



export default adminRouterMap