<template>
	<ul class="wrapper-menus" :class="types">
		<template v-if="types == 'mobile'">
			<template v-for="(value,i) in menus">
				<li :key="i" class="child-list">
					<router-link :to="value.Url" class="menus-link">
						<i :class="value.Icon"></i>
						<span :title="value.Title">{{value.Title}}</span>
					</router-link>
					<ul class="child-menus" v-if="value.Child != null  && value.Child.length > 0">
						<template v-for="(item,j) in value.Child">
							<li class="wrapperv-menus-child" :key="j">
								<router-link :to="item.Url" class="wrapperv-menus-child-link">
									<span class="menus-child-text">{{item.Title}}</span>
								</router-link>
							</li>
						</template>
					</ul>
				</li>
			</template>
		</template>
		<template v-else>
			<template v-for="(value,i) in menus">
				<li class="wrapper-item" :key="i">
					<router-link :to="value.Url" class="wrapper-link">
						<i v-if="value.Icon != ''" :class="value.Icon"></i>
						<span class="wrapper-item_title" :title="value.Title">{{value.Title}}</span>
					</router-link>
					<ul class="wrapperv-menus_sub" v-if="value.Child != null  && value.Child.length > 0">
						<template v-for="(item,j) in value.Child">
							<li class="wrapperv-menus-child" :key="j">
								<router-link :to="item.Url" class="wrapperv-menus-child-link">
									<i class="iconfont icon-denglu2 icon-wangyou-"></i>
									<span class="menus-child-text">{{item.Title}}</span>
									<i class="iconfont icon-denglu2 icon-wangzuo-"></i>
								</router-link>
							</li>
						</template>
					</ul>
				</li>
			</template>
		</template>
	</ul>
</template>
<script>
export default {
	props: ["menus", "types"]
}
</script>
<style scoped>
/************置顶*********/
ul.wrapper-menus.top {
	display: flex;
	flex-wrap: wrap;
}
ul.wrapper-menus.top a.wrapper-link i {
	color: #00a1d6;
	font-size: 20px;
	padding-right: 5px;
}
ul.wrapper-menus.top li.wrapper-item a {
	padding: 0 5px;
	cursor: pointer;
	display: block;
}

/************导航*********/
ul.wrapper-menus.nav {
	display: flex;
	flex-wrap: wrap;
	height: 50px;
	line-height: 50px;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}

ul.wrapper-menus.nav li.wrapper-item {
	position: relative;
	margin: auto 20px auto 0px;
}

ul.wrapper-menus.nav a.wrapper-link {
	display: block;
	margin-top: 10px;
}

ul.wrapper-menus.nav a.wrapper-link i {
	margin: 0 auto;
	display: block;
	background: #ffafc9;
	font-size: 12px;
	border-radius: 20%;
	text-shadow: 0 0px 0px rgba(0, 0, 0, 0);
	width: 22px;
	height: 15px;
	line-height: 15px;
	color: #fff;
	text-align: center;
}

ul.wrapper-menus.nav span.wrapper-item_title {
	font-size: 12px;
	display: block;
	height: 23px;
	padding: 0;
	margin: 0;
	text-align: center;
	line-height: 23px;
}
ul.wrapperv-menus_sub {
	background: #fff;
	margin: 0;
	min-width: 12rem;
	opacity: 0;
	padding: 10px 0;
	position: absolute;
	top: 100%;
	transform: scale3d(0.9, 0.7, 1);
	transform-origin: top left;
	visibility: hidden;
	z-index: 100;
	border-radius: 0 0 4px 4px;
	box-shadow: 0 15px 30px 5px rgba(0, 0, 0, 0.15);
}

li.wrapper-item:hover ul.wrapperv-menus_sub {
	opacity: 1;
	transform: scaleX(1);
	visibility: visible;
	transition: all 0.5s;
}
a.wrapperv-menus-child-link i {
	text-align: center;
	width: 10px;
	font-size: 10px;
	font-weight: 800;
	color: rgba(0, 0, 0, 0.33);
	padding: 0;
}

li.wrapperv-menus-child {
	line-height: 33.33333px;
	list-style-type: none;
	position: relative;
}

a.wrapperv-menus-child-link {
	font-size: 1rem;
	display: block;
	padding: 0 1rem;
	text-shadow: 0 0 2px #fff, 0 0 5px #fff, 0 0 10px #fff;
	white-space: nowrap;
}
a.wrapperv-menus-child-link i.iconfont.icon-denglu2.icon-wangzuo- {
	position: absolute;
	right: -50px;
	opacity: 0;
}
a.wrapperv-menus-child-link:hover i.iconfont.icon-denglu2.icon-wangzuo- {
	opacity: 1;
	transition: all 0.5s;
	position: relative;
	right: 0;
}
li.wrapper-item:hover span.wrapper-item_title {
	color: #00a1d6ba;
}
</style>
<style scoped>
@media (max-width: 600px) {
	a.menus-link {
		display: -webkit-box;
		display: -ms-flexbox;
		display: flex;
		white-space: nowrap;
		padding: 0 1rem;
		color: #22133d;
		-webkit-box-align: center;
		-ms-flex-align: center;
		align-items: center;
	}
	a.menus-link i {
		width: 2rem;
		height: 2rem;
		line-height: 2rem;
		margin-right: 1rem;
		border-radius: 50%;
		background: rgba(0, 0, 0, 0.2);
		color: #585858;
		text-align: center;
	}
	ul.child-menus {
		padding-left: 3rem;
		background: linear-gradient(rgba(0, 0, 0, 0.03), rgba(0, 0, 0, 0));
	}
	ul.wrapper-menus.mobile span {
		color: #2c2c2c;
	}
}
</style>