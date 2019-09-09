<template>
    <div class="modules-widget">
        <div class="widget-title">
            <template v-if=" widget.Title != ''">
                <i class="iconfont icon-denglu2 icon-paihangbang"></i>
                <span>{{ widget.Title }}</span>
            </template>
        </div>
        <template v-if="widget.Enabled == 1">
            <div class="widget-leaderboard">
                <template v-for="(value,i) in widget.Posts">
                    <li class="leaderboard_list" :key="value.Id">
                        <i class="leaderboard_list_num">{{ i + 1 }}</i>
                        <router-link :to="$root.GetPostsUrl(value.Id)" class="leaderboard_list_links">
                            <div class="leaderboard_list_preview">
                                <img class="thumbnail" v-lazy="value.Thumbnail" :key="value.Id" />
                            </div>
                            <div class="leaderboard_list_conter">
                                <p :title="value.Title">{{value.Title}}</p>
                            </div>
                        </router-link>
                    </li>
                </template>
            </div>
        </template>
        <template v-if="widget.Enabled == 2">
            <div class="widget-carousel">
                <myoo-carousel :carousel="widget.Posts"></myoo-carousel>
            </div>
        </template>
        <template v-if="widget.Enabled == 3">
            <div class="widget-code" v-html="widget.Code"></div>
        </template>
    </div>
</template>
<script>
export default {
    props: ["widget"],
}
</script>
<style scoped>
.widget-title {
    margin: 0;
    padding: 0;
    font-size: 17px;
    height: 40px;
    line-height: 40px;
    margin-bottom: 10px;
    color: #615f5f;
}
.widget-title i {
    color: #8c8c8c;
}
.widget-title span {
    color: #615f5f;
    font-size: 13px;
}
li.leaderboard_list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    margin-bottom: 10px;
    display: flex;
}

i.leaderboard_list_num {
    color: #fff;
    height: 18px;
    line-height: 18px;
    font-size: 12px;
    text-align: center;
    border-radius: 4px;
    padding: 0 5px;
    font-weight: bolder;
    font-style: normal;
    margin-right: 5px;
    background-color: #ffafc9;
}

a.leaderboard_list_links {
    display: -webkit-flex;
    display: flex;
}

.leaderboard_list_preview {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 80px;
    height: 50px;
    border-radius: 4px;
    overflow: hidden;
    display: none;
}

li.leaderboard_list:nth-child(1) .leaderboard_list_preview {
    display: inline-block;
}

img.thumbnail {
    width: 100%;
    height: 100%;
}

.leaderboard_list_conter {
    margin-left: 5px;
    height: 24px;
}

.leaderboard_list_conter p {
    -o-text-overflow: ellipsis;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
    word-break: break-all;
    line-height: 18px;
    height: 18px;
    overflow: hidden;
    color: #222;
    font-size: 1.1rem;
    font-weight: 400;
}

li.leaderboard_list:nth-child(1) .leaderboard_list_conter p {
    -webkit-line-clamp: 2;
    height: 36px;
}
.leaderboard_list_conter p:hover {
    color: #005ca4 !important;
}
.modules-widget {
    width: 100%;
    height: 100%;
}
.widget-carousel {
    width: 100%;
    height: 100%;
    max-height: calc(100% - 100px);
    min-height: 130px;
    position: relative;
    overflow: hidden;
    border-radius: 4px;
    margin-bottom: -100px;
}
.widget-code img {
    max-width: 100%;
}
</style>
