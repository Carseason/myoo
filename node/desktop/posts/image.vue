<template>
    <div class="posts-body">
        <div class="posts-bg"></div>
        <div class="posts-image">
            <div class="posts-context">
                <div class="posts-header">
                    <h1 class="posts-title">{{posts.Title}}</h1>
                    <div class="posts-category">
                        <router-link :to="'/category/' + posts.Category.Id">{{posts.Category.Name}}</router-link>
                    </div>
                    <div class="placeholder"></div>
                    <div class="posts-header-author">
                        <span>作者:</span>
                        <router-link :to="'/author/' + author.Id" class="posts-author-nickname" :title="author.Name">{{author.Name}}</router-link>
                    </div>
                    <div class="posts-date">{{posts.Date | dateForm3 }}</div>
                    <div class="posts-info">
                        <span>浏览：{{posts.Views | count}}</span>
                    </div>
                </div>
                <div class="posts-content" v-html="posts.Content"></div>
            </div>
        </div>
    </div>
</template>
<script>
export default {
    props: ["posts", "author"],

}
</script>
<style scoped>
.posts-image {
    margin-top: 30px;
    text-align: center;
}
h1.posts-title {
    font-size: 2rem;
    color: #676262;
    font-weight: 500;
    padding: 0;
    margin: 0;
    text-align: center;
    margin-bottom: 10px;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}

.posts-header {
    display: flex;
    flex-wrap: wrap;
    font-size: 12px;
    color: #999;
    padding-bottom: 10px;
    margin-bottom: 20px;
    border-bottom: 1px dashed #ddd;
}

.posts-category a {
    font-size: 12px;
    color: #999;
    margin-right: 5px;
}

.posts-date {
    margin-right: 10px;
}
.posts-content {
    font-size: 15px;
    line-height: 1.5;
    min-height: 200px;
}

.posts-header-author {
    margin: 0 10px;
}

.posts-header-author a {
    font-size: 12px;
    color: #999;
}
.posts-row {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: flex;
    flex-wrap: wrap;
}
</style>
