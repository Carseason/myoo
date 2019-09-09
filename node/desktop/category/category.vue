<template>
    <div class="body" :key="$route.params.id">
        <div class="category_wrap"></div>
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
            Title: "",
        }
    },
    watch: {
        "$route.params.id"(val) {
            this.Get()
        },
        "Page"(val) {
            this.$router.push("/category/" + this.$route.params.id + "/" + val)
        },
        '$store.state.Title': {
            immediate: true,
            handler() {
                this.$store.commit("title", this.Title);
            }
        },
    },
    created() {
        this.Get()
    },
    methods: {
        Get() {
            this.$axios.get('/category/' + this.$route.params.id + "/0").then(res => {
                if (res.data.success != 200) {
                    this.$router.push("/");
                    return;
                }
                this.Title = res.data.data.Name;
                this.MaxPage = res.data.data.MaxPage;
                this.MaxCount = res.data.data.MaxCount;
                this.$store.commit("title", this.Title);
            }).catch(error => {
                console.log(error);
            });
        }
    }
}
</script>
<style scoped>
.category_wrap {
    margin: 20px 0;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}
</style>