<template>
    <div>
        <myoo-header></myoo-header>
        <transition name="transitions" appear>
            <router-view></router-view>
        </transition>
        <myoo-footer></myoo-footer>
    </div>
</template>
<script>
import Header from "./header.vue";
import Footer from "./footer.vue";
export default {
    mounted() {
        this.$store.commit("title");
    },
    components: {
        "myoo-header": Header,
        "myoo-footer": Footer
    },
}
</script>
<style scoped>
@import url("https://at.alicdn.com/t/font_1376574_4f92lp5nb7c.css");
</style>
