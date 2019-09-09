<template>
    <div class="body">
        <transition name="transitions" appear>
            <router-view></router-view>
        </transition>
        <div class="sign-router">
            <router-link to="/sign/login" class="sign-router-link">
                <i class="iconfont icon-denglu2 icon-denglu"></i>
                <span>登录</span>
            </router-link>
            <router-link to="/sign/registered" class="sign-router-link" style="background: rgb(253, 139, 139)">
                <i class="iconfont icon-denglu2 icon-tianjiayuzhuceqiye"></i>
                <span>注册</span>
            </router-link>
            <router-link to="/sign/forgetpwd" class="sign-router-link" style="background: rgb(255, 169, 110);">
                <i class="iconfont icon-denglu2 icon-app_icons--"></i>
                <span>修改密码</span>
            </router-link>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Origin: "",

        }
    },
    watch: {
        '$store.state.Title': {
            immediate: true,
            handler() {
                this.$store.commit("title", "登录");
            }
        },
        "$store.state.Auth.Id"() {
            if (this.$store.state.Auth.Id > 0) {
                this.$router.push("/");
            }
        }
    },
    mounted() {
    },
}
</script>
<style scoped>
.sign-router {
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
.sign-router a.sign-router-link {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 60px;
    max-height: 60px;
    height: 100%;
    text-align: center;
    margin: 0 10px;
    font-size: 15px;
    box-shadow: 0 1px 10px rgba(15, 15, 15, 0.34901960784313724);
    background: #00a0d8;
    border-radius: 20%;
}
.sign-router a.sign-router-link:hover {
    max-width: 70px;
    max-height: 70px;
    transition: 0.3s;
}

.sign-router a.sign-router-link:hover i {
    font-size: 3rem;
    transition: 0.3s;
}
a.sign-router-link i {
    width: 100%;
    height: 100%;
    display: block;
    font-size: 2.5rem;
    color: #fff;
    line-height: 2;
}
a.sign-router-link span {
    color: #5b5b5b;
    font-size: 15px;
    margin-top: 5px;
    display: block;
}
</style>
