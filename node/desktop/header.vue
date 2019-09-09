<template>
    <header :class="$store.state.RouterPath">
        <div class="myoo-navmnus">
            <!-- <div class="myoo-navmnus_bg"></div>
            <div class="myoo-navmnus_mask"></div>-->
            <div class="body">
                <div class="wrapper-f">
                    <myoo-navmenus :menus="$store.state.Menus.TopMenus" types="top"></myoo-navmenus>
                </div>
                <div class="placeholder"></div>
                <div class="myoo-wrapper-search">
                    <form class="wrapper-searchform" method="get" @submit.prevent="Search()">
                        <input type="text" accesskey="s" v-model.trim="S" placeholder="搜索些什么呢?" required="required" class="wrapper-search-keyword" />
                        <button type="submit" title="搜索" class="wrapper-search-submit">
                            <i class="iconfont icon-denglu2 icon-sousuo"></i>
                        </button>
                    </form>
                </div>
                <div class="placeholder"></div>
                <div class="wrapper-r">
                    <myoo-auth></myoo-auth>
                </div>
            </div>
        </div>
        <div class="myoo-banner">
            <div class="body">
                <router-link to="/" class="myoo-logo"></router-link>
                <!-- <div class="myoo-search">
                    <router-link to="/rank/all" class="link-ranking">
                        <i class="iconfont icon-denglu2 icon-paihangbang"></i>
                        <span>排行榜</span>
                    </router-link>
                    <form class="searchform" method="get" @submit.prevent="Search()">
                        <input type="text" accesskey="s" v-model.trim="S" placeholder="搜索些什么呢?" required="required" class="search-keyword" />
                        <button type="submit" title="搜索" class="search-submit">
                            <i class="iconfont icon-denglu2 icon-sousuo"></i>
                        </button>
                    </form>
                </div>-->
            </div>
        </div>
        <div class="myoo-wrapper" :class="{active : $store.state.ScrollTop > 180}">
            <div class="body">
                <myoo-navmenus :menus="$store.state.Menus.NavMenus" types="nav"></myoo-navmenus>
            </div>
        </div>
        <div v-if="$store.state.Inc.Ad.Header != '' " class="ad" v-html="$store.state.Inc.Ad.Header"></div>
    </header>
</template>
<script>
import Auth from "@/desktop/auth/auth.vue";
export default {
    data() {
        return {
            S: "",
        }
    },
    components: {
        "myoo-auth": Auth,
    },
    mounted() {
    },
    methods: {
        Search() {
            if (this.S == "") {
                return;
            }
            window.location.href = this.$store.state.Mobile ? this.$store.state.Inc.MobileSearch.replace(/@Myoo@/i, this.S) : this.$store.state.Inc.Search.replace(/@Myoo@/i, this.S);
        }
    }
}
</script>
<style scoped>
header {
    position: relative;
    z-index: 99;
}
/* .myoo-navmnus_bg, */
.myoo-banner {
    /* background-image: url(https://i.loli.net/2019/09/01/SJQ5tLu4afgMdoD.png); */
    background-image: url(/src/img/hf.jpg);
}
.myoo-navmnus {
    /* position: relative; */
    position: fixed;
    width: 100%;
    height: 45px;
    line-height: 45px;
    background: #fff;
    color: #222;
    top: 0;
    z-index: 100;
}
.myoo-banner {
    border-top: 1px solid #e5e5e5;
    /* margin-top: -50px; */
    margin-top: 45px;
    z-index: 98;
    background-color: #eee;
    background-position: center;
    overflow: hidden;
    position: relative;
    vertical-align: middle;
    width: 100%;
    height: 180px;
    object-fit: cover;
    cursor: pointer;
    background-size: 100%;
    /* background-size: cover; */
    background-repeat: no-repeat;
    -moz-background-size: 100% 100%;
}
/* .myoo-navmnus_bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-position: center -10px;
    background-repeat: no-repeat;
    filter: blur(4px);
}
.myoo-navmnus_mask {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: hsla(0, 0%, 100%, 0.4);
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
} */

.myoo-wrapper {
    background: #fff;
}
.myoo-wrapper.active {
    position: fixed;
    top: 45px;
    box-shadow: 0 1px 0 rgba(0, 0, 0, 0.1);
    background: #f8f8f8;
    border-bottom: 1px solid #e5e5e5;
    width: 100%;
    z-index: 99;
}

a.myoo-logo {
    position: absolute;
    width: 230px;
    height: 80px;
    left: 0;
    bottom: 10px;
    background: transparent no-repeat 0;
    background-image: url(/src/img/logo.png);
    background-repeat: no-repeat;
    -moz-background-size: 100% 100%;
    background-size: 100%;
}
a.myoo-logo:hover {
    -webkit-animation: tada 1.5s 0.2s ease both;
    -moz-animation: tada 1.5s 0.2s ease both;
}
.myoo-search {
    position: absolute;
    bottom: 20px;
    right: 0;
    width: 270px;
    height: 32px;
    background-color: rgba(0, 0, 0, 0.12);
    border-radius: 4px;
    font-size: 12px;
    z-index: 10;
    display: flex;
    flex-wrap: wrap;
}

a.link-ranking {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 68px;
    margin-right: 2px;
    text-align: center;
    height: 100%;
    line-height: 32px;
    font-size: 12px;
    font-weight: 700;
    background-color: hsla(0, 0%, 100%, 0.88);
    border-radius: 4px;
    color: #f25d8e;
}

form.searchform {
    background-color: hsla(0, 0%, 100%, 0.88);
    height: 100%;
    border-radius: 4px;
    display: flex;
    flex-wrap: wrap;
}

input.search-keyword {
    width: 152px;
    color: #222;
    font-size: 12px;
    overflow: hidden;
    height: 100%;
    padding: 0 10px;
    border: 0;
    box-shadow: none;
    background-color: transparent;
}

button.search-submit {
    border-radius: 0 4px 4px 0;
    right: 0;
    width: 48px;
    min-width: 0;
    cursor: pointer;
    height: 100%;
    margin: 0;
    padding: 0;
    border: 0;
    font-size: 18px;
}
form.searchform button.search-submit:hover {
    color: #ffafc9;
}
a.link-ranking:hover {
    background-color: #fff;
}

.wrapper-f {
    display: flex;
    flex-wrap: wrap;
}

.wrapper-f li.nav-item > a {
    display: inline-block;
    padding: 0 5px;
    cursor: pointer;
}
.wrapper-f li.nav-item > a:hover {
    background-color: hsla(0, 0%, 100%, 0.3);
}

.wrapper-f li.nav-item i {
    color: #00a1d6;
    font-size: 20px;
    padding-right: 5px;
}

li.nav-item:hover ul.nav-menus_sub {
    opacity: 1;
    transform: scaleX(1);
    visibility: visible;
    transition: all 0.5s;
}
ul.nav-menus_sub {
    background: #fff;
    margin: 0;
    min-width: 12rem;
    opacity: 0;
    padding: 10px 0;
    position: absolute;
    top: 100%;
    transform: scale3d(0.9, 0.7, 1);
    transform-origin: top left;
    visibility: hidden;
    z-index: 1;
    border-radius: 0 0 4px 4px;
    box-shadow: 0 15px 30px 5px rgba(0, 0, 0, 0.15);
}
li.menus-child {
    line-height: 33.33333px;
    list-style-type: none;
    position: relative;
}
a.menus-child-link {
    font-size: 1rem;
    display: block;
    padding: 0 1rem;
    text-shadow: 0 0 2px #fff, 0 0 5px #fff, 0 0 10px #fff;
    white-space: nowrap;
}
.wrapper-f li.nav-item li.menus-child i {
    text-align: center;
    width: 10px;
    font-size: 10px;
    font-weight: 800;
    color: rgba(0, 0, 0, 0.33);
    padding: 0;
}
li.menus-child i.iconfont.icon-denglu2.icon-wangzuo- {
    position: absolute;
    right: -50px;
    opacity: 0;
}
li.menus-child:hover i.iconfont.icon-denglu2.icon-wangzuo- {
    opacity: 1;
    transition: all 0.5s;
    position: relative;
    right: 0;
}
span.menus-child-text {
    text-shadow: 0 0 2px #fff, 0 0 5px #fff, 0 0 10px #fff;
    white-space: nowrap;
}

form.wrapper-searchform {
    vertical-align: top;
    height: 35px;
    line-height: 35px;
    display: flex;
    flex-wrap: wrap;
}

.myoo-wrapper-search {
    width: 350px;
    text-align: center;
    padding: 5px 0;
}

input.wrapper-search-keyword {
    border: 1px solid #e5e5e5;
    border-right: none;
    height: 100%;
    padding: 0 10px;
    border-radius: 3px 0 0 3px;
    width: 300px;
    vertical-align: top;
    color: #999;
    box-sizing: border-box;
}

button.wrapper-search-submit {
    border: none;
    width: 46px;
    height: 100%;
    font-size: 14px;
    background: #fd4c5d;
    color: #fff;
    border-radius: 0 3px 3px 0;
    padding: 0;
    cursor: pointer;
    position: relative;
}
</style>
