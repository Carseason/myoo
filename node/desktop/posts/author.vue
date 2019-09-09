<template>
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
                <button class="gz" v-if="author.Followers.Status" @click="UpdateFollowers()" title="取消关注">已关注</button>
                <button class="gz" v-else @click="UpdateFollowers()" title="关注作者">+ 关注</button>
            </template>
        </a>
        <!-- <a class="posts-author-medal">
            <template v-for="(value ,i ) in author.Group.Medal">
                <img :src="value" class="medal" :key="i" />
            </template>
        </a>-->
    </div>
</template>
<script>
export default {
    props: ["author"],
    methods: {
        /**************关注**********/
        UpdateFollowers() {
            if (this.$store.state.Auth.Id == 0) {
                return
            }
            this.$axios.post("/followers", {
                id: this.author.Id
            }).then(res => {
                if (res.data.success != 200) {
                    this.$store.commit("msg", res.data.message)
                    return
                }
                this.author.Followers.Status = !this.author.Followers.Status;
                if (this.author.Followers.Status) {                    //已关注
                    this.author.Fans.Count += 1;
                } else {
                    this.author.Fans.Count = this.author.Fans.Count - 1;
                }
            }).catch(error => {
                console.log(error);
            });
        },
    }
}
</script>
<style scoped>
.posts-author {
    border-radius: 12px;
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
    margin-bottom: 10px;
}

.posts-author img.avatar {
    width: 100px;
    height: 100px;
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
