<template>
	<div class="posts-author">
		<template v-if="Status">
			<router-link :to="'/author/' + Author.Id" class="posts-author-avatar">
				<img :src="Author.Avatar" class="avatar" />
			</router-link>
			<router-link :to="'/author/' + Author.Id" class="posts-author-nickname" :title="Author.Name">{{Author.Name}}</router-link>
			<a class="posts-author-group">
				<span class="capabilities">{{Author.Group.NickName}}</span>
			</a>
			<a class="posts-author-description">{{Author.Description}}</a>
			<a class="posts-author-info">
				<span>投稿:{{Author.Posts.Count | count }}</span>
				<span class="block"></span>
				<span>评论:{{Author.Comments.Count | count}}</span>
				<span class="block"></span>
				<span>粉丝:{{Author.Fans.Count | count}}</span>
			</a>
			<a class="author-followe">
				<template v-if="$store.state.Auth.Id == 0 || $store.state.Auth.Id  == Author.Id">
					<button class="gz isActive">+ 关注</button>
				</template>
				<template v-else>
					<button class="gz" v-if="Author.Followers.Status" @click="UpdateFollowers()" title="取消关注">已关注</button>
					<button class="gz" v-else @click="UpdateFollowers()" title="关注作者">+ 关注</button>
				</template>
			</a>
			<!-- <a class="posts-author-medal">
            <template v-for="(value ,i ) in Author.Group.Medal">
                <img :src="value" class="medal" :key="i" />
            </template>
			</a>-->
		</template>
		<template v-else>
			<load></load>
		</template>
	</div>
</template>
<script>
export default {
	props: ["author", "id"],
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
					Medal: [],
				}
			},
		}
	},
	created() {
		this.$axios.get("/author/" + this.id).then(res => {
			if (res.data.data != null) {
				this.Author = res.data.data;
			}
			this.Status = true;
		}).catch(error => {
			console.log(error);
		});
	},
	methods: {
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
	}
}
</script>
<style scoped>
.posts-author {
	display: flex;
	flex-wrap: wrap;
	text-align: center;
	margin: 10px 0;
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
