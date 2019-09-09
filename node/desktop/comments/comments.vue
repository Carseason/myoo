<template>
    <div id="comments">
        <div class="comment-theme">
            <span>评论（{{MaxCount | count}}）</span>
        </div>
        <div class="comment-edit">
            <commentsTemplate :id="id" :parent="0"></commentsTemplate>
        </div>
        <div class="comment-list">
            <template v-for="value in Comments">
                <li class="comment-list__container" :key="value.Id">
                    <div class="parent-comment">
                        <div class="comment-face">
                            <router-link :to="'/author/'+ value.Author.Id">
                                <img :src="value.Author.Avatar" class="avatar" />
                            </router-link>
                        </div>
                        <div class="comment-body">
                            <div class="comments-nickname">
                                <router-link :to="'/author/'+ value.Author.Id" class="comment-info">{{ value.Author.Name}}</router-link>
                            </div>
                            <div class="comment-tire">
                                <div class="comment-data">{{value.Date | dateForm3 }}</div>
                            </div>
                            <div class="comment-content" v-html="value.Content"></div>
                        </div>
                    </div>
                </li>
            </template>
        </div>
        <navpage :page.sync="Page" :maxpage="MaxPage" :status="0"></navpage>
    </div>
</template>
<script>
import CommentsTemplate from './comments_edit.vue';
export default {
    props: ["id"],
    components: {
        'commentsTemplate': CommentsTemplate,
    },
    data() {
        return {
            Page: 1,
            MaxCount: 0,
            MaxPage: 0,
            Comments: [],
        }
    },
    watch: {
        Page(val) {
            this.GetComments(val)
        }
    },
    created() {
        this.$axios.get("/comments/" + this.id + "/0").then(res => {
            this.MaxPage = res.data.data.MaxPage ? res.data.data.MaxPage : 0;
            this.MaxCount = res.data.data.MaxCount ? res.data.data.MaxCount : 0;
            if (this.MaxCount > 0) {
                this.GetComments(1)
            }
        }).catch(error => {
            console.log(error);
        });
    },
    methods: {
        GetComments(val) {
            this.$axios.get("/comments/" + this.id + "/" + val).then(res => {
                this.Comments = res.data.data ? res.data.data : [];
            }).catch(error => {
                console.log(error);
            });
        }
    }
}
</script>
<style scoped>
div#comments {
    width: 100%;
    padding: 10px 0;
    display: -webkit-flex;
    display: flex;
    flex-wrap: wrap;
    border-top: 1px solid rgb(238, 238, 238);
}

.comment-theme {
    font-size: 18px;
    color: #222;
    line-height: 24px;
    margin: 10px 0;
    width: 100%;
    display: flex;
}

.comment-edit {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}

.comment-list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}

li.comment-list__container {
    width: 100%;
    padding: 10px 0;
    display: flex;
    flex-wrap: wrap;
    border-top: 1px dashed #ddd;
}

.parent-comment {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    display: flex;
    flex-wrap: wrap;
}

.comment-face {
    position: relative;
    margin-right: 10px;
}

.comment-face img.avatar {
    width: 75px;
    height: 75px;
    padding: 3px;
    border: 1px solid #e8e6e6;
    background: #fff;
    border-radius: 50%;
}

.comment-body {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
}
a.comment-info {
    color: #333;
    font-size: 15px;
    line-height: 15px;
    padding: 5px 0;
    margin-bottom: 5px;
}

.comment-tire {
    display: flex;
    flex-wrap: wrap;
}

.comment-data {
    font-size: 12px;
    color: #888;
}

.comment-content {
    margin-top: 5px;
    font-size: 14px;
}
</style>
