<template>
	<div id="recommend">
		<h2 class="recommend-title">或许你会喜欢:</h2>
		<myoo-box :posts="Recommend" types="mobile"></myoo-box>
	</div>
</template>
<script>
export default {
	props: ["id"],
	data() {
		return {
			Recommend: [],
		}
	},
	created() {
		this.$axios.get('/recommend/' + this.id).then(res => {
			this.Recommend = res.data.data ? res.data.data : [];
		}).catch(error => {
			console.log(error);
		});
	},
}
</script>
<style scoped>
h2.recommend-title {
	font-size: 14px;
	margin: 0;
	padding: 0;
	color: #acacac;
}
</style>
