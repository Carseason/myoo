<template>
    <ul class="account-comments">
        <div class="account-comments-header">
            <i class="iconfont icon-denglu2 icon-xinxi"></i>我的评论
        </div>
        <router-view :key="Page"></router-view>
        <navpage :page.sync="Page" :maxpage="MaxPage" :status="0"></navpage>
    </ul>
</template>
<script>
export default {
    data() {
        return {
            Page: this.$route.params.page ? this.$route.params.page : 1,
            MaxCount: 0,
            MaxPage: 0,
        }
    },
    watch: {
        Page(val) {
            this.$router.push("/account/comments/" + val)
        },
    },
    created() {
        this.$axios.get("/account/comments/0").then(res => {
            if (res.data.success != 200) {
                this.$router.push("/");
                return;
            }
            this.MaxPage = res.data.data.MaxPage;
            this.MaxCount = res.data.data.MaxCount;
        }).catch(error => {
            console.log(error);
        });
    }
}
</script>
<style scoped>
ul.account-comments {
    background: #ffffff;
    width: 100%;
    box-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
    padding: 20px;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    position: relative;
    margin-bottom: 10px;
    border-radius: 5px;
    min-height: 600px;
}
.account-comments-header {
    font-size: 18px;
    color: #222;
    vertical-align: top;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    height: 50px;
    line-height: 50px;
}

.account-comments-header i {
    margin-right: 10px;
    color: #717171;
}
</style>
