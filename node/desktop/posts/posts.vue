<template>
	<div class="body">
		<div class="filling"></div>
		<template v-if="Status">
			<div class="posts-thumbnail">
				<img :src="Posts.Thumbnail" alt class="thumbnail" />
			</div>
			<div class="posts-row">
				<h1 class="posts-title">{{Posts.Title}}</h1>
				<div class="posts-header">
					<div class="posts-category">
						<router-link :to="'/category/' + Posts.Category.Id">{{Posts.Category.Name}}</router-link>
					</div>
					<div class="placeholder"></div>
					<div class="posts-date">{{Posts.Date | dateForm3 }}</div>
					<div class="posts-info">
						<span>浏览：{{Posts.Views | count}}</span>
					</div>
				</div>
			</div>
			<template v-if="Posts.Type == 'video'">
				<player :multimedia="Posts.Multimedia"></player>
			</template>
			<template v-else-if="Posts.Type == 'music'">
				<music :multimedia="Posts.Multimedia"></music>
			</template>
			<div class="posts-row">
				<div class="posts-row-f">
					<div class="posts-content" v-html="Posts.Content"></div>
					<div class="posts-features">
						<template v-if="Posts.Supported.Status">
							<a class="features-star on" title="取消赞">
								<i class="iconfont icon-denglu2 icon-appreciatefill" @click="UpdateSupported()"></i>
								<span>{{Posts.Supported.Count | count }}</span>
							</a>
						</template>
						<template v-else>
							<a class="features-star" title="点赞">
								<i class="iconfont icon-denglu2 icon-appreciatefill" @click="UpdateSupported()"></i>
								<span>{{Posts.Supported.Count | count }}</span>
							</a>
						</template>

						<template v-if="Posts.Favorite.Status">
							<a class="features-favorite on" title="取消收藏">
								<i class="iconfont icon-denglu2 icon-favorfill" @click="UpdateFavorite()"></i>
							</a>
						</template>
						<template v-else>
							<a class="features-favorite">
								<i class="iconfont icon-denglu2 icon-favorfill" title="收藏" @click="UpdateFavorite()"></i>
							</a>
						</template>
						<router-link :to="'/download/' + Posts.Download" v-if="Posts.Download != '' ">
							<i class="iconfont icon-denglu2 icon-xiazai" title="下载"></i>
						</router-link>
					</div>
					<comments :id="Posts.Id"></comments>
				</div>
				<div class="posts-row-r">
					<author :id="Posts.Author"></author>
					<recommend :id="Posts.Category.Id"></recommend>
				</div>
			</div>
			<div class="posts-row">
				<div class="posts-row-f"></div>
			</div>
		</template>
		<template v-else>
			<load></load>
		</template>
	</div>
</template>
<script>
import Player from "./player.vue"
import Music from "./music.vue"
import Author from "./author.vue"
import Recommend from "./recommend.vue"
import Comments from "@/desktop/comments/comments.vue"
export default {
	components: {
		"author": Author,
		"comments": Comments,
		"player": Player,
		"music": Music,
		"recommend": Recommend,
	},
	data() {
		return {
			Status: false,
			Title: "",
			Posts: {},
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
		this.$axios.get('/posts/' + this.$route.params.id).then(res => {
			if (res.data.success !== 200) {
				window.location.href = "/"
			}
			this.Posts = res.data.data;
			this.Title = this.Posts.Title;
			this.$store.commit("title", this.Title);
			this.Status = true;
		}).catch(error => {
			console.log(error);
		});
	},
	methods: {
		UpdateSupported() {
			if (this.$store.state.Auth.Id == 0) {
				this.$store.commit("msg", "请先登录")
				return
			}
			this.$axios.post('/posts/' + this.Posts.Id, {
				type: "supported",
			}).then(res => {
				if (res.data.success == 200) {
					this.Posts.Supported.Status = !this.Posts.Supported.Status;
					if (this.Posts.Supported.Status) {
						this.Posts.Supported.Count = this.Posts.Supported.Count + 1;
					} else {
						this.Posts.Supported.Count = this.Posts.Supported.Count - 1;
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
			this.$axios.post('/posts/' + this.Posts.Id, {
				type: "favorite",
			}).then(res => {
				if (res.data.success == 200) {
					this.Posts.Favorite.Status = !this.Posts.Favorite.Status;
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
.filling {
	height: 95px;
}
.posts-row {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	display: flex;
	flex-wrap: wrap;
}
.posts-row-f {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 70%;
}
.posts-row-r {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 30%;
	padding-left: 40px;
}
.posts-header {
	display: flex;
	flex-wrap: wrap;
	font-size: 12px;
	color: #999;
	padding-bottom: 10px;
	margin-bottom: 20px;
	border-bottom: 1px dashed #ddd;
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}

.posts-category a {
	font-size: 12px;
	color: #999;
	margin-right: 5px;
}

.posts-date {
	margin-right: 10px;
}
h1.posts-title {
	color: #262626;
	line-height: 1.4;
	font-size: 30px;
	font-weight: 700;
	margin-bottom: 1.5rem;
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
.posts-thumbnail {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
.posts-thumbnail img.thumbnail {
	max-width: 100%;
	height: 150px;
	width: 100%;
}

.posts-content {
	min-height: 150px;
	font-size: 16px;
	line-height: 1.8;
	color: #555;
}
.posts-content img {
	max-width: 100%;
}
.post-content p {
	color: #555;
}
</style>
<style scoped>
.posts-features {
	display: flex;
	flex-wrap: wrap;
	text-align: right;
	position: relative;
	float: right;
	margin: 25px 0 0 0;
}
.posts-features > a {
	margin-right: 8px;
	cursor: pointer;
	transition: all 0.3s;
	white-space: nowrap;
	color: #505050;
	position: relative;
	text-align: left;
}

.posts-features i {
	font-size: 28px;
	vertical-align: top;
	margin-right: 6px;
	width: 28px;
	height: 28px;
	color: #757575;
}
.posts-features > a.on i,
.posts-features > a:hover i {
	color: #00a1d6;
}
a.features-star span {
	font-size: 15px;
}
</style>
<style scoped>
@media (max-width: 600px) {
	.filling {
		height: 0;
	}
	.posts-row {
		padding: 0 10px;
	}
	.posts-row-f {
		max-width: 100%;
	}
	.posts-row-r {
		max-width: 100%;
		padding-left: 0;
	}
	.posts-content {
		min-height: 0;
	}
}
</style>