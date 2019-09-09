<template>
    <div class="comment-editor">
        <template v-if="$store.state.Auth.Id == 0">
            <textarea cols="45" rows="5" maxlength="200" placeholder="登录后才能发布评论......" required="required" disabled="disabled" ref="editarea"></textarea>
        </template>
        <template v-else>
            <textarea cols="45" rows="5" maxlength="200" required="required" placeholder="写评论......" v-model.trim="CommentsContent" ref="editarea"></textarea>
        </template>
        <div class="comment-box">
            <a class="emoji" @click="CommentsEmoji = true">
                <img src="https://i.loli.net/2019/09/06/73GX9JFIw6etHOK.png" />
            </a>
            <div class="emoji-content" v-if="CommentsEmoji">
                <div class="enter-header">
                    <i title="关闭" class="iconfont icon-denglu2 icon-close" @click="CommentsEmoji = false"></i>
                </div>
                <ul class="emoji-box">
                    <a title="大笑" @click="UpdateCommentsContent('[大笑]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/wbmvoD5ycFJ79KA.gif" />
                    </a>
                    <a title="喷血" @click="UpdateCommentsContent('[喷血]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/EbY8FCLmKcQIWHz.gif" />
                    </a>
                    <a title="抠鼻" @click="UpdateCommentsContent('[抠鼻]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/xfhvsjE284nPVip.gif" />
                    </a>
                    <a title="吐" @click="UpdateCommentsContent('[吐]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/98nzCEk4pHriBV1.gif" />
                    </a>
                    <a title="微笑" @click="UpdateCommentsContent('[微笑]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/YoNItSivPL2ZJAn.gif" />
                    </a>
                    <a title="笑哭" @click="UpdateCommentsContent('[笑哭]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/PX6M1eFxrYVZNdC.gif" />
                    </a>
                    <a title="斜眼笑" @click="UpdateCommentsContent('[斜眼笑]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/vcDieIRrf25KlAz.gif" />
                    </a>
                    <a title="阴险" @click="UpdateCommentsContent('[阴险]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/tNqeoXlZcywU41R.gif" />
                    </a>
                    <a title="doge" @click="UpdateCommentsContent('[doge]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/sRAJywb56qphYK1.gif" />
                    </a>
                    <a title="狗头" @click="UpdateCommentsContent('[狗头]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/Ww6GxbhD1Vamq8A.gif" />
                    </a>
                    <a title="猪头" @click="UpdateCommentsContent('[猪头]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/jGbMDPfFphlrXuz.gif" />
                    </a>
                    <a title="马" @click="UpdateCommentsContent('[马]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/MnKwdZi4eBVj3LD.gif" />
                    </a>
                    <a title="牛" @click="UpdateCommentsContent('[牛]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/AguMSPy1DTN7BVt.gif" />
                    </a>
                    <a title="啤酒" @click="UpdateCommentsContent('[啤酒]')" class="emoji-list">
                        <img src="https://i.loli.net/2019/08/05/q8i1MQUCE5YrOxF.gif" />
                    </a>
                </ul>
            </div>
            <span class="placeholder"></span>
            <button type="button" class="comment-box-primary" @click="SubmitComments()">回 复</button>
        </div>
    </div>
</template>
<script>
export default {
    props: ['id', 'parent'],
    data() {
        return {
            CommentsEmoji: false,
            CommentsContent: '',
        }
    },
    template: ``,
    mounted() {
    },
    methods: {
        UpdateCommentsContent(e) {
            var start = this.$refs.editarea.selectionStart;
            this.CommentsContent = this.CommentsContent.slice(0, start) + e + this.CommentsContent.slice(start, this.$refs.editarea.length)
            this.CommentsEmoji = false;
            this.$refs.editarea.focus()
            return;
        },
        SubmitComments() {  //提交评论
            if (this.CommentsContent == "") {
                return;
            }
            this.$axios.post('/comments/' + this.id, {
                content: this.CommentsContent,  //内容
            }).then(res => {
                if (res.data.success == 200) {
                    this.CommentsContent = "";
                }
                this.$store.commit("msg", res.data.message);
            }).catch(error => {
                console.log(error);
            });
        }
    },
}
</script>
<style scoped>
.comment-editor {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    margin-bottom: 10px;
}

.comment-editor textarea {
    font-size: 14px;
    width: 100%;
    border: 1px solid #d9d9d9;
    text-align: left;
    transition: all 0.3s linear;
    color: #555;
    background-color: #fff;
    border-radius: 3px;
    padding: 10px;
    outline: none;
    font-family: -webkit-body;
}

.comment-box {
    position: relative;
    display: flex;
    flex-wrap: wrap;
    padding: 5px;
}

button.comment-box-primary {
    background: #00a1d6;
    color: #fff;
    padding: 4px 18px;
    border-radius: 4px;
    transition: all 0.3s;
    -ms-user-select: none;
    border: none;
}

a.emoji {
    font-size: 15px;
}

.emoji-content {
    top: 30px;
    display: block;
    position: absolute;
    z-index: 99;
    width: 500px;
    border-radius: 3px;
    border: 1px solid #d9d9d9;
    background: #fff;
}

.enter-header {
    text-align: right;
    font-size: 1.5rem;
}

ul.emoji-box {
    display: flex;
    flex-wrap: wrap;
    padding: 5px;
}

ul.emoji-box a {
    -webkit-box-flex: 0;
    flex: 0 0 7.1%;
    max-width: 7.1%;
    font-size: 16px;
    padding: 0 4px;
    line-height: 40px;
    border: 1px solid #d9d9d9;
}

a.emoji-list img {
    width: 100%;
}

.enter-header i {
    background: rgba(0, 0, 0, 0.03);
    border-top-right-radius: 1rem;
    border-bottom-left-radius: 1rem;
    cursor: pointer;
    color: #666;
    font-size: 1.5rem;
    font-weight: 1000;
    padding: 5px 10px;
}
.enter-header i:hover {
    background: #fb72996b;
    color: #fff;
    border-top-right-radius: 3px;
}
</style>
