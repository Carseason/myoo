/*Vue路由*/
import Vue from 'vue';
import store from '@/vuex/vuex.js';
import axios from '@/axios/axios.js';
import VueRouter from 'vue-router'
Vue.use(VueRouter);

import Template from "@/desktop/desktop.vue"
import Index from "@/desktop/index/index.vue"

import Category from "@/desktop/category/category.vue"
import Category_Page from "@/desktop/category/category_page.vue"

import Sign from "@/desktop/sign/sign.vue"
import Sign_Login from "@/desktop/sign/login.vue"
import Sign_Registered from "@/desktop/sign/registered.vue"
import Sign_Forgetpwd from "@/desktop/sign/forgetpwd.vue"
import Sign_Resetpwd from "@/desktop/sign/resetpwd.vue"


import Author from "@/desktop/author/author.vue"
import Author_Dynamic from "@/desktop/author/dynamic.vue"
import Author_Posts from "@/desktop/author/posts.vue"
import Author_Posts_Page from "@/desktop/author/posts_page.vue"
import Author_Followers from "@/desktop/author/followers.vue"
import Author_Followers_Page from "@/desktop/author/followers_page.vue"
import Author_Fans from "@/desktop/author/fans.vue"
import Author_Fans_Page from "@/desktop/author/fans_page.vue"




import Account from "@/desktop/account/account.vue"
import Account_Info from "@/desktop/account/info.vue"
import Account_NewPosts from "@/desktop/account/newposts.vue"
import Account_Posts from "@/desktop/account/posts.vue"
import Account_Posts_Page from "@/desktop/account/posts_page.vue"
import Account_Favorite from "@/desktop/account/favorite.vue"
import Account_Favorite_Page from "@/desktop/account/favorite_page.vue"
import Account_Comments from "@/desktop/account/comments.vue"
import Account_Comments_Page from "@/desktop/account/comments_page.vue"



import Posts from "@/desktop/posts/posts.vue"

import Download from "@/desktop/download/download.vue"


import Admin_Index from "@/admin/index.vue"
import Admin_Options from "@/admin/options.vue"
import Admin_Statistics from "@/admin/statistics.vue"
import Admin_Category from "@/admin/category.vue"
import Admin_Menus from "@/admin/menus.vue"
import Admin_Box from "@/admin/box.vue"
import Admin_Posts from "@/admin/posts.vue"
import Admin_Posts_Page from "@/admin/posts_page.vue"
import Admin_User from "@/admin/user.vue"
import Admin_Users from "@/admin/users.vue"
import Admin_Users_Page from "@/admin/users_page.vue"
import Admin_Group from "@/admin/group.vue"
import Admin_ModifyPosts from "@/admin/modifyposts.vue"
import Admin_Theme from "@/admin/theme.vue"

const router = new VueRouter({  //创建路由对象
    mode: 'history',        //hash模式，history模式去掉#
    routes: [
        {
            path: "/", component: Template, name: "", children: [
                { path: '/', component: Index, name: "index" },
                {
                    path: 'category/:id', component: Category, name: "category", children: [
                        { path: ':page', component: Category_Page, name: "category" },
                    ]
                },
                {
                    path: '/sign', component: Sign, name: "sign", children: [
                        { path: 'login', component: Sign_Login, name: "sign" },
                        { path: 'registered', component: Sign_Registered, name: "sign" },
                        { path: 'forgetpwd', component: Sign_Forgetpwd, name: "sign" },
                        { path: 'resetpwd/:id', component: Sign_Resetpwd, name: "sign" },
                    ]
                },
                {
                    path: 'author/:id', component: Author, name: "author", children: [

                        { path: '', component: Author_Dynamic, name: "author" },
                        {
                            path: 'posts', component: Author_Posts, name: "author", children: [
                                { path: ':page', component: Author_Posts_Page, name: "author" },
                            ]
                        },
                        {
                            path: 'followers', component: Author_Followers, name: "author", children: [
                                { path: ':page', component: Author_Followers_Page, name: "author" },
                            ]
                        },
                        {
                            path: 'fans', component: Author_Fans, name: "author", children: [
                                { path: ':page', component: Author_Fans_Page, name: "author" },
                            ]
                        },
                    ]
                },
                {
                    path: 'account', component: Account, name: "account", children: [
                        { path: 'info', component: Account_Info, name: "account" },
                        { path: 'newposts', component: Account_NewPosts, name: "account" },
                        {
                            path: 'posts', component: Account_Posts, name: "account", children: [
                                { path: ':page', component: Account_Posts_Page, name: "account" },
                            ]
                        },
                        {
                            path: 'favorite', component: Account_Favorite, name: "account", children: [
                                { path: ':page', component: Account_Favorite_Page, name: "account" },
                            ]
                        },
                        {
                            path: 'comments', component: Account_Comments, name: "account", children: [
                                { path: ':page', component: Account_Comments_Page, name: "account" },
                            ]
                        },
                    ], beforeEnter: (to, from, next) => {
                        axios.get("/account").then(res => {
                            if (res.data.success === 200) {
                                next()
                            } else {
                                next("/")
                            }
                        }).catch(error => {
                            next("/")
                        });
                    }

                },
                { path: '/av/:id', component: Posts, name: "posts" },
                { path: '/download/:id', component: Download, name: "download" },

            ]
        },

        {
            path: '/admin/', component: Admin_Index, name: "admin", children: [
                { path: 'statistics', component: Admin_Statistics, name: "admin" },
                { path: 'modifyposts/:id', component: Admin_ModifyPosts, name: "admin" },
                { path: 'theme', component: Admin_Theme, name: "admin" },

                { path: 'options', component: Admin_Options, name: "admin" },
                { path: 'category', component: Admin_Category, name: "admin" },
                { path: 'menus', component: Admin_Menus, name: "admin" },
                { path: 'box', component: Admin_Box, name: "admin" },
                {
                    path: 'posts', component: Admin_Posts, name: "admin",
                    children: [
                        { path: ':page', component: Admin_Posts_Page, name: "admin" }
                    ]
                },
                { path: 'user/:id', component: Admin_User, name: "admin" },
                { path: 'group', component: Admin_Group, name: "admin" },
                {
                    path: 'users', component: Admin_Users, name: "admin",
                    children: [
                        { path: ':page', component: Admin_Users_Page, name: "admin" }
                    ]
                },
            ], beforeEnter: (to, from, next) => {
                axios.get("/admin").then(res => {
                    if (res.data.success === 200) {
                        next()
                    } else {
                        next("/")
                    }
                }).catch(error => {
                    next("/")
                });
            }
        },
    ],

    /*覆盖默认路由高亮类*/
    linkActiveClass: 'on'
});


router.beforeEach((to, from, next) => {
    store.state.RouterPath = to.name;
    switch (to.name) {
        case "category":
            if (to.params.page == null) {
                next("/category/" + to.params.id + "/1");
            }
    }
    next()
})
export default router;

