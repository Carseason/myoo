import Vue from 'vue';
import Vuex from 'vuex'
import axios from '@/axios/axios.js';
Vue.use(Vuex)
export default new Vuex.Store({
    state: {
        Mobile: (/Mobile/i.test(navigator.userAgent)),
        RouterPath: "",
        ScrollTop: 0,
        ScrollStatus: false,
        Title: "",
        Auth: {
            Id: 0,
            Name: "",
            Email: "",
            Date: "",
            Avatar: "",
            Description: "",
            Group: {
                Lv: 0,
                NickName: "",
                Medal: [],
            }
        },
        Message: {
            Status: false,
            Title: "消息",
            Content: "",
        },
        Loading: false,
        Inc: {
            Sitename: "",
            Sitedescription: "",
            Message: "",
            Searach: "",
            MobileSearch: "",
            Ad: {
                Header: "",
                Footer: "",
            }
        },
        Menus: {
            Links: [],
            TopMenus: [],
            NavMenus: []
        },
        CheckCode: {
            Key: "",
            BaseCode: "",
        }
    },
    mutations: {    //方法
        auth: (state, val) => {
            state.Auth = val;
        },
        localstorage: (state, val) => {
            if (val) {
                localStorage.setItem("MYOO", val);
                return
            } else {
                localStorage.removeItem("MYOO");
            }
        },
        msg: (state, val) => {
            state.Message.Status = !state.Message.Status;
            state.Message.Content = val;
        },
        load: (state, status) => {
            [state.Loading = !state.Loading] = [status]
        },
        scrollTop: (state) => {
            window.addEventListener('scroll', () => {
                state.ScrollTop = document.documentElement.scrollTop || document.body.scrollTop;
                state.ScrollStatus = state.ScrollTop >= 210
            });
        },
        inc: (state, val) => {
            state.Inc = val
        },
        title: (state, e) => {
            if (e == null) {
                state.Title = state.Inc.Sitename + " - " + state.Inc.Sitedescription;
            } else {
                state.Title = e + " - " + state.Inc.Sitename;
            }
            document.title = state.Title;
            return;
        },
        menus: (state, val) => {
            if (val.NavMenus != null) {
                state.Menus.NavMenus = val.NavMenus
            }
            if (val.Links != null) {
                state.Menus.Links = val.Links
            }
            if (val.TopMenus != null) {
                state.Menus.TopMenus = val.TopMenus
            }

        },
        checkcode: (state) => {
            axios.get('/checkcode').then(res => {
                if (res.data.data != null) {
                    state.CheckCode = res.data.data
                }
            });
        },
    }
})