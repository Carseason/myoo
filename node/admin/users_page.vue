<template>
    <div class="row">
        <template v-for="value in Users">
            <li class="my-posts-list" :key="value.Id">
                <div>
                    <img :src="value.Avatar" class="avatar" />
                </div>
                <div>{{value.Id}}</div>
                <div>{{value.Name}}</div>
                <div>{{value.Email}}</div>
                <div>{{value.Lv}}</div>
                <div>{{value.Date | dateForm2 }}</div>
                <div>
                    <router-link :to="'/admin/user/'+value.Id">修改</router-link>
                    <span class="block"></span>
                    <router-link :to="'/author/'+value.Id">查看</router-link>
                </div>
            </li>
        </template>
    </div>
</template>

<script>
export default {
    data() {
        return {
            Users: []
        }
    },

    created() {
        this.$axios.get('/admin/users/' + this.$route.params.page).then(res => {
            if (res.data.data != null) {
                this.Users = res.data.data;
            }
        }).catch(error => {
            console.log(error);
        });
    },
    methods: {

    }
}
</script>
<style scoped>
li.my-posts-list {
    height: 70px;
    line-height: 70px;
    display: flex;
    flex-wrap: wrap;
    position: relative;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    font-size: 14px;
    margin-bottom: 5px;
    border-bottom: 1px solid #0000000d;
}
li.my-posts-list div {
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}
li.my-posts-list > * {
    text-align: center;
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 14.2%;
    height: 100%;
    color: #666;
}
li.my-posts-list img.thumbnail {
    max-width: 100px;
    width: 100%;
}
li.my-posts-list img.avatar {
    width: 50px;
    height: 50px;
    padding: 0;
}
</style>
