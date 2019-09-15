<template>
	<div id="myoo-auth">
		<template v-if="$store.state.Auth.Id > 0 ">
			<router-link to="/account" class="auth-links">
				<img :src="$store.state.Auth.Avatar" alt="加载中" class="avatar" />
			</router-link>
			<ul class="auth-nav">
				<div class="auth-avatar">
					<img :src="$store.state.Auth.Avatar" alt="加载中" class="avatar" />
				</div>
				<div class="auth-info">
					<span class="auth-nickname">{{$store.state.Auth.Name}}</span>
					<span class="auth-description">{{$store.state.Auth.Description}}</span>
				</div>
				<div class="auth-group">
					<li>
						<i class="iconfont icon-denglu2 icon-dengji"></i>
						{{ $store.state.Auth.Group.NickName}}
					</li>
					<li>
						<i class="iconfont icon-denglu2 icon-xunzhang"></i>
						{{ $store.state.Auth.Group.Lv}}
					</li>
				</div>
				<div class="auth-menus">
					<router-link :to="'/author/' + $store.state.Auth.Id">
						<i class="iconfont icon-denglu2 icon-shouye"></i>
						我的空间
					</router-link>
					<router-link to="/account">
						<i class="iconfont icon-denglu2 icon-shezhi"></i>
						个人中心
					</router-link>
				</div>
				<div class="auth-footer">
					<router-link to="/admin" style="text-align: left;">
						<i class="iconfont icon-denglu2 icon-anquan"></i>
						管理
					</router-link>
					<a style="text-align: right;" @click="OutLogin()">
						<i class="iconfont icon-denglu2 icon-weizhi"></i>
						退出
					</a>
				</div>
			</ul>
		</template>
		<template v-else>
			<router-link to="/sign/login" class="auth-login" title="登录">
				<img src="/src/avatar/login.jpg" class="avatar" />
			</router-link>
		</template>
	</div>
</template>
<script>
export default {
	methods: {
		OutLogin() {
			this.$store.commit("localstorage")
			location.reload();
		},
	}
}
</script>
<style scoped>
a.auth-login {
	display: block;
	width: 32px;
	height: 32px;
	position: relative;
}

a.auth-login img.avatar {
	border: 0 solid #fff;
	width: 100%;
	height: 100%;
	border-radius: 50%;
}

#myoo-auth {
	position: relative;
	padding: 0 7px;
	margin: 0 5px;
}

a.auth-links {
	display: block;
	width: 32px;
	height: 32px;
	position: relative;
}

a.auth-links img.avatar {
	border: none;
	width: 100%;
	height: 100%;
	border-radius: 50%;
}

ul.auth-nav {
	left: -100px;
	top: 35px;
	display: -webkit-flex;
	display: flex;
	flex-wrap: wrap;
	position: absolute;
	opacity: 0;
	visibility: hidden;
	width: 260px;
	margin: 0 auto;
	background: rgba(255, 255, 255, 0.95);
	box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
	border-radius: 4px;
	text-align: center;
}

#myoo-auth:hover ul.auth-nav {
	opacity: 1;
	visibility: visible;
	transition: all 0.5s;
}

ul.auth-nav > div {
	max-width: 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	padding: 0 10px;
}
.auth-avatar,
.auth-info {
	margin: 10px 0;
}
.auth-avatar img.avatar {
	width: 64px;
	height: 64px;
	border-radius: 50%;
	border: 2px solid #ffffff;
}

.auth-info span {
	color: #6f6f6f;
	text-align: center;
	width: 100%;
	height: 28px;
	line-height: 28px;
	display: block;
	white-space: nowrap;
	overflow: hidden;
	text-overflow: ellipsis;
	margin: 2px 0;
}
span.auth-nickname {
	font-size: 15px;
}
span.auth-description {
	font-size: 12px;
}
.auth-menus {
	border-top: 1px solid #e5e9ef;
	padding: 10px;
	display: flex;
	flex-wrap: wrap;
}

.auth-menus a {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50%;
	line-height: 35px;
	cursor: pointer;
	font-size: 12px;
	color: #666;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}

.auth-footer {
	width: 100%;
	height: 30px;
	line-height: 30px;
	background-color: #f4f5f7;
	border-radius: 0 0 4px 4px;
	display: flex;
	flex-wrap: wrap;
}

.auth-footer a {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50%;
	cursor: pointer;
	font-size: 12px;
	color: #666;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
ul.auth-nav a:hover {
	color: #fb7299;
}

.auth-group {
	display: flex;
	flex-wrap: wrap;
	height: 30px;
	line-height: 30px;
	text-align: left;
}

.auth-group li {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50%;
	color: #615f5f;
}
</style>
<style  scoped>
@media (max-width: 600px) {
	#myoo-auth:hover ul.auth-nav {
		display: none;
	}
}
</style>