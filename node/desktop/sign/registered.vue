<template>
	<div class="sign-login">
		<div class="sign-title">
			<span>注册</span>
		</div>
		<div class="sign-body">
			<div class="sign-body-f">
				<div class="sign-body-bg"></div>
			</div>
			<div class="sign-body-r">
				<form method="post" @submit.prevent="Update()">
					<input type="email" name="email" placeholder="您的邮箱" v-model="User.Email" required="required" />
					<input type="text" name="pass" placeholder="用户名" v-model="User.Name" required="required" />
					<input type="password" name="pass" placeholder="您的密码" v-model="User.Pwd" required="required" />
					<input type="password" name="pass" placeholder="重复输入密码" v-model="User.Pwd2" required="required" />
					<p class="message">{{Message}}</p>
					<template v-if="Status">
						<button type="submit">注册</button>
					</template>
					<template v-else>
						<div class="registered-off">很抱歉,当前注册已关闭!</div>
					</template>
				</form>
			</div>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			Message: "",
			Status: true,
			User: {
				Email: "",
				Name: "",
				Pwd: "",
				Pwd2: "",
			},

		}
	},
	created() {
		this.$store.commit("load", true)
		this.$axios.get("/sign/registered").then(res => {
			this.$store.commit("load", false);
			if (res.data.data != null) {
				this.Status = res.data.data;
			}
		}).catch(error => {
			console.log(error);
		});
	},
	watch: {
		'$store.state.Title': {
			immediate: true,
			handler() {
				this.$store.commit("title", "注册");
			}
		},
	},
	mounted() {
	},
	methods: {
		Update() {
			if (this.User.Email == "" || this.User.Pwd == "") {
				this.$store.commit("msg", "请输入邮箱或密码");
				return;
			}
			if (this.User.Pwd != this.User.Pwd2) {
				this.$store.commit("msg", "两次输入的密码不一致");
				return;
			}
			this.$store.commit("load", true);
			this.$axios.post('/sign/registered', {
				email: this.User.Email,
				name: this.User.Name,
				pwd: this.User.Pwd,
				pwd2: this.User.Pwd2,
			}).then(res => {
				this.$store.commit("load", false);
				this.Message = res.data.message;
				if (res.data.success === 200) {
					this.$store.commit("localstorage", res.data.data)
					this.$store.commit("msg", "注册成功,正在跳转至首页....")
					setTimeout(() => {
						if (document.referrer == "") {
							window.location.href = "/"
						}
						window.location.href = document.referrer
					}, 1500);
				}
			}).catch(error => {
				console.log(error);
			});
			return;
		}
	}
}
</script>
<style scoped>
.sign-login {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 60%;
	margin: 0 auto;
}
.sign-title {
	height: 28px;
	margin: 25px auto 50px auto;
	border-bottom: 1px solid #ddd;
	text-align: center;
}
.sign-title span {
	height: 56px;
	line-height: 56px;
	margin: 0 auto;
	padding: 0 20px;
	font-size: 30px;
	background: #fff;
	text-align: center;
}
.sign-body {
	display: flex;
	flex-wrap: wrap;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}

.sign-body-f {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50%;
	padding: 20px;
}
.sign-body-bg {
	background-image: url("");
	background-position: center;
	overflow: hidden;
	position: relative;
	vertical-align: middle;
	width: 100%;
	height: 100%;
	object-fit: cover;
	cursor: pointer;
	background-size: cover;
	background-repeat: no-repeat;
	-moz-background-size: 100% 100%;
	border-radius: 5px;
}
.sign-body-r {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 50%;
}

.sign-body form {
	border-left: 1px solid #ddd;
	padding: 20px;
	width: 100%;
}
.sign-body form input {
	padding: 10px;
	border: 1px solid rgba(151, 151, 151, 0.47843137254901963);
	width: 100%;
	height: 40px;
	border-radius: 5px;
	margin-bottom: 15px;
}

button:hover {
	opacity: 0.9;
}
.sign-body form button {
	width: 100%;
	height: 40px;
	line-height: 40px;
	font-size: 16px;
	color: #fff;
	background-color: #00a0d8;
	border-radius: 5px;
	border: none;
}
.registered-off {
	width: 100%;
	height: 40px;
	line-height: 40px;
	font-size: 16px;
	color: #fff;
	background-color: #232323;
	border-radius: 5px;
	border: none;
	text-align: center;
	cursor: not-allowed;
}
</style>
<style scoped>
@media (max-width: 600px) {
	.sign-login {
		max-width: 100%;
	}
	.sign-body-f {
		display: none;
	}
	.sign-body-r {
		max-width: 100%;
	}
}
</style>