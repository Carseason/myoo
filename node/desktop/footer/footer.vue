<template>
	<footer :class="$store.state.RouterPath">
		<div class="body">
			<!-- <div v-if="$store.state.Inc.Ad.Footer != '' " class="ad" v-html="$store.state.Inc.Ad.Footer"></div> -->
			<div class="footer-container-links">
				<div class="footer-container-links-header">
					<h2 class="footer-container-links-header-title">
						<i class="iconfont icon-denglu2 icon-link"></i>
						<span class="footer-container-links-header-title-text">友情链接</span>
					</h2>
				</div>
				<nav class="footer-container-links-body">
					<ul class="footer-container-links-body-menu">
						<template v-for="(value,i) in $store.state.Menus.Links">
							<li class="footer-container-links-body-menu-text" :key="i">
								<a target="_blank" :href="value.Url" class="footer-container-links-body-menu-is-text">
									<span class="footer-container-links-body-menu-is-text">{{value.Title}}</span>
								</a>
							</li>
						</template>
					</ul>
				</nav>
			</div>
		</div>
		<div class="myoo-footer-back" :class="{ on : $store.state.ScrollStatus}">
			<div title="公告" @click="$store.commit('msg' ,$store.state.Inc.Message)">
				<i class="iconfont icon-denglu2 icon-gonggao"></i>
			</div>
			<div title="返回顶部" @click="BackTop()">
				<i class="iconfont icon-denglu2 icon-xiangxia"></i>
			</div>
		</div>
		<message v-if="$store.state.Message.Status"></message>
		<loading v-if="$store.state.Loading"></loading>
	</footer>
</template>
<script>
export default {
	methods: {
		BackTop: function () {
			//返回顶部
			document.documentElement.scrollTop = 0;
		}
	}
};
</script>
<style scoped>
footer {
	margin-top: 100px;
	padding-bottom: 100px;
}
h2.footer-container-links-header-title {
	font-size: 1rem;
	margin: 0;
}
nav.footer-container-links-body {
	-webkit-box-shadow: none;
	border: 0;
	box-shadow: none;
	padding: 0;
}
ul.footer-container-links-body-menu {
	margin: 10px 0;
}
li.footer-container-links-body-menu-text {
	margin: 0 0.5rem 0.5rem 0;
	display: inline-block;
	list-style-type: none;
}
a.footer-container-links-body-menu-is-text {
	display: inline-block;
	padding: 0.2rem 0.3rem;
	color: #9ba3ab;
}

.myoo-footer-back {
	position: fixed;
	z-index: 100;
	right: -100px;
	bottom: 10%;
	display: block;
	text-align: center;
	transition: 0.5s;
}
.myoo-footer-back.on {
	right: 30px;
}
.myoo-footer-back > div {
	width: 48px;
	height: 48px;
	line-height: 48px;
	cursor: pointer;
	background-color: #f6f9fa;
	border: 1px solid #e5e9efb3;
	border-radius: 4px;
	font-size: 2rem;
	margin: 10px 0;
}
.myoo-footer-back > div:hover {
	background-color: #00a1d6;
	color: #fff;
}
.footer-container-links {
	padding: 10px;
}
</style>
<style scoped>
@media (max-width: 600px) {
	footer {
		margin-top: 10px;
		padding-bottom: 10px;
	}
}
</style>