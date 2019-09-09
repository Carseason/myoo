<template>
    <myoo-user :user="Fans"></myoo-user>
</template>
<script>
export default {
    data() {
        return {
            Fans: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/author/" + this.$route.params.id + "/fans/" + this.$route.params.page).then(res => {
            this.Fans = res.data.data ? res.data.data : [];
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
</style>
