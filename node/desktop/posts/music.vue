<template>
	<div id="player">
		<div class="music-theme">
			<div class="music-list-title">
				<span>歌曲列表（{{MusicList.length}}首）</span>
			</div>
			<div class="placeholder"></div>
		</div>
		<ul class="music-list" v-if="Status">
			<aplayer autoplay :music="MusicList[0]" :list="MusicList" />
		</ul>
	</div>
</template>
<script>
import Aplayer from "vue-aplayer";
Aplayer.disableVersionBadge = true;
export default {
	props: ["multimedia"],
	components: {
		Aplayer
	},
	data() {
		return {
			Path: 0,
			Status: false,
			MusicList: []
		}
	},
	mounted() {
		this.multimedia.map(item => {
			this.MusicList.push({
				title: item.Title,
				src: item.Url,
				artist: "",
				pic: "",
			})
		})
		this.Status = true;
	},
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
.music-theme {
	font-size: 18px;
	color: #222;
	line-height: 24px;
	margin-top: 18px;
	margin-bottom: 18px;
	width: 100%;
	display: flex;
}
ul.music-list {
	width: 100%;
	border-top: 1px solid #ccd0d7;
	overflow: auto;
}
</style>
