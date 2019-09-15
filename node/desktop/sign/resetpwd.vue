<template>
	<div class="sign-login">
		<div class="sign-title">
			<span>修改密码</span>
		</div>
		<div class="sign-body">
			<form method="post" @submit.prevent="Update()">
				<input type="password" name="pass" placeholder="您的密码" v-model="User.Pwd" required="required" />
				<input type="password" name="pass" placeholder="重复输入密码" v-model="User.Pwd2" required="required" />
				<p class="message">{{Message}}</p>
				<button type="submit">重置密码</button>
			</form>
		</div>
	</div>
</template>
<script>
export default {
	data() {
		return {
			Message: "",
			User: {
				Pwd: "",
				Pwd2: "",
			},

		}
	},
	created() {
	},
	watch: {
		'$store.state.Title': {
			immediate: true,
			handler() {
				this.$store.commit("title", "修改密码");
			}
		},
	},
	mounted() {
		if (this.$route.params.id == null || this.$route.params.id == "") {
			this.$router.push("/");
		}
	},
	methods: {
		Update() {
			if (this.User.Pwd != this.User.Pwd2) {
				this.$store.commit("msg", "两次输入的密码不一致");
				return;
			}
			this.$store.commit("load", true);
			this.$axios.post('/sign/check/' + this.$route.params.id, {
				pwd: this.User.Pwd,
				pwd2: this.User.Pwd2,
			}).then(res => {
				this.Message = res.data.message;
				this.$store.commit("load", false);
				this.$store.commit("msg", res.data.message)
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

.sign-body form {
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