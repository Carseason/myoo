<template>
    <div class="account-comments-row">
        <template v-for="value in Comments">
            <li class="account-comments-list" :key="value.Id">
                <span>{{value.Data | dateForm3}}</span>
                您评论了《
                <router-link :to="$root.GetPostsUrl(value.Posts.Id)">{{value.Posts.Title}}</router-link>》：
                <div class="account-comments-content" v-html="value.Content"></div>
            </li>
        </template>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Comments: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/account/comments/" + this.$route.params.page).then(res => {
            this.Comments = res.data.data ? res.data.data : [];
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
li.account-comments-list {
    display: flex;
    flex-wrap: wrap;
    margin: 10px 0;
    padding: 10px;
    font-size: 16px;
    cursor: pointer;
    border-bottom: 1px solid #d7d9da3d;
}

li.account-comments-list span {
    margin-right: 10px;
}

li.account-comments-list a {
    color: #ffafc9;
    margin: 0 5px;
}
</style>
