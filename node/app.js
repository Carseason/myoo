/*******************Vue*******************/
import Vue from 'vue';
/*******************axios*******************/
import axios from '@/axios/axios.js';
Vue.prototype.$axios = axios;
/*******************vuex*******************/
import store from '@/vuex/vuex.js';
/*******************路由*******************/
import router from '@/router/router.js';
/*******************懒加载*******************/
import VueLazyload from 'vue-lazyload'
Vue.use(VueLazyload, {
    preLoad: 1.3,
    error: '/src/img/lazyload.jpg',
    loading: '/src/img/lazyload.jpg',
    attempt: 1,
    dispatchEvent: true,
})
/*******************css*******************/
import "@/style/style.css";
/*******************组件*******************/
import navpage from "@/modules/navpage.vue"
import loading from "@/modules/loading.vue"
import message from "@/modules/message.vue"
import carousel from "@/modules/carousel.vue"
import box from "@/modules/box.vue"
import user from "@/modules/user.vue"
import widget from "@/modules/widget.vue"
import navmenus from "@/modules/navmenus.vue"
import checkcode from "@/modules/checkcode.vue"
import comments from "@/desktop/comments/comments.vue"
Vue.component('navpage', navpage)
Vue.component('loading', loading)
Vue.component('message', message)
Vue.component('myoo-carousel', carousel)
Vue.component('myoo-box', box)
Vue.component('myoo-user', user)
Vue.component('myoo-widget', widget)
Vue.component('myoo-navmenus', navmenus)
Vue.component('myoo-checkcode', checkcode)
Vue.component('myoo-comments', comments)

/*******************过滤器*******************/
Vue.filter("dateForm", function (time) {
    if (time) {
        var dt = new Date(time);
        var y = dt.getFullYear();
        var m = dt.getMonth() + 1;
        var d = dt.getDate();
        return y + "-" + m + "-" + d;
    } else {
        return null;
    }
})
Vue.filter("dateForm2", function (time) {
    if (time) {
        var dt = new Date(time);
        var y = dt.getFullYear();
        var m = dt.getMonth() + 1;
        var d = dt.getDate();
        var hh = dt.getHours();
        var mm = dt.getMinutes();
        var ss = dt.getSeconds();
        return y + "年" + m + "月" + d + "日" + hh + ":" + mm;
    } else {
        return null;
    }
})

Vue.filter("dateForm3", function (time) {
    if (time) {
        let dt = new Date(time);
        let checktime = dt.getTime();
        let nowtime = new Date().getTime();
        let result = (nowtime - checktime) / 1000;
        if (result <= 60) {
            return "刚刚";
        } else if (result > 60 && result <= 3600) {
            return Math.ceil((result / (60))) + '分钟前';
        } else if (result > 3600 && result <= 86400) {
            return Math.ceil((result / (3600))) + '小时前';
        } else if (result > 86400 && result <= 2592000) {
            return Math.ceil((result / (86400))) + '天前';
        } else if (result > 2592000 && result <= 31104000) {
            return Math.ceil((result / (2592000))) + '个月前';
        } else {
            var y = dt.getFullYear();
            var m = dt.getMonth() + 1;
            var d = dt.getDate();
            return y + "年" + m + "月" + d + "日";
        }
        return result;
    } else {
        return null;
    }
})

Vue.filter("count", function (e) {
    return e;
})

const Myoo = new Vue({
    el: '#main',
    data: {
    },
    mounted() {
        this.$store.commit("scrollTop")
    },
    methods: {
        GetPostsUrl(e) {
            return '/av/' + e;
        },
    },
    beforeCreate() {
        this.$axios.all([
            this.$axios.get('/inc'),
            this.$axios.get('/menus'),
            this.$axios.get('/auth'),
        ]).then(this.$axios.spread((inc, menus, auth) => {
            if (inc.data.data != null) {
                this.$store.commit("inc", inc.data.data);
                this.$store.commit("title");
            }
            if (menus.data.data != null) {
                this.$store.commit("menus", menus.data.data);
            }
            if (auth.data.success === 200 && auth.data.data != null) {
                this.$store.commit("auth", auth.data.data);
            }

        }));
    },
    created() {
    },
    components: {
    },
    router,
    store
});