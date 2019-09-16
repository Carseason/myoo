<template>
	<div>
		<myoo-header></myoo-header>

		<!-- <transition name="transitions" appear> -->
		<router-view></router-view>
		<!-- </transition> -->
		<myoo-footer></myoo-footer>
	</div>
</template>
<script>
import Header from "./header/header.vue";
import Footer from "./footer/footer.vue";
export default {
	components: {
		"myoo-header": Header,
		"myoo-footer": Footer
	}
};
</script>
<style scoped>
@import url("https://at.alicdn.com/t/font_1376574_t5oaz7ekxm.css");
</style>