<template>
	<form class="account-row" method="post" @submit.prevent="Update()">
		<ul class="account-newposts-header">攒写新文章</ul>
		<ul class="account-newposts-title">
			<input type="text" placeholder="标题" required="required" v-model.trim="Title" />
		</ul>
		<ul class="account-newposts-content">
			<quillEditor v-model.trim="Content" :options="editorOption"></quillEditor>
			<span class="count">{{Content.length}}</span>
		</ul>
		<ul class="account-newposts-category">
			<span class="title">投稿分类:</span>
			<select class="posts-category" v-model.number="Category">
				<option :value="0">选择投稿分类</option>
				<template v-for="value in CategoryAll">
					<option :value="value.Id" :key="value.Id">{{value.Name}}</option>
				</template>
			</select>
		</ul>
		<ul class="account-newposts-thumbnail">
			<span class="title">文章封面:</span>
			<input placeholder="封面特色图链接URL" type="url" v-model.trim="Thumbnail" />
			<div :style="'background-image: url(' +Thumbnail + ');' "></div>
		</ul>
		<ul class="account-newposts-types">
			<span class="title">文章类型:</span>
			<div class="account-newposts-flex">
				<button class="editor-type" type="button" :class="{ on :Types == 'standard'}" @click="Types = 'standard'">文章</button>
				<button class="editor-type" type="button" :class="{ on :Types == 'image'}" @click="Types = 'image'">图片</button>
				<button class="editor-type" type="button" :class="{ on :Types == 'music'}" @click="Types = 'music'">音乐</button>
				<button class="editor-type" type="button" :class="{ on :Types == 'video'}" @click="Types = 'video'">视频</button>
			</div>
			<div class="account-newposts-flex">
				<template v-if="Types == 'music'">
					<li v-for="(value,i) in Music" :key="i" class="account-newposts-list">
						<div class="account-newposts-list-c">
							<input placeholder="音乐名称" required="required" type="text" class="input f" v-model.trim="value.Title" />
							<input placeholder="音乐链接" required="required" type="url" class="input f" v-model.trim="value.Url" />
						</div>
						<div class="account-newposts-list-r">
							<i class="back" title="删除" @click="Music.splice(i,1)">-</i>
							<i class="add" title="增加" @click="Music.splice(i+1 ,0, {Title:'',Url:''})">+</i>
						</div>
					</li>
					<div class="account-newposts-icon">
						<i class="add" title="增加" @click="Music.push({Title:'',Url:''})">+</i>
					</div>
				</template>
				<template v-else-if="Types == 'video'">
					<li v-for="(value,i) in Video" :key="i" class="account-newposts-list">
						<div class="account-newposts-list-c">
							<input placeholder="视频名称" required="required" type="text" class="input c" v-model.trim="value.Title" />
							<input placeholder="视频链接" required="required" type="url" class="input c" v-model.trim="value.Url" />
						</div>
						<div class="account-newposts-list-r">
							<i class="back" title="删除" @click="Video.splice(i,1)">-</i>
							<i class="add" title="增加" @click="Video.splice(i+1 ,0, {Title:'',Url:''})">+</i>
						</div>
					</li>
					<div class="account-newposts-icon">
						<i class="add" title="增加" @click="Video.push({Title:'',Url:''})">+</i>
					</div>
				</template>
			</div>
		</ul>

		<ul class="account-newposts-download">
			<span class="title">下载资源:</span>
			<div class="account-newposts-flex">
				<li v-for="(value,i) in Download" :key="i" class="account-newposts-list">
					<div class="account-newposts-list-c">
						<input placeholder="下载名称" required="required" type="text" class="input e" v-model.trim="value.Title" />
						<input placeholder="下载链接" required="required" type="url" class="input e" v-model.trim="value.Url" />
						<input placeholder="提取密码" type="text" class="input e" v-model.trim="value.DownloadPwd" />
						<input placeholder="解压密码" type="text" class="input e" v-model.trim="value.ExtractPwd" />
					</div>
					<div class="account-newposts-list-r">
						<i class="back" title="删除" @click="Download.splice(i,1)">-</i>
						<i class="add" title="增加" @click="Download.splice(i+1 ,0, {Title:'',Url:'',DownloadPwd:'',ExtractPwd:''})">+</i>
					</div>
				</li>
			</div>
			<div class="account-newposts-icon">
				<i class="add" title="增加" @click="Download.push({Title:'',Url:'',DownloadPwd:'',ExtractPwd:''})">+</i>
			</div>
		</ul>
		<ul class="account-newposts-submit">
			<button class="button" type="submit" v-if="Status">投稿</button>
		</ul>
	</form>
</template>
<script>
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import { quillEditor } from 'vue-quill-editor'
export default {
	components: {
		quillEditor
	},
	data() {
		return {
			Status: true,
			CategoryAll: [],
			Music: [],
			Video: [],
			Download: [],
			Title: "",
			Content: "",
			Types: 'standard',
			Thumbnail: "",
			Category: 0,
			editorOption: {				placeholder: "请输入投稿内容...", modules: {
					toolbar: [
						[{ 'header': [1, 2, 3, 4, 5, 6, false] }],  //H标签
						[{ 'color': [] }, { 'background': [] }],          // 字体颜色和背景颜色
						[{ 'link': 'link' }],
						['bold', 'italic', 'underline', 'strike'],  //字体加粗,斜体,下划线,删除线
						['blockquote', 'code-block'],               //blockquote模块,代码模块
						[{ 'list': 'ordered' }, { 'list': 'bullet' }],//枚举清单
						[{ 'script': 'sub' }, { 'script': 'super' }], //字体上标下标
						[{ 'indent': '-1' }, { 'indent': '+1' }],     //缩进
						[{ 'font': [] }],                           //字体
						[{ 'align': [] }],                                //文字对齐
					]
				}
			},
		}
	},
	mounted() {
		if (sessionStorage.getItem("Myoo-NewPosts")) {
			let value = JSON.parse(sessionStorage.getItem("Myoo-NewPosts"))
			this.Download = value.Download;
			this.Video = value.Video;
			this.Music = value.Music
			this.Title = value.Title;
			this.Content = value.Content;
			this.Thumbnail = value.Thumbnail;
			this.Types = value.Types;
			this.Category = value.Category;
		}
	},
	watch: {

		Content() {
			sessionStorage.setItem("Myoo-NewPosts", JSON.stringify({
				Title: this.Title,
				Content: this.Content,
				Thumbnail: this.Thumbnail,
				Types: this.Types,
				Category: this.Category,
				Download: this.Download,
				Video: this.Video,
				Music: this.Music,
			}));
		}
	},
	created() {
		this.$store.commit("load", true)
		this.$axios.get("/account/newposts").then(res => {
			if (res.data.success != 200) {
				this.$router.push("/")
			}
			this.$store.commit("load", false);
			this.CategoryAll = res.data.data.Category ? res.data.data.Category : [];
		}).catch(error => {
			console.log(error);
		});
	},
	methods: {
		Update() {
			if (this.Title == "") {
				this.$store.commit("msg", "投稿标题不能为空");
				return;
			}
			if (this.Category == 0) {
				this.$store.commit("msg", "请选择投稿分类");
				return;
			}
			this.$axios.post("/account/newposts", {
				title: this.Title,
				content: this.Content,
				type: this.Types,
				category: this.Category,
				thumbnail: this.Thumbnail,
				music: JSON.stringify(this.Music),
				video: JSON.stringify(this.Video),
				download: JSON.stringify(this.Download),
			}).then(res => {
				this.$store.commit("msg", res.data.message);
				if (res.data.success == 200) {
					this.Status = false;
				}
			}).catch(error => {
				console.log(error);
			});
		}
	}
}
</script>
<style scoped>
input,
select {
	padding: 5px;
	border: 1px solid #ddd;
	box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.07);
	height: 40px;
}
span.title {
	display: block;
	margin-bottom: 6px;
	width: 100%;
	position: relative;
	min-height: 25px;
	font-size: 16px;
	color: #444;
}
form.account-row {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
form.account-row ul {
	background: #ffffff;
	width: 100%;
	box-shadow: 0 1px 2px rgba(0, 0, 0, 0.2);
	padding: 20px;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	position: relative;
	margin-bottom: 10px;
	border-radius: 5px;
}

ul.account-newposts-header {
	font-size: 23px;
	font-weight: 400;
	margin: 0;
	padding: 9px 0 4px 0;
	line-height: 29px;
	color: #23282d;
	margin-bottom: 10px;
}

form.account-row ul {
	margin-bottom: 15px;
}

ul.account-newposts-title input {
	min-height: 50px;
	width: 100%;
	color: #000;
	font-size: 1.7em;
	padding: 11px 10px;
}
ul.account-newposts-category select,
ul.account-newposts-thumbnail input {
	height: 40px;
	font-size: 15px;
	width: 100%;
}
ul.account-newposts-thumbnail > div {
	width: 150px;
	height: 150px;
	background-size: 100%;
	background-repeat: no-repeat;
	-moz-background-size: 100% 100%;
	margin-top: 10px;
}
.account-newposts-flex button.editor-type {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	padding: 5px;
	border: none;
	color: #424242;
}

.account-newposts-flex {
	display: flex;
	flex-wrap: wrap;
}

.account-newposts-flex button.editor-type.on,
.account-newposts-flex button.editor-type:hover {
	background: #2086bf;
	color: #fff;
}

.account-newposts-icon {
	margin: 10px 0;
	flex: 0 0 100%;
	-webkit-box-flex: 0;
	text-align: center;
	max-width: 100%;
}

ul.account-newposts-types i,
ul.account-newposts-download i {
	background: #2086bf;
	color: #fff;
	height: 30px;
	line-height: 30px;
	border-radius: 2px;
	cursor: pointer;
	text-align: center;
	margin: 0 5px;
	font-size: 20px;
	font-weight: 1000;
	display: inline-block;
}
ul.account-newposts-types i:hover,
ul.account-newposts-download i:hover {
	opacity: 0.9;
}
.account-newposts-icon i {
	width: 100%;
}

li.account-newposts-list {
	margin: 5px 0;
	flex: 0 0 100%;
	-webkit-box-flex: 0;
	display: flex;
	flex-wrap: wrap;
	padding: 0 10px;
	width: 100%;
	position: relative;
	font-size: 14px;
	color: #444;
}

.account-newposts-list-c {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	display: flex;
	flex-wrap: wrap;
	height: 30px;
	line-height: 30px;
}

.account-newposts-list-r {
	flex: 0 0 100%;
	max-width: 90px;
	overflow: hidden;
	height: 30px;
	line-height: 30px;
}

.account-newposts-list-c input {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	margin: 0 5px;
	height: 100%;
}

.account-newposts-list-r i {
	width: 30px;
}

form.account-row button.button {
	background: #00a206;
	color: #fff;
	height: 30px;
	line-height: 30px;
	border-radius: 2px;
	cursor: pointer;
	text-align: center;
	font-size: 15px;
	width: 100%;
	border: none;
	margin-top: 10px;
}
</style>
<style >
form.account-row .ql-editor.ql-blank,
form.account-row .ql-editor {
	background: #fff;
	min-height: 500px !important;
	padding: 20px;
}
</style>
<style scoped>
@media (max-width: 600px) {
	form.account-row ul {
		margin-bottom: 0;
		border-radius: 0;
	}
	.account-newposts-list-c {
		-webkit-box-flex: 0;
		flex: 0 0 100%;
		max-width: 100%;
		height: auto;
	}
	.account-newposts-list-c input {
		-webkit-box-flex: 0;
		flex: 0 0 100%;
		max-width: 45%;
		height: auto;
	}
	.account-newposts-list-r {
		max-width: 100%;
		text-align: right;
	}
}
</style>