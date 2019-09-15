<template>
	<div id="player">
		<div class="video-player">
			<div id="dplayer" ref="dplayer"></div>
		</div>
		<div class="player-module">
			<div class="player-select">
				<div class="player-module_title">
					<i class="iconfont icon-denglu2 icon-liebiao"></i>
					<h4>选集</h4>
				</div>
				<ul class="player-list">
					<template v-for="(value,i) in multimedia">
						<li :key="i" @click="SetPlayerr(i)" :class="{ on : i == SelectPlayer }">{{value.Title}}</li>
					</template>
				</ul>
			</div>
		</div>
	</div>
</template>
<script>
import hls from 'hls.js';
window.Hls = hls;
import DPlayer from 'dplayer';
import 'dplayer/dist/DPlayer.min.css';
export default {
	props: ["multimedia"],
	data() {
		return {
			SelectPlayer: 0,
			Dp: {},
		}
	},
	mounted() {
		this.NewPlayer();
	},
	methods: {
		NewPlayer() {
			this.Dp = new DPlayer({
				container: this.$refs.dplayer,
				autoplay: false,
				video: {
					url: this.multimedia[0].Url,
				},
			});
		},
		SetPlayer(e) {
			this.SelectPlayer = e;
			this.Dp.switchVideo({
				url: this.multimedia[SelectPlayer].Url,
			});
		},
	}
}
</script>
<style scoped>
#player {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	display: flex;
	flex-wrap: wrap;
	margin: 10px 0;
}
.video-player {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 70%;
}
.player-module {
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 30%;
	padding-left: 40px;
}
.player-select {
	padding: 10px;
	position: relative;
	background-color: #f4f4f4;
	border-radius: 4px;
	height: 100%;
}
.player-module_title {
	font-size: 16px;
}
.player-module_title h4 {
	display: inline-block;
	color: #212121;
	font-weight: 400;
	margin: 0 0 0 5px;
	padding: 0;
}
ul.player-list {
	position: relative;
	display: block;
	overflow: auto;
	text-overflow: ellipsis;
	padding-top: 10px;
	padding: 10px 0;
	max-height: 0;
	min-height: 95%;
	-ms-flex: 0 0 100%;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
ul.player-list li {
	display: block;
	height: 30px;
	line-height: 30px;
	font-size: 15px;
	overflow: hidden;
	-o-text-overflow: ellipsis;
	text-overflow: ellipsis;
	white-space: nowrap;
	cursor: pointer;
	padding: 0 10px;
	margin: 5px 0;
	border-radius: 4px;
}
ul.player-list li.on,
ul.player-list li:hover {
	background-color: #fff;
	color: #00a1d6;
}
</style>
<style scoped>
@media (max-width: 600px) {
	.video-player {
		max-width: 100%;
	}
	.player-module {
		max-width: 100%;
		padding-left: 0;
	}
	.player-select {
		background-color: rgba(255, 255, 255, 0);
		height: 100%;
	}
	ul.player-list {
		max-height: 100%;
		min-height: 100%;
		flex-wrap: nowrap;
		display: flex;
	}
	ul.player-list::-webkit-scrollbar {
		display: none;
	}
	ul.player-list li {
		-webkit-box-flex: 0;
		flex: 0 0 100%;
		max-width: 7.1rem;
		margin-right: 0.5rem;
		font-size: 0.8rem;
		border-radius: 0.4rem;
		border: 1px solid rgba(0, 0, 0, 0.1);
		text-align: center;
		height: 4rem;
		line-height: 4rem;
	}
	ul.player-list li.on {
		border: 1px solid #2d8cf0;
		color: #2d8cf0;
	}
}
</style>