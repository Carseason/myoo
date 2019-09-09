<template>
    <div class="autho-fans">
        <router-view :key="Page"></router-view>
        <navpage :page.sync="Page" :maxpage="MaxPage" :status="0"></navpage>
    </div>
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
            this.$router.push("/author/" + this.$route.params.id + "/fans/" + val)
        },
    },
    created() {
        this.$axios.get("/author/" + this.$route.params.id + "/fans/0").then(res => {
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
</style>