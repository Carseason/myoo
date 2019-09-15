<template>
    <div class="rwo">
        <template v-for="value in Posts">
            <li class="my-posts-list" :key="value.Id">
                <div>
                    <img :src="value.Thumbnail" :key="value.Id" class="thumbnail" />
                </div>
                <div>{{value.Title}}</div>
                <div>{{value.Author.Name}}</div>
                <template>
                    <div v-if=" value.Type == 'standard' ">文章</div>
                    <div v-else-if=" value.Type == 'image' ">图片</div>
                    <div v-else-if=" value.Type == 'music'">音乐</div>
                    <div v-else-if=" value.Type == 'video' ">视频</div>
                </template>
                <div>{{value.Date | dateForm2 }}</div>
                <template>
                    <div v-if=" value.Status == 'publish'">已发布</div>
                    <div v-else class="red">待审核</div>
                </template>
                <div>
                    <router-link :to="'/admin/modifyposts/'+value.Id">修改</router-link>
                    <span class="block"></span>
                    <a :href="$root.GetPostsUrl(value.Id)" target="_blank">查看</a>
                </div>
            </li>
        </template>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Posts: []
        }
    },
    created() {
        this.$axios.get('/admin/posts/' + this.$route.params.page).then(res => {
            if (res.data.data != null) {
                this.Posts = res.data.data;
            }
        }).catch(error => {
            console.log(error);
        });
    },
    methods: {}
}
</script>
<style scoped>
li.my-posts-list {
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
li.my-posts-list img.thumbnail {
    max-width: 100px;
    width: 100%;
}
.red {
    color: red;
}
</style>
