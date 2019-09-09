<template>
    <div class="account-posts-row">
        <template v-for="value in Posts">
            <li class="account-posts-list" :key="value.Id">
                <div>
                    <img class="thumbnail" v-lazy="value.Thumbnail" />
                </div>
                <div>{{value.Title}}</div>
                <div>
                    <router-link :to="'/category/' + value.Category.Id">{{value.Category.Name}}</router-link>
                </div>
                <template>
                    <div v-if=" value.Type == 'standard' ">文章</div>
                    <div v-else-if=" value.Type == 'image' ">图片</div>
                    <div v-else-if=" value.Type == 'music'">音乐</div>
                    <div v-else-if=" value.Type == 'video' ">视频</div>
                </template>

                <div>{{value.Date | dateForm3 }}</div>
                <template v-if=" value.Status == 'pedding'">
                    <div>待审核</div>
                </template>
                <template v-else>
                    <div>已发布</div>
                    <div>
                        <router-link :to="$root.GetPostsUrl(value.Id)">查看</router-link>
                    </div>
                </template>
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
        this.$axios.get('/account/posts/' + this.$route.params.page).then(res => {
            if (res.data.success != 200) {
                this.$router.push("/");
                return;
            }
            this.Posts = res.data.data ? res.data.data : [];
        }).catch(error => {
            console.log(error);
        });
    }
}
</script>
<style scoped>
li.account-posts-list {
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

li.account-posts-list > div {
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

li.account-posts-list img.thumbnail {
    max-width: 100px;
    width: 100%;
    height: 100%;
}
</style>
