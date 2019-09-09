<template>
    <div class="body">
        <div class="posts-sidebar">
            <div class="posts-row">
                <div class="posts-thumbnail">
                    <img :src="posts.Thumbnail" alt class="thumbnail" />
                </div>
                <div class="posts-modules">
                    <div class="posts-header">
                        <div class="posts-title">{{posts.Title}}</div>
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
                <div class="music-theme">
                    <div class="music-list-title">
                        <span>歌曲列表（{{posts.Multimedia.length}}首）</span>
                    </div>
                    <div class="placeholder"></div>
                    <myoo-features :posts="posts"></myoo-features>
                </div>
                <ul class="music-list" v-if="Status">
                    <aplayer autoplay :music="MusicList[0]" :list="MusicList" />
                </ul>
            </div>
            <div class="posts-row">
                <myoo-comments :id="posts.Id"></myoo-comments>
            </div>
        </div>
        <div class="posts-widget">
            <myoo-author :author="author"></myoo-author>
        </div>
    </div>
</template>
<script>
import Aplayer from "vue-aplayer";
Aplayer.disableVersionBadge = true;

import author from "./author.vue"
import features from "./features.vue";

export default {
    props: ["posts", "author"],
    components: {
        "myoo-author": author,
        "myoo-features": features,
        Aplayer
    },
    data() {
        return {
            Path: 0,
            Status: false,
            MusicList: []
        }
    },
    mounted() {
        this.posts.Multimedia.map(item => {
            this.MusicList.push({
                title: item.Title,
                src: item.Url,
                artist: "",
                pic: "",
            })
        })
        this.Status = true;
    },
    methods: {
        // UpdatePlayer(e) {
        //     this.Path = e;
        //     this.MusicList.src = this.posts.Multimedia[e].Url;
        //     this.MusicList.title = this.posts.Multimedia[e].Title;
        // }
    }
}
</script>
<style scoped>
.posts-row {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: flex;
    flex-wrap: wrap;
}
.posts-sidebar {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(80% - 30px);
    margin-right: 30px;
}
.posts-thumbnail {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 370px;
    height: 370px;
    border-radius: 2px;
    overflow: hidden;
}

.posts-thumbnail img {
    max-width: 100%;
    height: 100%;
}
.posts-modules {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: calc(100% - 430px);
    margin: 0 30px;
}
.posts-widget {
    -ms-flex: 0 0 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 20%;
}
.posts-title {
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

.music-theme {
    font-size: 18px;
    color: #222;
    line-height: 24px;
    margin-top: 18px;
    margin-bottom: 18px;
    width: 100%;
    display: flex;
}

ul.music-list {
    width: 100%;
    border-top: 1px solid #ccd0d7;
    overflow: auto;
}

ul.music-list li > span {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 30px;
    text-align: center;
}
ul.music-list li > div {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
    color: #222;
    overflow: hidden;
}

ul.music-list li:hover,
ul.music-list li.on {
    background: #efefef;
}
</style>
