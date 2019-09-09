<template>
    <div class="account-favorite-row">
        <template v-for="value in Posts">
            <li class="account-favorite-list" :key="value.Id">
                <div>
                    <img class="thumbnail" v-lazy="value.Thumbnail" />
                </div>
                <div>{{value.Title}}</div>

                <div>{{value.Date | dateForm3 }}</div>

                <div>
                    <router-link :to="$root.GetPostsUrl(value.Id)">查看</router-link>
                </div>
            </li>
        </template>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Posts: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/account/favorite/" + this.$route.params.page).then(res => {
            this.Posts = res.data.data ? res.data.data : [];
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
li.account-favorite-list {
    height: 70px;
    line-height: 70px;
    display: flex;
    flex-wrap: wrap;
    position: relative;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    font-size: 14px;
    margin-bottom: 5px;
    border-bottom: 1px solid #0000000d;
}

li.account-favorite-list > div {
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

li.account-favorite-list img.thumbnail {
    max-width: 100px;
    width: 100%;
    height: 100%;
}
</style>