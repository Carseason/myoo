<template>
    <div class="row">
        <h1 class="row-title">
            <router-link to="/admin/user/0">新建用户</router-link>
        </h1>
        <ul class="row-list">
            <li class="my-posts-list">
                <div>头像</div>
                <div>ID</div>
                <div>用户名</div>
                <div>邮箱</div>
                <div>等级</div>
                <div>注册日期</div>
                <div></div>
            </li>
            <transition name="transitions" appear>
                <router-view :key="Page"></router-view>
            </transition>
        </ul>
        <navpage :page.sync="Page" :maxpage="MaxPage" :status="0"></navpage>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Page: this.$route.params.page,
            MaxCount: 0,
            MaxPage: 0,
        }
    },
    watch: {
        "Page"(val) {
            this.$router.push("/admin/users/" + val)
        }
    },
    created() {
        this.$axios.get('/admin/users/0').then(res => {
            if (res.data.success == 403) {
                this.$router.push("/");
                return;
            }
            this.MaxPage = res.data.data.MaxPage;
            this.MaxCount = res.data.data.MaxCount;
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
ul.row-list {
    margin: auto;
    border: 1px solid rgba(255, 255, 255, 0.29);
    box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
    border-radius: 1rem;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}
li.my-posts-list {
    display: flex;
    flex-wrap: wrap;
    position: relative;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    font-size: 14px;
    margin-bottom: 5px;
    border-bottom: 1px solid #0000000d;
    font-weight: 700;
    height: 30px;
    line-height: 30px;
}
li.my-posts-list div {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}
li.my-posts-list > * {
    text-align: center;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 14.2%;
    height: 100%;
    color: #666;
}
</style>
