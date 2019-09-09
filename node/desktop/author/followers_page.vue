<template>
    <myoo-user :user="Followers"></myoo-user>
</template>
<script>
export default {
    data() {
        return {
            Followers: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/author/" + this.$route.params.id + "/followers/" + this.$route.params.page).then(res => {
            this.Followers = res.data.data ? res.data.data : [];
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
</style>
