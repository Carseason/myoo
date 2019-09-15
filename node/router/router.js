/*Vue路由*/
import Vue from 'vue';
import store from '@/vuex/vuex.js';
import axios from '@/axios/axios.js';
import VueRouter from 'vue-router'
Vue.use(VueRouter);


import desktop from "@/router/desktop.js"
import admin from "@/router/admin.js"
const router = new VueRouter({  //创建路由对象
    mode: 'history',        //hash模式，history模式去掉#
    routes: [
        desktop,
        admin
    ],

    /*覆盖默认路由高亮类*/
    linkActiveClass: 'on'
});


router.beforeEach((to, from, next) => {
    store.state.RouterPath = to.name;
    next()
})
export default router;

