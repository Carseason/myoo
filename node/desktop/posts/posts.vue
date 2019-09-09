<template>
    <div class="body" :class="Posts.Type">
        <template v-if="Status">
            <myoo-video v-if="Posts.Type == 'video'" :posts="Posts" :author="Author"></myoo-video>
            <myoo-music v-else-if="Posts.Type == 'music'" :posts="Posts" :author="Author"></myoo-music>
            <myoo-image v-else-if="Posts.Type == 'image'" :posts="Posts" :author="Author"></myoo-image>
            <myoo-standard v-else :posts="Posts" :author="Author"></myoo-standard>
        </template>
        <template></template>
    </div>
</template>
<script>
import Player from './video.vue';
import Music from './music.vue';
import Image from './image.vue';
import Standard from './standard.vue';
export default {
    components: {
        "myoo-video": Player,
        "myoo-music": Music,
        "myoo-image": Image,
        'myoo-standard': Standard,
    },
    data() {
        return {
            Status: false,
            Title: "",
            Posts: {},
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
                    Medal: [],
                }
            },
        }
    },
    mounted() {
    },
    watch: {
        '$store.state.Title': {
            immediate: true,
            handler() {
                this.$store.commit("title", this.Title)
            }
        },
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get('/posts/' + this.$route.params.id).then(res => {
            if (res.data.success !== 200) {
                window.location.href = "/"
            }
            this.Posts = res.data.data;
            this.Title = this.Posts.Title;
            this.GetAuthor(this.Posts.Author)
            this.$store.commit("title", this.Title)
        }).catch(error => {
            console.log(error);
        });
    },
    methods: {
        GetAuthor(id) {
            this.$axios.get("/author/" + id).then(res => {
                if (res.data.data != null) {
                    this.Author = res.data.data;
                }
                this.$store.commit("load", false)
                this.Status = true;
            }).catch(error => {
                console.log(error);
            });
        },



        /**************关注**********/
        UpdateFollowers() {
            if (this.$store.state.Auth.Id == 0) {
                return
            }
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
        },

        /**************点赞收藏**********/
        UpdateFeatures(e) {
            if (this.$store.state.Auth.Id == 0) {
                this.$store.commit("msg", "请先登录")
                return
            }
            this.$axios.post('/posts/' + this.$route.params.id, {
                type: e
            }).then(res => {
                if (res.data.success == 200) {
                    switch (e) {
                        case "favorite":
                            this.Posts.Features.Favorite = !this.Posts.Features.Favorite;
                            break;
                        case "supported":
                            this.Posts.Features.Supported.Status = !this.Posts.Features.Supported.Status;
                            if (this.Posts.Features.Supported.Status) {
                                this.Posts.Features.Supported.Count = this.Posts.Features.Supported.Count + 1;
                            } else {
                                this.Posts.Features.Supported.Count = this.Posts.Features.Supported.Count - 1;
                            }
                            break;
                    }
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
</style>
