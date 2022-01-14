import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import filter from './filter'
import Highlight from './directive/highlight'
import DefaultLayout from  './layout/Layout.vue'
import AdminLayout from './layout/Admin.vue'
import EmptyLayout from './layout/Empty.vue'
import ElementUI from 'element-ui';


import 'element-ui/lib/theme-chalk/index.css';

import './permission'

Vue.config.productionTip = false


Vue.component('default-layout', DefaultLayout)
Vue.component('admin-layout', AdminLayout)
Vue.component('empty-layout', EmptyLayout)
Vue.use(Highlight)
Vue.use(ElementUI);


Object.keys(filter).forEach(key => Vue.filter(key, filter[key]))


new Vue({
    router,
    store,
    render: h => h(App)
}).$mount('#app')