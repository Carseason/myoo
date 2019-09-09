<template>
    <ul class="account-favorite">
        <li class="account-favorite-header">
            <div>封面</div>
            <div>标题</div>
            <div>日期</div>
            <div>状态</div>
            <div></div>
        </li>
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
            this.$router.push("/account/favorite/" + val)
        },

    },
    created() {
        this.$axios.get("/account/favorite/0").then(res => {
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
ul.account-favorite {
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
li.account-favorite-header {
    display: flex;
    flex-wrap: wrap;
    position: relative;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    font-size: 14px;
    margin-bottom: 5px;
    border-bottom: 1px solid #0000000d;
    background: #d7edff;
    font-weight: 700;
    height: 30px;
    line-height: 30px;
}

li.account-favorite-header > div {
    text-align: center;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 14.2%;
    height: 100%;
    color: #666;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}
</style>
