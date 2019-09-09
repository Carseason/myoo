<template>
    <div class="body">
        <form class="row" method="post" @submit.prevent="Update()">
            <h1 class="row-title">用户信息</h1>

            <ul class="row-list">
                <li class="list-list">
                    <div class="list-f">用户头像</div>
                    <div class="list-c">
                        <img :src="User.Avatar" class="avatar" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">用户名</div>
                    <div class="list-c">
                        <input type="text" placeholder="用户名" required="required" v-model.trim="User.Name" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">用户邮箱</div>
                    <div class="list-c">
                        <input type="email" placeholder="用户邮箱" required="required" v-model.trim="User.Email" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">用户等级</div>
                    <div class="list-c">
                        <input type="number" placeholder="用户等级" min="1" required="required" v-model.trim="User.Lv" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">用户签名</div>
                    <div class="list-c">
                        <input type="text" placeholder="签名" v-model.trim="User.Description" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">用户密码</div>
                    <div class="list-c">
                        <input type="text" placeholder="留空则不更改" v-model.trim="User.Pwd" />
                    </div>
                </li>
            </ul>

            <button v-if="$route.params.id == 0" class="button" type="submit">创建用户</button>
            <button v-else class="button" type="submit">更新</button>
        </form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            User: {
                Id: 0,
                Name: "",
                Lv: 1,
                Email: "",
                Pwd: "",
                Description: "",
                Avatar: "",
            }
        }
    },
    created() {
        if (this.$route.params.id > 0) {
            this.$axios.get('/admin/user/' + this.$route.params.id).then(res => {
                if (res.data.success == 403) {
                    this.$router.push("/");
                    return;
                }
                this.User = res.data.data;
            }).catch(res => {
                console.log(error);
            });
        }
    },
    methods: {
        Update() {
            if (this.$route.params.id == 0 && this.User.Pwd == "") {
                this.$store.commit("msg", "用户密码不能为空");
                return;
            }
            if (this.User.Email == "") {
                this.$store.commit("msg", "用户邮箱不能为空");
                return;
            }
            if (this.User.Name == "") {
                this.$store.commit("msg", "用户名不能为空");
                return;
            }
            if (this.User.Lv <= 0) {
                this.$store.commit("msg", "用户等级不能为异常");
                return;
            }
            this.$axios.post('/admin/user/' + this.$route.params.id, {
                name: this.User.Name,
                lv: this.User.Lv,
                email: this.User.Email,
                pwd: this.User.Pwd,
                description: this.User.Description,
            }).then(res => {
                this.$store.commit("msg", res.data.message);
            }).catch(error => {
                console.log(error);
            });
        }
    }


}
</script>
<style scoped>
form.row {
    margin: auto;
    border: 1px solid rgba(255, 255, 255, 0.29);
    box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
    border-radius: 1rem;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
}
h1.row-title {
    font-size: 23px;
    font-weight: 400;
    margin: 0;
    padding: 9px 0 4px 0;
    line-height: 29px;
    color: #23282d;
    margin-left: 20px;
}
li.list-list {
    display: flex;
    flex-wrap: wrap;
    padding: 0 10px;
    margin-bottom: 10px;
    width: 100%;
    position: relative;
    min-height: 30px;
    font-size: 14px;
    color: #444;
}
li.list-list > div {
    min-height: 30px;
    line-height: 30px;
}
.list-f {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 10%;
    padding: 0 10px;
}
.list-c {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
    height: 100%;
    display: -webkit-flex;
    display: flex;
    flex-wrap: wrap;
}
.list-c > * {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
    margin: 0 5px;
}
input,
select {
    padding: 5px;
    border: none;
    box-shadow: inset 0 1px 2px rgba(0, 0, 0, 0.07);
    background: rgba(255, 255, 255, 0);
}
button.button {
    width: 100%;
    height: 30px;
    cursor: pointer;
    background: #008ec2;
    color: #fff;
    border: 0;
    border-radius: 4px;
    margin: 10px;
}
img.avatar {
    height: 100%;
    width: 100%;
    max-width: 100px;
    max-height: 100px;
}
</style>
