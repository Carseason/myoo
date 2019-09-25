<template>
	<div class="myoo-boxmodules" :class="types" :module="module">
		<template v-for="value in posts">
			<li class="row-box-modules" :key="value.Id">
				<div class="modules-container">
					<router-link :to="$root.GetPostsUrl(value.Id)" :class="'modules-link'">
						<img class="modules-thumbnail" v-lazy="value.Thumbnail" :key="value.Id" />
					</router-link>
					<h3 class="modules-widget">
						<router-link :to="$root.GetPostsUrl(value.Id)" class="modules-title" :title="value.Title">{{value.Title}}</router-link>
						<div class="modules-widget-put">
							<div class="modules-widget-author">
								<router-link :to="'/author/' + value.Author.Id" class="modules-widget-author-link">
									<div class="modules-widget-author-avatar">
										<img :src="value.Author.Avatar" class="avatar" />
									</div>
									<div class="modules-widget-author-nickname">{{value.Author.Name}}</div>
								</router-link>
							</div>
							<div class="modules-widget-time">{{value.Date | dateForm3}}</div>
						</div>
					</h3>
				</div>
			</li>
		</template>
	</div>
</template>
<script>
export default {
	props: ["posts", "types", "module"]
};
</script>
<style scoped>
.myoo-boxmodules {
	display: flex;
	flex-wrap: wrap;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
.myoo-boxmodules.types li.row-box-modules {
	max-width: 20%;
}
.myoo-boxmodules.mobile li.row-box-modules {
	max-width: 50%;
}

.myoo-boxmodules li.row-box-modules {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 16.6666%;
	margin: 10px 0;
}

.modules-container {
	padding: 0 5px;
}

a.modules-link {
	border-radius: 4px;
	display: block;
	height: 0;
	overflow: hidden;
	position: relative;
}
.myoo-boxmodules a.modules-link {
	padding-bottom: 50%;
}
.myoo-boxmodules[module="2"] a.modules-link {
	padding-bottom: 100%;
}

img.modules-thumbnail {
	width: 100%;
	height: 100%;
	position: absolute;
	top: 0;
	object-fit: cover;
}

a.modules-title {
	width: 100%;
	height: 100%;
	font-size: 12px;
	color: #9ba3ab;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
	display: inline-block;
}
a.modules-title:hover,
.modules-widget-author-nickname:hover {
	color: #00a1d6ba;
}

h3.modules-widget {
	height: 50px;
	padding: 5px;
	margin: 0;
	font-size: 12px;
	font-weight: 400 !important;
	position: relative;
	overflow: hidden;
}

.modules-widget-put {
	position: absolute;
	width: 100%;
	bottom: 0;
	left: 0;
	height: 25px;
	line-height: 25px;
	padding: 0 5px;
	color: #9ba3ab;
	background-color: #fff;
	transition: all 0.4s;
	display: flex;
	flex-wrap: wrap;
}

.modules-widget-author {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 70%;
}

.modules-widget-time {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 30%;
	text-align: right;
}

img.avatar {
	border-radius: 50%;
	width: 20px;
	height: 20px;
}

.modules-widget-author-avatar {
	margin-right: 0.5rem;
}

a.modules-widget-author-link {
	display: flex;
}

.modules-widget-time {
	font-size: 10px;
}

.modules-widget-author-nickname {
	color: #9ba3ab;
	overflow: hidden;
	text-overflow: ellipsis;
	white-space: nowrap;
}
</style>

<style scoped>
@media (max-width: 600px) {
	.myoo-boxmodules li.row-box-modules,
	.myoo-boxmodules.types li.row-box-modules {
		max-width: 50%;
	}
}
</style>
