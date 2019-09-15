<template>
	<div class="msg-mask">
		<div class="msg-bg"></div>
		<div class="msg-box">
			<div class="msg-box__header">
				<div class="msg-box__title">
					<span>{{$store.state.Message.Title}}</span>
				</div>
			</div>
			<div class="msg-box__content">
				<div class="msg-box__msg">
					<p v-html="$store.state.Message.Content"></p>
				</div>
			</div>
			<div class="msg-box__btns">
				<button type="button" class="msg-button-primary" @click="Message">
					<span>确定</span>
				</button>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
		}
	},
	methods: {
		Message() {
			this.$store.commit("msg", "")
		}
	}

}
</script>
<style scoped>
.msg-mask {
	position: fixed;
	top: 0;
	bottom: 0;
	left: 0;
	right: 0;
	margin: auto;
	width: 100%;
	height: 100%;
	text-align: center;
	z-index: 2001;
	display: flex;
}
.msg-bg {
	width: 100%;
	height: 100%;
	opacity: 0.5;
	background: #000;
	position: absolute;
}
.msg-box {
	background: #fff;
	position: relative;
	margin: auto;
	width: 400px;
	padding-bottom: 10px;
	vertical-align: middle;
	background-color: #fff;
	border-radius: 4px;
	border: 1px solid #ebeef5;
	font-size: 18px;
	box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
	text-align: left;
	overflow: hidden;
	backface-visibility: hidden;
	display: flex;
	flex-wrap: wrap;
}
.msg-box > div {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
.msg-box__header {
	position: relative;
	padding: 15px 15px 10px;
}
.msg-box__title {
	padding-left: 0;
	margin-bottom: 0;
	font-size: 18px;
	line-height: 1;
	color: #303133;
}

.msg-box__content {
	position: relative;
	padding: 10px 15px;
	color: #606266;
	font-size: 14px;
}
.msg-box__msg {
	margin: 0;
}
.msg-box__msg p {
	margin: 0;
	line-height: 24px;
}
.msg-box__btns {
	padding: 5px 15px 0;
	text-align: right;
}
button.msg-button-primary {
	color: #fff;
	background-color: #409eff;
	border-color: #409eff;
	display: inline-block;
	line-height: 1;
	white-space: nowrap;
	cursor: pointer;
	border: 1px solid #dcdfe6;
	-webkit-appearance: none;
	text-align: center;
	box-sizing: border-box;
	outline: none;
	margin: 0;
	transition: 0.1s;
	font-weight: 500;
	-moz-user-select: none;
	-webkit-user-select: none;
	-ms-user-select: none;
	padding: 12px 20px;
	font-size: 14px;
	border-radius: 4px;
	padding: 9px 15px;
	font-size: 12px;
	border-radius: 3px;
}
button.msg-button-primary:hover {
	opacity: 0.9;
}
@media (max-width: 500px) {
	.msg-box {
		width: 90%;
	}
}
</style>

