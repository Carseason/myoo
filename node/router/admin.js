
import axios from '@/axios/axios.js';

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
export default {
    path: '/admin/', component: Admin_Index, name: "admin", children: [
        { path: 'statistics', component: Admin_Statistics, name: "admin" },
        { path: 'modifyposts/:id', component: Admin_ModifyPosts, name: "admin" },
        { path: 'theme', component: Admin_Theme, name: "admin" },

        { path: 'options', component: Admin_Options, name: "admin" },
        { path: 'category', component: Admin_Category, name: "admin" },
        { path: 'menus', component: Admin_Menus, name: "admin" },
        { path: 'box', component: Admin_Box, name: "admin" },
        {
            path: 'posts', component: Admin_Posts, name: "admin", children: [
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
};