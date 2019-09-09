<template>
    <div class="body">
        <div class="account-bg"></div>
        <router-view></router-view>
        <div class="account-menus">
            <router-link to="/account">
                <i class="iconfont icon-denglu2 icon-zhuye"></i>
                <span>主页</span>
            </router-link>
            <router-link to="/account/info">
                <i class="iconfont icon-denglu2 icon-shezhi1"></i>
                <span>资料</span>
            </router-link>
            <router-link to="/account/newposts">
                <i class="iconfont icon-denglu2 icon-bianji"></i>
                <span>投稿</span>
            </router-link>
            <router-link to="/account/posts/1">
                <i class="iconfont icon-denglu2 icon-rili"></i>
                <span>文章</span>
            </router-link>
            <router-link to="/account/favorite/1">
                <i class="iconfont icon-denglu2 icon-xingbiao"></i>
                <span>收藏</span>
            </router-link>
            <router-link to="/account/comments/1">
                <i class="iconfont icon-denglu2 icon-xinxi"></i>
                <span>评论</span>
            </router-link>
            <!-- <router-link to>
                <i class="iconfont icon-denglu2 icon-tixing"></i>
                <span>消息</span>
            </router-link>-->
            <router-link :to="'/author/' + $store.state.Auth.Id + '/followers/1'">
                <i class="iconfont icon-denglu2 icon-qunzu"></i>
                <span>关注</span>
            </router-link>
            <router-link :to="'/author/' + $store.state.Auth.Id + '/fans/1'">
                <i class="iconfont icon-denglu2 icon-qunzu"></i>
                <span>粉丝</span>
            </router-link>
            <router-link :to="'/author/' + $store.state.Auth.Id">
                <i class="iconfont icon-denglu2 icon-zhuye"></i>
                <span>空间</span>
            </router-link>
            <a @click="OutLogin()">
                <i class="iconfont icon-denglu2 icon-tuichu"></i>
                <span>注销</span>
            </a>
        </div>
    </div>
</template>
<script>
export default {
    watch: {
        '$store.state.Title': {
            immediate: true,
            handler() {
                this.$store.commit("title", "个人中心")
            }
        },
    },
    mounted() {
    },
    methods: {
        OutLogin() {
            this.$store.commit("localstorage", null)
            window.location.href = "/";
        },
    }
}
</script>

<style scoped>
.account-menus {
    width: 100%;
    height: 100px;
    margin: 0 auto;
    position: fixed;
    z-index: 99;
    bottom: 20px;
    left: 0;
    right: 0;
    justify-content: center;
    -ms-flex-wrap: nowrap;
    display: flex;
    flex-wrap: nowrap;
    -webkit-box-pack: start;
    -ms-flex-pack: start;
    overflow-x: auto;
    overflow-y: hidden;
}

.account-menus a {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 60px;
    max-height: 60px;
    height: 100%;
    text-align: center;
    margin: 0 10px;
    font-size: 15px;
    box-shadow: 0 1px 10px rgba(15, 15, 15, 0.34901960784313724);
    background-color: rgb(49, 110, 181);
    border-radius: 20%;
}
.account-menus a:hover {
    max-width: 70px;
    max-height: 70px;
    transition: 0.3s;
}
.account-menus a:nth-child(1) {
    background-color: rgb(49, 110, 181);
}
.account-menus a:nth-child(2) {
    background-color: rgb(255, 153, 190);
}

.account-menus a:nth-child(3) {
    background-color: rgb(136, 237, 191);
}
.account-menus a:nth-child(4) {
    background-color: rgb(141, 101, 196);
}
.account-menus a:nth-child(5) {
    background-color: rgb(68, 167, 223);
}
.account-menus a:nth-child(6) {
    background-color: rgb(255, 181, 131);
}
.account-menus a:nth-child(7) {
    background-color: rgb(119, 117, 116);
}
.account-menus a:nth-child(8) {
    background-color: rgb(91, 180, 255);
}
.account-menus a:nth-child(9) {
    background-color: rgb(91, 180, 255);
}
.account-menus a:nth-child(10) {
    background-color: rgb(54, 54, 54);
}
.account-menus a:nth-child(11) {
    background-color: rgb(249, 129, 140);
}

.account-menus a i {
    width: 100%;
    height: 100%;
    display: block;
    font-size: 2.5rem;
    color: #fff;
    line-height: 2;
}
.account-menus a:hover i[data-v-b34c4d8c] {
    font-size: 3rem;
    transition: 0.3s;
}
.account-menus a span {
    color: #fff;
    font-size: 15px;
    margin-top: 5px;
    display: block;
}

.account-bg {
    width: 100%;
    height: 100%;
    z-index: -1;
    position: fixed;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
    background-image: linear-gradient(
            60deg,
            rgba(255, 165, 150, 0.5) 5%,
            rgba(0, 228, 255, 0.35)
        ),
        url(/src/img/bg.jpg);
    background-size: cover;
}
</style>
