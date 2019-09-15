<template>
	<ul class="account-row">
		<li class="account-list">
			<div class="account-list-title">
				<i class="iconfont icon-denglu2 icon-wodeziliaoicon"></i>
				<span>我的资料</span>
			</div>
			<div class="account-dialiy-exp">
				<li class="account-item">
					<label class="account-label">ID:</label>
					<div class="account-el">{{User.Id}}</div>
				</li>
				<li class="account-item">
					<label class="account-label">邮箱:</label>
					<div class="account-el">{{User.Email}}</div>
				</li>
				<li class="account-item">
					<label class="account-label">昵称:</label>
					<div class="account-el">{{User.Name}}</div>
				</li>
				<li class="account-item">
					<label class="account-label">注册于:</label>
					<div class="account-el">{{User.Date | dateForm2}} | 您是本站第 {{User.Id}} 位用户</div>
				</li>
				<li class="account-item">
					<label class="account-label">Lv:</label>
					<div class="account-el">
						<span class="capabilities-lv">{{User.Group.Lv}}</span>
						<span class="capabilities-nicename">{{User.Group.NickName}}</span>
					</div>
				</li>
				<!-- <li class="account-item">
                    <label class="account-label">勋章:</label>
                    <div class="account-el">
                        <template v-for="(value ,i ) in User.Group.Medal">
                            <img :src="value" class="medal" :key="i" />
                        </template>
                    </div>
				</li>-->
			</div>
		</li>
		<li class="account-list">
			<div class="account-list-title">
				<i class="iconfont icon-denglu2 icon-qianming"></i>
				我的签名
			</div>
			<div class="account-dialiy-exp">
				<li class="account-item">
					<label class="account-label">签名:</label>
					<div class="account-el">
						<textarea placeholder="暂无签名。。。" class="textarea a" v-model="User.Description"></textarea>
					</div>
				</li>
				<li class="account-item">
					<label class="account-label"></label>
					<div class="account-el">
						<button type="submit" @click="Update('description')">更新签名</button>
					</div>
				</li>
			</div>
		</li>
		<li class="account-list">
			<div class="account-list-title">
				<i class="iconfont icon-denglu2 icon-mima"></i>我的密码
			</div>
			<div class="account-dialiy-exp">
				<li class="account-item">
					<label class="account-label">密码:</label>
					<div class="account-el">
						<input type="password" v-model.trim="Pwd" placeholder="请输入6位以上的密码" class="input a" />
					</div>
				</li>
				<li class="account-item">
					<label class="account-label"></label>
					<div class="account-el">
						<button type="submit" @click="Update('pwd')">更新密码</button>
					</div>
				</li>
			</div>
		</li>
		<li class="account-list">
			<div class="account-list-title">
				<i class="iconfont icon-denglu2 icon-touxiang"></i>我的头像
			</div>
			<div class="account-dialiy-exp">
				<label class="account-label" style="max-width: 20%;"></label>
				<label class="account-update">
					<i class="iconfont icon-denglu2 icon-Image"></i>
					<span>上传择头像</span>
					<input type="file" name="avatar" id="uppic" accept="image/jpeg, image/jpg, image/png" hidden @change="changeImage($event)" />
				</label>
				<div class="border-line"></div>
				<div class="account-el" style="margin: auto 0px;">
					<img :src="User.Avatar" class="avatar" />
					<span class="pre-info">当前头像</span>
				</div>
			</div>
			<p class="descript">请选择图片上传：大小150 * 150像素支持JPG、PNG等格式，图片需小于2M</p>
		</li>
		<li class="account-list" v-if="CropperStatus">
			<cropper :imgfile="ImgFile"></cropper>
		</li>
	</ul>
</template>
<script>
import cropper from "@/desktop/account/cropper.vue";
export default {
	components: {
		cropper: cropper
	},
	data() {
		return {
			User: this.$store.state.Auth,
			Status: false,
			Pwd: "",
			ImgFile: "",
			CropperStatus: false,
		};
	},
	mounted() { },
	watch: {
		"$store.state.Auth"() {
			this.User = this.$store.state.Auth;
		},

	},
	created() {

	},
	methods: {
		changeImage(e) {
			let reader = new FileReader();
			let file = e.target.files[0];
			if (!/\.(jpg|png|jpge)$/.test(file.name)) {
				console.log(file.name)
				this.$store.commit("msg", "请上传 jpg | png | jpge | jpeg 格式的图像文件");
				return;
			}
			reader.readAsDataURL(file); // 读出 base64
			reader.onloadend = () => {
				this.ImgFile = reader.result;
			};
			this.CropperStatus = true;
		},
		Update(types) {
			if (types == "pwd" && this.Pwd.length < 5) {
				this.$store.commit("msg", "密码不能少于6位");
				return;
			}
			this.$axios.post("/account/info", {
				type: types,
				description: this.User.Description,
				pwd: this.Pwd,
			}).then(res => {
				this.$store.commit("msg", res.data.message);
				if (types == "pwd" && res.data.success == 200) {
					this.$store.commit("localstorage")
					setTimeout(() => {
						window.location.href = "/"
					}, 1500);
				}

			}).catch(error => {
				console.log(error);
			});
			return;
		},
	}
};
</script>
<style scoped>
ul.account-row {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
li.account-list {
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
.account-list-title {
	font-size: 20px;
	color: #222;
	vertical-align: top;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	height: 50px;
}
.account-list-title i {
	margin-right: 10px;
	color: #717171;
}

.account-dialiy-exp {
	-webkit-box-flex: 0;
	max-width: 100%;
	display: flex;
	flex-wrap: wrap;
	flex: 0 0 100%;
	margin: 20px 0px;
}

li.account-item {
	margin-bottom: 22px;
	flex-wrap: wrap;
	display: flex;
	position: relative;
	-webkit-box-flex: 0;
	max-width: 100%;
	flex: 0 0 100%;
	font-size: 14px;
}

label.account-label {
	-webkit-box-flex: 0;
	max-width: 100px;
	text-align: right;
	margin-right: 10px;
	vertical-align: middle;
	color: rgb(72, 87, 106);
	flex: 0 0 100%;
}

.account-el {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
}
label.account-update {
	cursor: pointer;
	background: #f1f2f5;
	width: 180px;
	height: 180px;
	max-width: 180px;
	border: 1px solid #e5e9ef;
	border-radius: 4px;
	transition: all 0.6s ease;
}
label.account-update i {
	margin: 50px auto 0;
	width: 50px;
	height: 50px;
	font-size: 3rem;
	font-weight: 100;
	color: #a5a5a5;
	display: block;
}
label.account-update span {
	display: block;
	font-family: MicrosoftYaHei;
	font-size: 16px;
	color: #6d757a;
	line-height: 20px;
	margin-top: 10px;
	padding: 0 12px;
	letter-spacing: 0;
	text-align: center;
}
.border-line {
	height: 182px;
	width: 1px;
	background: #e5e9ef;
	margin: 0 40px;
}
.account-el img.avatar {
	width: 150px;
	height: 150px;
	border-radius: 50%;
}
span.pre-info {
	display: block;
	margin-top: 20px;
	font-size: 12px;
	color: #99a2aa;
	margin-left: 25px;
}
p.descript {
	line-height: 16px;
	color: #99a2aa;
	text-align: center;
	width: 100%;
}

textarea.textarea.a {
	padding: 2px 6px;
	resize: vertical;
	overflow: auto;
	width: 300px;
	height: 70px;
	line-height: 1.5;
	color: #1f2d3d;
	background-color: #fff;
	background-image: none;
	border: 1px solid #bfcbd9;
	border-radius: 4px;
}

form.account-form {
	width: 100%;
}

li.account-item button {
	color: #fff;
	width: 110px;
	height: 30px;
	background: #00a1d6;
	border: none;
	border-radius: 4px;
}
li.account-item input.input.a {
	overflow: auto;
	width: 300px;
	resize: vertical;
	padding: 5px 10px;
	color: #1f2d3d;
	background-color: #fff;
	background-image: none;
	border: 1px solid #bfcbd9;
	border-radius: 4px;
}

.capabilities-nicename {
	margin: 0 8px;
	color: #fff;
	background: #ff6688;
	padding: 0 5px;
	border: 1px solid #ff6688;
	border-radius: 6px;
}
img.medal {
	max-width: 100%;
	margin: 10px;
}
</style>
<style scoped>
@media (max-width: 600px) {
	li.account-list {
		margin-bottom: 0;
		border-radius: 0;
	}
}
</style>