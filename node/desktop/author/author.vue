<template>
    <div class="body" v-if="Status">
        <div class="author-wrapper"></div>
        <div class="author-info">
            <div class="author-avatar">
                <img :src="Author.Avatar" class="avatar" />
            </div>
            <div class="author-config">
                <div class="author-user">
                    <span class="name">{{Author.Name}}</span>
                    <span class="capabilities">{{Author.Group.NickName}}</span>
                </div>
                <div class="author-data">注册于 {{Author.Date | dateForm2}}</div>
                <div class="author-description">{{Author.Description}}</div>
                <!-- <div class="author-medal">
                    <template v-for="(value ,i ) in Author.Group.Medal">
                        <img :src="value" class="medal" :key="i" />
                    </template>
                </div>-->
            </div>
            <div class="placeholder"></div>
            <div class="author-right">
                <ruby>
                    {{Author.Followers.Count}}
                    <rt>关注</rt>
                </ruby>
                <span class="right"></span>
                <ruby>
                    {{Author.Fans.Count}}
                    <rt>粉丝</rt>
                </ruby>
                <span class="right"></span>
                <ruby>
                    {{Author.Posts.Count}}
                    <rt>投稿</rt>
                </ruby>
                <span class="right"></span>
                <ruby>
                    {{Author.Comments.Count}}
                    <rt>评论</rt>
                </ruby>
            </div>

            <div class="author-followe" style="text-align: right; -webkit-box-flex: 0; flex: 0 0 100%;">
                <template v-if="$store.state.Auth.Id == 0 || $store.state.Auth.Id  == Author.Id">
                    <button class="gz isActive">+ 关注</button>
                </template>
                <template v-else>
                    <button class="gz" v-if="Author.Followers.Status" @click="UpdateFollowers()" title="取消关注">已关注</button>
                    <button class="gz" v-else @click="UpdateFollowers()" title="关注作者">+ 关注</button>
                </template>
            </div>
        </div>
        <div class="author-nav">
            <router-link :to="'/author/'+$route.params.id">
                <i class="iconfont icon-denglu2 icon-new-home"></i>
                <span>主页</span>
            </router-link>
            <router-link :to="'/author/'+$route.params.id +'/posts/1'">
                <i class="iconfont icon-denglu2 icon-neirong"></i>
                <span>作品</span>
            </router-link>
            <router-link :to="'/author/'+$route.params.id +'/followers/1'">
                <i class="iconfont iconfont icon-denglu2 icon-guanzhu"></i>
                <span>关注</span>
            </router-link>
            <router-link :to="'/author/'+$route.params.id +'/fans/1'">
                <i class="iconfont iconfont icon-denglu2 icon-fensi"></i>
                <span>粉丝</span>
            </router-link>
        </div>
        <div class="author-space">
            <router-view></router-view>
        </div>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Status: false,
            Author: {
                Id: 0,
                Name: "",
                Avatar: "",
                Description: "",
                Date: "",
                Lv: 0,
                Posts: {
                    Count: 0,
                },
                Fans: {
                    Count: 0,
                },
                Followers: {
                    Count: 0,
                    Status: false,
                },
                Comments: {
                    Count: 0,
                },
                Group: {
                    Lv: 0,
                    NickName: "",
                    //    Medal: [],
                }
            }
        }
    },
    watch: {
        '$store.state.Title': {
            immediate: true,
            handler() {
                this.$store.commit("title", this.Author.Name + "的空间");
            }
        },
        "$route.params.id"() {
            this.Get(this.$route.params.id)
        }
    },
    created() {
        this.Get(this.$route.params.id)
    },
    methods: {
        Get(val) {
            this.$store.commit("load", true)
            this.$axios.get("/author/" + val).then(res => {
                if (res.data.success != 200) {
                    this.$router.push("/")
                }
                this.$store.commit("load", false);
                if (res.data.data != null) {
                    this.Author = res.data.data;
                    this.$store.commit("title", this.Author.Name + "的空间")
                    this.Status = true;
                }
            }).catch(error => {
                console.log(error);
            });
        },
        UpdateFollowers() {
            this.$axios.post("/followers", {
                id: this.Author.Id
            }).then(res => {
                if (res.data.success != 200) {
                    this.$store.commit("msg", res.data.message)
                    return
                }
                this.Author.Followers.Status = !this.Author.Followers.Status;
                if (this.Author.Followers.Status) {                    //已关注
                    this.Author.Fans.Count += 1;
                } else {
                    this.Author.Fans.Count = this.Author.Fans.Count - 1;
                }
            }).catch(error => {
                console.log(error);
            });
        }
    }
}
</script>

<style scoped>
.author-wrapper {
    width: 100%;
    height: 400px;
    margin-top: -100px;
    background: #fbfbfb;
    background-size: 100%;
    background-repeat: no-repeat;
    background-image: url(/src/img/author.jpg);
}

.author-right {
    text-align: right;
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
}
.author-info {
    background: #fff;
    box-shadow: 0 0 0 1px #eee;
    border-radius: 0 0 4px 4px;
    padding: 20px;
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    position: relative;
}

.author-nav {
    height: 66px;
    line-height: 66px;
    text-align: center;
    background: #fff;
    box-shadow: 0 0 0 1px #eee;
    border-radius: 0 0 4px 4px;
    padding: 0 20px;
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    position: relative;
}

.author-space {
    margin-top: 20px;
    background: #fff;
    border: 1px solid #eee;
    border-radius: 4px;
    padding: 15px 20px;
    width: 100%;
    min-height: 500px;
}

.author-avatar {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 150px;
    margin-right: 10px;
}

.author-avatar img.avatar {
    width: 150px;
    height: 150px;
    padding: 5px;
    border: 1px solid #e8e6e6;
    cursor: pointer;
    object-fit: cover;
    border-radius: 50%;
}

.author-config > div {
    padding: 5px;
}

.author-description {
    border-top: 1px dashed #ddd;
}

span.right {
    /* border-right: 1px solid rgba(204, 204, 204, 0.71); */
    margin: 0px 5px;
}
span.name {
    font-size: 16px;
    font-weight: bold;
    color: #c66;
}

span.capabilities {
    margin: 0 8px;
    color: #fff;
    background: #ff6688;
    padding: 0 5px;
    border: 1px solid #ff6688;
    border-radius: 6px;
}
button.gz {
    background: #00a1d6;
    color: #fff;
    padding: 4px 18px;
    border-radius: 4px;
    transition: all 0.3s;
    -ms-user-select: none;
    border: none;
}

.author-nav a {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 6.5%;
    margin: 0 5px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
}

.author-nav i {
    margin-right: 4px;
    font-size: 20px;
    color: #00c091;
}
.author-nav a:nth-child(1) i {
    color: #00c091;
}
.author-nav a:nth-child(2) i {
    color: #02b5da;
}
.author-nav a:nth-child(3) i {
    color: #fb7299;
}
.author-nav a:nth-child(4) i {
    color: #fb7299;
}
.author-nav a {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 6.5%;
    margin: 0 5px;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
}

.author-nav i {
    margin-right: 4px;
    font-size: 20px;
    color: #00c091;
}

.author-nav span {
    font-size: 13px;
    font-weight: 400;
}
.author-nav a.router-link-exact-active.on,
.author-nav a:hover {
    transition: all 0.5s;
    border-bottom: 3px solid rgba(0, 161, 214, 0.7294117647058823);
}

.gz.isActive,
.toolbar-item.isActive {
    opacity: 0.5;
    cursor: not-allowed !important;
}

h3.author-title {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: -webkit-flex;
    display: flex;
    color: #00000087;
    font-weight: 400;
    padding: 5px;
    font-size: 20px;
    margin: 0;
}

.author-medal {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}
.author-medal img.medal {
    max-width: 100%;
    margin: 10px;
}
.author-right rt {
    font-size: 12px;
}
</style>
