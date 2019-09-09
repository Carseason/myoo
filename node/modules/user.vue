<template>
    <div class="myoo-usermodules">
        <template v-for="value in user">
            <li class="row-user-list" :key="value.Id">
                <router-link :to="'/author/' + value.Id" class="row-user-links">
                    <img :src="value.Avatar" class="avatar" />
                    <span>{{value.Name}}</span>
                </router-link>
            </li>
        </template>
    </div>
</template>
<script>
export default {
    props: ["user"]
}
</script>
<style scoped>
li.row-user-list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100px;
    margin: 0 10px;
}

li.row-user-list img.avatar {
    width: 100px;
    height: 100px;
    padding: 2px;
    border: 1px solid #e8e6e6;
    cursor: pointer;
    object-fit: cover;
    border-radius: 50%;
}

li.row-user-list span {
    display: block;
    width: 100px;
    padding: 5px;
    font-size: 14px;
    color: #9a9a9a;
    text-align: center;
}
li.row-user-list span:hover {
    color: #00a1d6ba;
}
.myoo-usermodules {
    display: flex;
    flex-wrap: wrap;
    margin: 10px 0;
}
</style>
