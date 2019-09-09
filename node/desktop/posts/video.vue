<template>
    <div class="posts-body">
        <div class="posts-row">
            <div class="posts-player">
                <div id="dplayer" ref="dplayer"></div>
            </div>
            <div class="player-module">
                <div class="player-module_title">
                    <i class="iconfont icon-denglu2 icon-liebiao"></i>
                    <h4>选集</h4>
                </div>
                <ul class="player-list">
                    <template v-for="(value,i) in this.posts.Multimedia">
                        <li :key="i" @click="UpdatePlayer(i)" :class="{ on : i == Path }">{{value.Title}}</li>
                    </template>
                </ul>
            </div>
        </div>

        <div class="posts-row">
            <div class="posts-wrap">
                <div class="posts-row">
                    <div class="placeholder"></div>
                    <myoo-features :posts="posts"></myoo-features>
                </div>
                <div class="posts-row">
                    <div class="posts-thumbnail">
                        <img :src="posts.Thumbnail" alt class="thumbnail" />
                    </div>
                    <div class="posts-modules">
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
                <div class="posts-row">
                    <div class="posts-comments">
                        <myoo-comments :id="posts.Id"></myoo-comments>
                    </div>
                </div>
            </div>
            <div class="posts-widget">
                <myoo-author :author="author"></myoo-author>
            </div>
        </div>
    </div>
</template>
<script>
import hls from 'hls.js';
window.Hls = hls;
import DPlayer from 'dplayer';
import 'dplayer/dist/DPlayer.min.css';
import author from "./author.vue"
import features from "./features.vue";

export default {
    components: {
        "myoo-author": author,
        "myoo-features": features,
    },
    props: ["posts", "author"],
    data() {
        return {
            dp: {},
            Path: 0,
        }
    },
    mounted() {
        this.dp = new DPlayer({
            container: this.$refs.dplayer,
            autoplay: false,
            video: {
                url: this.posts.Multimedia[0].Url,
            },
        });
    },
    methods: {
        UpdatePlayer(e) {
            this.Path = e;
            this.dp.switchVideo({
                url: this.posts.Multimedia[e].Url,
            });
            return;
        },
    }
}
</script>
<style scoped>
.posts-body {
    display: flex;
    flex-wrap: wrap;
    width: 100%;
    margin: 10px 0;
}
/* .posts-bg {
    position: fixed;
    top: 0;
    left: 0;
    z-index: -1;
    width: 100%;
    height: 100%;
    background-color: rgb(240, 240, 240);
} */
.posts-row {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: flex;
    flex-wrap: wrap;
}

.posts-player {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(80% - 30px);
    margin-right: 30px;
}

.player-module {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 20%;
    padding: 10px;
    position: relative;
    background-color: #f4f4f4;
    border-radius: 4px;
    height: 100%;
}
.player-module_title {
    font-size: 16px;
}

.player-module_title h4 {
    display: inline-block;
    color: #212121;
    font-weight: 400;
    margin: 0 0 0 5px;
    padding: 0;
}
.posts-comments {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
}

ul.player-list {
    position: relative;
    display: block;
    overflow: auto;
    text-overflow: ellipsis;
    padding-top: 10px;
    padding: 10px 0;
    max-height: 0;
    min-height: 95%;
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}
ul.player-list li {
    display: block;
    height: 30px;
    line-height: 30px;
    font-size: 15px;
    overflow: hidden;
    -o-text-overflow: ellipsis;
    text-overflow: ellipsis;
    white-space: nowrap;
    cursor: pointer;
    padding: 0 10px;
    margin: 5px 0;
    border-radius: 4px;
}
ul.player-list li.on,
ul.player-list li:hover {
    background-color: #fff;
    color: #00a1d6;
}

h1.posts-title {
    font-size: 2rem;
    color: #676262;
    font-weight: 500;
    padding: 0;
    margin: 0;
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
}

.posts-sidebar {
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
}
.posts-context {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    margin: 20px 0;
    padding: 10px 0;
    border-top: 1px solid #eee;
    padding-bottom: 0.1px;
    display: flex;
    flex-wrap: wrap;
}

.posts-wrap {
    padding: 10px 0;
    margin: 10px 0;
    border-top: 1px solid #e7e7e7;
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(80% - 30px);
    margin-right: 30px;
    display: flex;
    flex-wrap: wrap;
}

.posts-thumbnail {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 180px;
    height: 180px;
    border-radius: 2px;
    overflow: hidden;
    background-color: #f4f5f7;
    margin-bottom: 20px;
}

.posts-thumbnail img {
    max-width: 100%;
    height: 100%;
}

.posts-modules {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(100% - 240px);
    margin: 0 30px;
}
</style>