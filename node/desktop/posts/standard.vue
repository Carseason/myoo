<template>
    <div class="posts-body">
        <div class="posts-bg"></div>
        <div class="posts-sidebar">
            <div class="posts-box">
                <div class="posts-context">
                    <div class="posts-header">
                        <h1 class="posts-title">{{posts.Title}}</h1>
                        <div class="posts-category">
                            <router-link :to="'/category/' + posts.Category.Id">{{posts.Category.Name}}</router-link>
                        </div>
                        <div class="placeholder"></div>
                        <div class="posts-date">{{posts.Date | dateForm3 }}</div>
                        <div class="posts-info">
                            <span>浏览：{{posts.Views | count}}</span>
                        </div>
                    </div>
                    <div class="posts-content" v-html="posts.Content"></div>
                </div>
            </div>
        </div>
        <div class="posts-widget">
            <div class="posts-widget-item">
                <div class="posts-author">
                    <router-link :to="'/author/' + author.Id" class="posts-author-avatar">
                        <img :src="author.Avatar" class="avatar" />
                    </router-link>
                    <router-link :to="'/author/' + author.Id" class="posts-author-nickname" :title="author.Name">{{author.Name}}</router-link>
                    <a class="posts-author-group">
                        <span class="capabilities">{{author.Group.NickName}}</span>
                    </a>
                    <a class="posts-author-description">{{author.Description}}</a>
                    <a class="posts-author-info">
                        <span>投稿:{{author.Posts.Count | count }}</span>
                        <span class="block"></span>
                        <span>评论:{{author.Comments.Count | count}}</span>
                        <span class="block"></span>
                        <span>粉丝:{{author.Fans.Count | count}}</span>
                    </a>
                    <a class="author-followe">
                        <template v-if="$store.state.Auth.Id == 0 || $store.state.Auth.Id  == author.Id">
                            <button class="gz isActive">+ 关注</button>
                        </template>
                        <template v-else>
                            <button class="gz" v-if="author.Followers.Status" @click="$parent.UpdateFollowers()" title="取消关注">已关注</button>
                            <button class="gz" v-else @click="$parent.UpdateFollowers()" title="关注作者">+ 关注</button>
                        </template>
                    </a>
                    <a class="posts-author-medal">
                        <template v-for="(value ,i ) in author.Group.Medal">
                            <img :src="value" class="medal" :key="i" />
                        </template>
                    </a>
                </div>
                <div class="posts-box">
                    <myoo-features :posts="posts"></myoo-features>
                </div>
            </div>
        </div>
        <div class="posts-comments">
            <div class="posts-box">
                <myoo-comments :id="posts.Id"></myoo-comments>
            </div>
        </div>
    </div>
</template>
<script>
import features from "./features.vue";

export default {
    components: {
        "myoo-features": features,
    },
    props: ["posts", "author"],

}
</script>
<style scoped>
.posts-bg {
    position: fixed;
    top: 0;
    left: 0;
    z-index: -1;
    width: 100%;
    height: 100%;
    background-color: rgb(240, 240, 240);
}
.body {
    margin-top: 50px;
}

.posts-bg {
    position: fixed;
    top: 0;
    left: 0;
    z-index: -1;
    width: 100%;
    height: 100%;
    background-color: rgb(242, 243, 245);
}

.posts-body {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    height: 100%;
    margin-top: 50px;
    position: relative;
}

.posts-sidebar,
.posts-comments {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(80% - 30px);
    margin-right: 30px;
}

.posts-widget {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 20%;
    position: absolute;
    width: 100%;
    right: 0;
    top: 0;
    bottom: 0;
}
.posts-box {
    background-color: #fff;
    border: 1px solid #e3e8ec;
    border-radius: 12px;
    width: 100%;
    margin-bottom: 20px;
    padding: 30px;
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

.posts-content img {
    max-width: 100%;
}
.posts-widget-item {
    transition: all 0.3s;
    position: fixed;
    z-index: 10;
    width: 280px;
}

.posts-author {
    background-color: #fff;
    border: 1px solid #e3e8ec;
    border-radius: 12px;
    padding: 30px;
    display: flex;
    flex-wrap: wrap;
    text-align: center;
    margin-bottom: 10px;
}

.posts-author a {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    margin-bottom: 13px;
}

.posts-author img.avatar {
    width: 100px;
    height: 100px;
    margin-top: -80px;
    border-radius: 50%;
    padding: 3px;
    border-width: 1px;
    border-style: solid;
    border-color: rgb(232, 230, 230);
    border-image: initial;
    background: rgb(255, 255, 255);
}

a.posts-author-avatar {
}

a.posts-author-group span {
    border-radius: 5px;
    border: 1px solid;
    line-height: 14px;
    vertical-align: middle;
    font-size: 12px;
    font-weight: normal;
    color: #5896de;
    border-color: #5896de;
    padding: 2px 5px;
}

.posts-author-description {
    word-break: break-all;
    text-overflow: ellipsis;
    height: 50px;
    overflow: hidden;
}
a.posts-author-info span {
    margin: 0 5px;
    font-size: 12px;
    color: #999;
}
a.posts-author-nickname {
    font-size: 14px;
    font-weight: 600;
    color: #666;
    word-break: keep-all;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}
a.posts-author-medal img.medal {
    max-width: 100%;
    margin: 10px;
}
.posts-author a:hover,
.posts-author a span:hover {
    color: #ffbcd2;
}
</style>