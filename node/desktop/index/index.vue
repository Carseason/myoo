<template>
	<div class="body">
		<template v-for="( item , i) in Data">
			<div class="row-box" :key="i" ref="BoxScroll">
				<template v-if="item.Module == 0">
					<div class="row-box-carousel">
						<div class="row-box-carousel-f">
							<myoo-carousel :carousel="item.Posts"></myoo-carousel>
						</div>
						<div class="row-box-carousel-r" v-if="item.Widget.Enabled > 0">
							<template v-for="value in item.Widget.Posts">
								<li class="row-box-carousel-modules" :key="value.Id">
									<div class="modules-container">
										<router-link :to="$root.GetPostsUrl(value.Id)" class="modules-link">
											<img class="modules-thumbnail" v-lazy="value.Thumbnail" :key="value.Id" />
										</router-link>
										<h3 class="box_list-title">
											<router-link :to="$root.GetPostsUrl(value.Id)" class="box_list-title-link" :title="value.Title">{{value.Title}}</router-link>
										</h3>
									</div>
								</li>
							</template>
						</div>
					</div>
				</template>
				<template v-else>
					<div class="row-box_list" :class="{ left : item.Widget.Enabled > 0}">
						<div class="box_header">
							<h2 class="box_header-title">
								<a class="box_header-llink" :href="item.Url != '' ? item.Url :'javascript:void(0)'">
									<img :src="item.Icon" v-if=" item.Icon != '' " />
									<span :title="item.Title">{{item.Title}}</span>
								</a>
							</h2>
							<div class="placeholder"></div>
							<a class="box_header-more" :href="item.Url != '' ? item.Url :'javascript:void(0)'">
								<span>更多</span>
								<i class="iconfont icon-denglu2 icon-right"></i>
							</a>
						</div>
						<myoo-box :posts="item.Posts" :types=" item.Widget.Enabled > 0 ?'types' :'' " :module="item.Module"></myoo-box>
					</div>
					<div class="row-box_list right" v-if="item.Widget.Enabled  > 0">
						<myoo-widget :widget="item.Widget"></myoo-widget>
					</div>
				</template>
				<div v-if=" item.Code != ''" class="ad" v-html="item.Code"></div>
			</div>
		</template>
		<div v-if="Count > 0">
			<div class="box-elevator" :class="{on :TopPx}">
				<template v-for="(value,i) in Data">
					<li @click="Scrollshow(i)" :class="{on : scrollshow == i}" :key="i">{{value.Title}}</li>
				</template>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			Data: [],
			Count: 0,
			scrollshow: -1,
			TopPx: false
		};
	},
	methods: {
		Scrollshow: function (e) {
			this.$refs.BoxScroll[e].scrollIntoView({
				//滑动效果
				behavior: "smooth"
			});
		}
	},
	watch: {
		"$store.state.Title": {
			immediate: true,
			handler() {
				this.$store.commit("title");
			}
		},
		"$store.state.ScrollTop": function () {
			if (this.Count == 0) {
				return;
			}
			if (this.$store.state.ScrollStatus) {
				this.TopPx = true;
				for (let i = 0; i < this.Count; i++) {
					if (
						this.$refs.BoxScroll[i].offsetTop <=
						this.$store.state.ScrollTop
					) {
						this.scrollshow = i;
					} else {
						break;
					}
				}
			} else {
				this.TopPx = false;
				this.scrollshow = -1;
				return;
			}
		}
	},
	created: function () {
		this.$store.commit("load", true);
		this.$axios
			.get("/index")
			.then(res => {
				if (res.data.data != null) {
					this.Data = res.data.data;
					this.Count = this.Data.length;
				}
				this.$store.commit("load", false);
			})
			.catch(error => {
				console.log(error);
			});
	}
};
</script>
<style scoped>
.row-box {
	display: flex;
	flex-wrap: wrap;
	width: 100%;
	margin: 10px 0;
}

.row-box-carousel {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	display: -webkit-flex;
	display: flex;
	flex-wrap: wrap;
	margin-top: 10px;
	padding: 0 5px;
}

.row-box-carousel-f {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 35%;
	overflow: hidden;
	position: relative;
	height: 220px;
	border-radius: 4px;
}

.row-box-carousel-r {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 65%;
	display: flex;
	flex-wrap: wrap;
	margin: -10px 0;
}

li.row-box-carousel-modules {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: calc(25% - 10px);
	padding: 0 0 0 10px;
	margin: 10px 0 0 10px;
	height: 100px;
}

.modules-container {
	position: relative;
	width: 100%;
	height: 100%;
	border-radius: 4px;
}

a.modules-link {
	border-radius: 4px;
	display: block;
	height: 100%;
	overflow: hidden;
	position: relative;
}

img.modules-thumbnail {
	width: 100%;
	height: 100%;
}

h3.box_list-title {
	border-radius: 4px;
	position: absolute;
	bottom: 0;
	width: 100%;
	height: 30px;
	background: -webkit-linear-gradient(transparent, rgba(0, 0, 0, 0.5));
	padding: 5px 10px;
	margin: 0;
	font-size: 12px;
	font-weight: 400;
}

a.box_list-title-link {
	color: #fff;
	font-size: 13px;
	width: 100%;
	word-break: break-all;
	display: block;
	text-overflow: ellipsis;
	height: 100%;
	overflow: hidden;
	white-space: nowrap;
}

li.row-box-carousel-modules:hover h3.box_list-title {
	height: 100%;
}

.row-box_list {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
.row-box_list.left {
	max-width: 80%;
}
.row-box_list.right {
	max-width: 20%;
	padding: 0 5px 0 20px;
}

.box_header {
	-webkit-box-pack: start;
	justify-content: flex-start;
	-webkit-box-align: center;
	align-items: center;
	display: flex;
	padding: 0 5px;
}

h2.box_header-title {
	display: flex;
	font-weight: 700;
	margin: 0;
	position: relative;
}
h2.box_header-title a.box_header-llink {
	display: flex;
	height: 40px;
}
h2.box_header-title a.box_header-llink img {
	width: 30px;
	height: 30px;
	margin-right: 5px;
}

h2.box_header-title span {
	font-size: 24px;
	font-weight: 400;
	color: #222;
}
h2.box_header-title span:hover {
	color: #00a1d6;
}

.box-code img {
	max-width: 100%;
}
.box-elevator.on {
	top: 120px;
}
.box-elevator {
	position: fixed;
	transition: top 0.3s;
	background-color: #f6f9fa;
	border: 1px solid #e5e9ef;
	border-radius: 4px;
	min-width: 70px;
	top: 295px;
	margin-left: 10px;
}
.box-elevator li {
	padding: 0 5px;
	height: 30px;
	line-height: 30px;
	text-align: center;
	transition: background-color 0.3s, color 0.3s;
	cursor: pointer;
	-ms-user-select: none;
	user-select: none;
	text-overflow: ellipsis;
	word-wrap: break-word;
	word-break: break-all;
	overflow: hidden;
}

.box-elevator li.on,
.box-elevator li:hover {
	background-color: #00a1d6;
	color: #fff;
}

a.box_header-more {
	background-color: #fff;
	border: 1px solid #ccd0d7;
	color: #555;
	border-radius: 4px;
	text-align: center;
	padding: 0 7px;
	transition: all 0.2s;
	height: 22px;
	line-height: 22px;
	display: flex;
}

a.box_header-more i {
	font-size: 10px;
}

a.box_header-more:hover {
	background-color: #ccd0d7;
}
</style>


<style scoped>
@media (max-width: 600px) {
	.myoo-boxmodules {
		flex-wrap: nowrap !important;
		overflow-y: hidden;
	}
	.myoo-boxmodules::-webkit-scrollbar {
		display: none;
	}
	.row-box_list.left {
		max-width: 100%;
	}
	.row-box_list.right {
		max-width: 100%;
		padding: 0 5px;
		display: none;
	}
	.row-box-carousel-f {
		max-width: 100%;
		height: 150px;
	}
	.row-box-carousel-r {
		max-width: 100%;
		margin: 10px 0;
		flex-wrap: nowrap;
		overflow-y: hidden;
	}
	.row-box-carousel-r::-webkit-scrollbar {
		display: none;
	}
	.row-box-carousel {
		margin-top: 0;
	}
	li.row-box-carousel-modules {
		max-width: 50%;
		margin: 0 2.5px 0 0;
		padding: 0 2.5px 0 0;
	}
	.box-elevator {
		display: none;
	}
	.row-box {
		margin: 10px 0 0 0;
	}
}
</style>