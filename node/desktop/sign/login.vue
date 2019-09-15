<template>
	<div class="sign-login">
		<div class="sign-title">
			<span>登录</span>
		</div>
		<div class="sign-body">
			<div class="sign-body-f">
				<div class="sign-body-bg"></div>
			</div>
			<div class="sign-body-r">
				<form method="post" @submit.prevent="Update()">
					<input type="email" name="email" placeholder="您的邮箱" v-model="User.Email" required="required" />
					<input type="password" name="pass" placeholder="您的密码" v-model="User.Pwd" required="required" />
					<!-- <myoo-checkcode :checkcode.sync="Check.Code"></myoo-checkcode> -->
					<p class="message">{{Message}}</p>
					<button type="submit">登录</button>
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
			User: {
				Email: "",
				Pwd: "",
			},
			// Check: {
			//     Code: "",
			//     Status: 0,
			// }
		}
	},
	created() {
	},
	watch: {
		'$store.state.Title': {
			immediate: true,
			handler() {
				this.$store.commit("title", "登录");
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
			// if (this.Check.Code.length != 6 || this.Check.Key == "") {
			//     this.$store.commit("msg", "验证码错误");
			//     return;
			// }
			this.$store.commit("load", true);
			this.$axios.post('/sign/login', {
				email: this.User.Email,
				pwd: this.User.Pwd,
				// checkcode: this.Check.Code,
				// checkkey: this.$store.state.CheckCode.Key,
			}).then(res => {
				this.$store.commit("load", false);
				this.Message = res.data.message;
				if (res.data.success === 200) {
					this.$store.commit("localstorage", res.data.data)
					this.$store.commit("msg", "登录成功,正在跳转....")
					setTimeout(() => {
						window.location.href = "/";//document.referrer
					}, 1500);
					//    this.$router.go(-1);
				} else {
					//    this.$store.commit("checkcode")
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