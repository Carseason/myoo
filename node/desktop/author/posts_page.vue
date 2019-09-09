<template>
    <myoo-box :posts="Posts"></myoo-box>
</template>
<script>
export default {
    data() {
        return {
            Posts: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/author/" + this.$route.params.id + "/posts/" + this.$route.params.page).then(res => {
            this.Posts = res.data.data ? res.data.data : [];
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
</style>