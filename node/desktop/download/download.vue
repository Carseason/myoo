<template>
	<div class="body">
		<div class="filling"></div>
		<div class="download-content" v-if="Status">
			<template v-for="(value,i) in Download">
				<fieldset class="download-list" :key="i">
					<legend class="download-title">
						<span>{{value.Title}}</span>
					</legend>
					<div class="download-pwd" v-if=" value.DownloadPwd != '' ">
						<i class="iconfont icon-denglu2 icon-mima"></i>
						<span>提取密码:</span>
						<input :value="value.DownloadPwd" type="text" readonly="readonly" />
					</div>
					<div class="download-ext" v-if=" value.ExtractPwd != '' ">
						<i class="iconfont icon-denglu2 icon-mima"></i>
						<span>解压密码:</span>
						<input :value="value.ExtractPwd" type="text" readonly="readonly" />
					</div>
					<div class="download-url">
						<a href="javascript:void(0)" target="_blank" @click="OpenDownload(i)">
							<i class="iconfont icon-denglu2 icon-xiazai"></i>
							<span>下载</span>
						</a>
					</div>
				</fieldset>
			</template>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			Download: [],
			Status: false,
		}
	},
	watch: {
		'$store.state.Title'() {
			this.$store.commit("title", "下载");
		},
	},
	created() {
		this.$axios.get('/download/' + this.$route.params.id).then(res => {
			switch (res.data.success) {
				case 403:
					this.$router.push("/")
					break;
				case 302:
					this.$store.commit("msg", res.data.message)
					break;
				case 200:
					this.Status = true;
					this.Download = res.data.data ? res.data.data : [];
					break;
			}
		}).catch(error => {
			console.log(error);
		});
	},
	methods: {
		OpenDownload(e) {
			window.open(this.Download[e].Url, '_blank');
			return;
		}
	},
}
</script>
<style scoped>
.download-code {
	margin: 50px 0 20px 0;
}
.download-content {
	margin: 10px 0;
	width: 100%;
}
fieldset.download-list {
	background: rgba(0, 0, 0, 0.05);
	border: 1px solid rgba(0, 0, 0, 0.02);
	border-radius: 4px;
	margin: 10px;
}
legend.download-title span {
	background: #4caf50;
	border-radius: 1rem;
	color: #fff;
	display: inline-block;
	padding: 5px 10px;
	white-space: nowrap;
}
fieldset.download-list > div {
	display: flex;
	padding: 0 10px;
	margin: 10px 0;
	flex-wrap: wrap;
	height: 25px;
	line-height: 25px;
}
fieldset.download-list i {
	font-size: 1rem;
}
fieldset.download-list input {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	text-align: center;
	border: none;
	margin-left: 20px;
	border-radius: 5px;
	color: #616161;
}
fieldset.download-list input:hover {
	border: none;
}
fieldset.download-list .download-url a {
	text-align: center;
	color: #fff;
	display: block;
	width: 100%;
	background: #5cb860;
	border-radius: 5px;
}
fieldset.download-list .download-url a:hover {
	background: #00b5e5;
}
</style>