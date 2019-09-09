<template>
    <myoo-box :posts="Data"></myoo-box>
</template>
<script>
export default {
    data() {
        return {
            Data: [],
        }
    },
    created() {
        this.$store.commit("load", true)
        this.$axios.get("/category/" + this.$route.params.id + "/" + this.$route.params.page).then(res => {
            if (res.data.data != null) {
                this.Data = res.data.data;
            }
            this.$store.commit("load", false)
        }).catch(error => {
            console.log(error);
        });
    },
}
</script>
<style scoped>
</style>
