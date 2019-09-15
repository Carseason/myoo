import axios from '@/axios/axios.js';

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

export default {
    path: "/", component: Template, name: "", children: [
        { path: '/', component: Index, name: "index" },
        {
            path: 'category/:id', component: Category, name: "category", children: [
                { path: ':page', component: Category_Page, name: "category" },
            ], beforeEnter: (to, from, next) => {
                let path = parseInt(to.params.page)
                if (!Number.isInteger(path) || path <= 0) {
                    next("/category/" + to.params.id + "/1");
                }
                next();
            }
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
};