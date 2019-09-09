<template>
    <div class="posts-features">
        <template v-if="posts.Supported.Status">
            <a class="features-star on" title="取消赞">
                <i class="iconfont icon-denglu2 icon-appreciatefill" @click="UpdateSupported()"></i>
                <span>{{posts.Supported.Count | count }}</span>
            </a>
        </template>
        <template v-else>
            <a class="features-star" title="点赞">
                <i class="iconfont icon-denglu2 icon-appreciate" @click="UpdateSupported()"></i>
                <span>{{posts.Supported.Count | count }}</span>
            </a>
        </template>

        <template v-if="posts.Favorite.Status">
            <a class="features-favorite on" title="取消收藏">
                <i class="iconfont icon-denglu2 icon-favorfill" @click="UpdateFavorite()"></i>
            </a>
        </template>
        <template v-else>
            <a class="features-favorite">
                <i class="iconfont icon-denglu2 icon-favor" title="收藏" @click="UpdateFavorite()"></i>
            </a>
        </template>
        <router-link :to="'/download/' + posts.Download" v-if="posts.Download != '' ">
            <i class="iconfont icon-denglu2 icon-xiazai" title="下载"></i>
        </router-link>
    </div>
</template>
<script>
export default {
    props: ["posts"],
    methods: {
        UpdateSupported() {
            if (this.$store.state.Auth.Id == 0) {
                this.$store.commit("msg", "请先登录")
                return
            }
            this.$axios.post('/posts/' + this.posts.Id, {
                type: "supported",
            }).then(res => {
                if (res.data.success == 200) {
                    this.posts.Supported.Status = !this.posts.Supported.Status;
                    if (this.posts.Supported.Status) {
                        this.posts.Supported.Count = this.posts.Supported.Count + 1;
                    } else {
                        this.posts.Supported.Count = this.posts.Supported.Count - 1;
                    }
                } else {
                    this.$store.commit("msg", res.data.message)
                }
            }).catch(error => {
                console.log(error);
            });
        },
        UpdateFavorite() {
            if (this.$store.state.Auth.Id == 0) {
                this.$store.commit("msg", "请先登录")
                return
            }
            this.$axios.post('/posts/' + this.posts.Id, {
                type: "favorite",
            }).then(res => {
                if (res.data.success == 200) {
                    this.posts.Favorite.Status = !this.posts.Favorite.Status;
                } else {
                    this.$store.commit("msg", res.data.message)
                }
            }).catch(error => {
                console.log(error);
            });
        },
    }
}
</script>
<style scoped>
.posts-features {
    display: flex;
    flex-wrap: wrap;
}
.posts-features > a {
    position: relative;
    cursor: pointer;
    color: #7c7b7b;
    line-height: 1;
    margin-left: 15px;
}

.posts-features i {
    font-size: 2.5rem;
    color: #9a9999;
}
.posts-features > a.on i {
    color: #00a1d6;
}
.posts-features i:hover {
    color: #80d0eb;
}
a.features-star span {
    position: absolute;
    top: 0;
    font-size: 10px;
}
</style>
