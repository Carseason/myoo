<template>
	<ul class="account-favorite">
		<router-view :key="Page"></router-view>
		<navpage :page.sync="Page" :maxpage="MaxPage" :status="0"></navpage>
	</ul>
</template>
<script>
export default {
	data() {
		return {
			Page: this.$route.params.page ? this.$route.params.page : 1,
			MaxCount: 0,
			MaxPage: 0,
		}
	},
	watch: {
		Page(val) {
			this.$router.push("/account/favorite/" + val)
		},

	},
	created() {
		this.$axios.get("/account/favorite/0").then(res => {
			if (res.data.success != 200) {
				this.$router.push("/");
				return;
			}
			this.MaxPage = res.data.data.MaxPage;
			this.MaxCount = res.data.data.MaxCount;
		}).catch(error => {
			console.log(error);
		});
	}
}
</script>
<style scoped>
ul.account-favorite {
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
	min-height: 600px;
}
</style>
