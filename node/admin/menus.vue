<template>
    <div class="body">
        <form class="row" method="post" @submit.prevent="Update()">
            <ul class="row-list">
                <h2 class="list-title">导航菜单</h2>
                <li class="list-list">
                    <div class="list-f">菜单设置</div>
                    <div class="list-c"></div>
                    <div class="list-r">
                        <i class="add" title="增加" @click="Menus.TopMenus.push({Icon:'',Title:'',Url:'/',Child:[]})">+</i>
                    </div>
                </li>
                <template v-for="(item,i) in Menus.TopMenus">
                    <li class="list-list" :key="i">
                        <div class="list-f"></div>
                        <div class="list-c">
                            <input class="input c" type="text" placeholder="图标class" v-model.trim="item.Icon" />
                            <input class="input c" type="text" placeholder="名称" v-model.trim="item.Title" />
                            <select class="select c" v-model.number="item.Url">
                                <option value="/">首页</option>
                                <template v-for="item in CategoryAll">
                                    <option :value="'/category/' + item.Id + '/1'" :key="item.Id">{{item.Name}}</option>
                                </template>
                            </select>
                        </div>
                        <div class="list-r">
                            <i class="add" title="增加" @click="item.Child.push({Title:'',Url:''})">+</i>
                            <i class="back" title="删除" @click="Menus.TopMenus.splice(i,1)">-</i>
                        </div>
                    </li>
                    <template v-for="(child,j) in item.Child">
                        <li class="list-list" :key=" i + '' + j">
                            <div class="list-f"></div>
                            <div class="list-c">
                                <input class="input c" type="text" style="opacity: 0;" />
                                <input class="input c" type="text" placeholder="名称" v-model.trim="child.Title" />
                                <select class="select c" v-model.number="item.Url">
                                    <option value="/">首页</option>
                                    <template v-for="item in CategoryAll">
                                        <option :value="'/category/' + item.Id +'/1'" :key="item.Id">{{item.Name}}</option>
                                    </template>
                                </select>
                            </div>
                            <div class="list-r">
                                <i class="back" title="删除" @click="item.Child.splice(j,1)">-</i>
                            </div>
                        </li>
                    </template>
                </template>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">导航菜单</h2>
                <li class="list-list">
                    <div class="list-f">菜单设置</div>
                    <div class="list-c"></div>
                    <div class="list-r">
                        <i class="add" title="增加" @click="Menus.NavMenus.push({Icon:'',Title:'',Url:'/',Child:[]})">+</i>
                    </div>
                </li>
                <template v-for="(item,i) in Menus.NavMenus">
                    <li class="list-list" :key="i">
                        <div class="list-f"></div>
                        <div class="list-c">
                            <input class="input c" type="text" placeholder="图标class" v-model.trim="item.Icon" />
                            <input class="input c" type="text" placeholder="名称" v-model.trim="item.Title" />
                            <select class="select c" v-model.number="item.Url">
                                <option value="/">首页</option>
                                <template v-for="item in CategoryAll">
                                    <option :value="'/category/' + item.Id + '/1'" :key="item.Id">{{item.Name}}</option>
                                </template>
                            </select>
                        </div>
                        <div class="list-r">
                            <i class="add" title="增加" @click="item.Child.push({Title:'',Url:'/'})">+</i>
                            <i class="back" title="删除" @click="Menus.NavMenus.splice(i,1)">-</i>
                        </div>
                    </li>
                    <template v-for="(child,j) in item.Child">
                        <li class="list-list" :key=" i + '' + j">
                            <div class="list-f"></div>
                            <div class="list-c">
                                <input class="input c" type="text" style="opacity: 0;" />
                                <input class="input c" type="text" placeholder="名称" v-model.trim="child.Title" />
                                <select class="select c" v-model.number="child.Url">
                                    <option :value="'/'">首页</option>
                                    <template v-for="item in CategoryAll">
                                        <option :value="'/category/' + item.Id +'/1'" :key="item.Id">{{item.Name}}</option>
                                    </template>
                                </select>
                            </div>
                            <div class="list-r">
                                <i class="back" title="删除" @click="item.Child.splice(j,1)">-</i>
                            </div>
                        </li>
                    </template>
                </template>
            </ul>
            <ul class="row-list">
                <h2 class="list-title">友联配置</h2>
                <li class="list-list">
                    <div class="list-f">友情链接</div>
                    <div class="list-c"></div>
                    <div class="list-r">
                        <i class="add" title="增加" @click="Menus.Links.push({Title:'',Url:''})">+</i>
                    </div>
                </li>
                <template v-for="(item,i) in Menus.Links">
                    <li class="list-list" :key="i">
                        <div class="list-f"></div>
                        <div class="list-c">
                            <input class="input c" type="text" placeholder="名称" v-model.trim="item.Title" />
                            <input class="input c" type="url" placeholder="链接" v-model.trim="item.Url" />
                        </div>
                        <div class="list-r">
                            <i class="back" title="删除" @click="Menus.Links.splice(i,1)">-</i>
                        </div>
                    </li>
                </template>
            </ul>
            <button class="button" type="submit">保存</button>
        </form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            CategoryAll: [],
            Menus: {
                TopMenus: [],
                NavMenus: [],
                Links: []
            }
        };
    },
    created() {
        this.$axios.all([
            this.$axios.get("/admin/menus"),
            this.$axios.get("/admin/category")
        ]).then(this.$axios.spread((menus, category) => {
            if (menus.data.data.NavMenus != null) {
                this.Menus.NavMenus = menus.data.data.NavMenus;
            }
            if (menus.data.data.TopMenus != null) {
                this.Menus.TopMenus = menus.data.data.TopMenus;
            }
            if (menus.data.data.Links != null) {
                this.Menus.Links = menus.data.data.Links;
            }
            if (category.data.data != null) {
                this.CategoryAll = category.data.data;
            }
        }));
    },
    methods: {
        Update() {
            this.$axios.post("/admin/menus", {
                data: JSON.stringify(this.Menus)
            }).then(res => {
                this.$store.commit("msg", res.data.message);
            }).catch(error => {
                console.log(error);
            });
            return;
        }
    }
};
</script>
<style scoped>
form.row {
    margin: auto;
    border: 1px solid rgba(255, 255, 255, 0.29);
    box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
    border-radius: 1rem;
    width: 100%;
    margin-bottom: 150px;
}

ul.row-list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    border-bottom: solid 1px #86828224;
}

h2.list-title {
    background: #0085ba;
    color: #fff;
    display: inline-block;
    text-decoration: none;
    font-size: 13px;
    line-height: 30px;
    height: 30px;
    margin: 0;
    width: 100px;
    text-align: center;
    cursor: pointer;
}

li.list-list {
    display: flex;
    flex-wrap: wrap;
    padding: 0 10px;
    margin-bottom: 6px;
    width: 100%;
    position: relative;
    min-height: 25px;
    font-size: 1rem;
    color: #717171;
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
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
    vertical-align: middle;
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

.list-r {
    flex: 0 0 100%;
    max-width: 120px;
    overflow: hidden;
    text-align: left;
}

.list-c > * {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
    margin: 0 5px;
}

form.row img {
    max-height: 50px;
    height: 100%;
    max-width: 50px;
    width: 100%;
    text-align: center;
    vertical-align: middle;
    object-fit: cover;
    padding: 10px;
}

input,
select,
textarea {
    padding: 5px 10px;
    background: rgba(255, 255, 255, 0);
    border: 1px solid rgb(191, 191, 191);
}
textarea {
    overflow: auto;
    padding: 2px 6px;
    line-height: 1.4;
    resize: vertical;
    height: 100px;
}
button.button {
    width: 100%;
    height: 30px;
    cursor: pointer;
    background: #008ec2;
    color: #fff;
    border: 0;
    border-radius: 4px;
    margin: 10px 0;
}
button:hover {
    opacity: 0.9;
}
li.list-list i {
    background: #00a206;
    color: #fff;
    height: 30px;
    width: 30px;
    line-height: 30px;
    border-radius: 2px;
    cursor: pointer;
    text-align: center;
    margin: 0 5px;
    font-size: 20px;
    font-weight: 1000;
    display: inline-block;
}
</style>

