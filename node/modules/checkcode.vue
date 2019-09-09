<template>
    <div class="check-code">
        <div class="checkcodebase">
            <img alt="加载失败,点我重试" :src="$store.state.CheckCode.BaseCode" @click="UpdateCheckCode()" title="验证码" />
        </div>
        <input type="text" class="checkcodeInput" name="code" placeholder="验证码" v-model="InputCode" required="required" />
    </div>
</template>
<script>
export default {
    name: "checkcode",
    props: ["checkcode"],
    data() {
        return {
            InputCode: "",
        }
    },
    created() {
        this.$store.commit("checkcode")
    },
    watch: {
        InputCode(val) {
            this.$emit('update:checkcode', val)
        },
    },
    methods: {
        UpdateCheckCode() {
            this.$store.commit("checkcode")
        }
    }
}
</script>
<style scoped>
.check-code {
    display: flex;
    flex-wrap: wrap;
    margin-top: 5px;
}
.check-code input.checkcodeInput {
    max-width: 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    padding: 10px;
    border: 1px solid rgba(151, 151, 151, 0.47843137254901963);
    width: 100%;
    border-radius: 5px;
    margin-bottom: 15px;
}
.check-code .checkcodebase {
    max-width: 100%;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    margin: 0 auto;
    text-align: center;
    margin-bottom: 10px;
}
.check-code .checkcodebase img {
    max-width: 100%;
    margin: 10px 0;
    border-radius: 4px;
    cursor: pointer;
}
</style>
