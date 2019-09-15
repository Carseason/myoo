<template>
    <div class="row">
        <form class="row" method="post" @submit.prevent="Update()">
            <template v-for="(value,i) in BoxItems">
                <ul class="row-list" :key="i">
                    <h2 class="list-title">模块 - {{ i}}</h2>
                    <li class="list-list">
                        <div class="list-f">模块标题</div>
                        <div class="list-c">
                            <input type="text" placeholder="标题" class="input c" required="required" v-model.trim="value.Title" />
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">模块链接</div>
                        <div class="list-c">
                            <input type="url" placeholder="链接" class="input c" v-model.trim="value.Url" />
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">模块图标</div>
                        <div class="list-c">
                            <input type="url" placeholder="图标URL" class="input c" v-model.trim="value.Icon" />
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">模块类型</div>
                        <div class="list-c">
                            <select class="select c" v-model.number="value.Module">
                                <option :value="0">幻灯片模块</option>
                                <option :value="1">模块一</option>
                            </select>
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">模块数量</div>
                        <div class="list-c">
                            <input type="number" min="1" placeholder="显示数量" class="input c" v-model.number="value.Number" />
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">排序方式</div>
                        <div class="list-c">
                            <select class="select c" v-model.number="value.Orderby">
                                <option :value="0">默认</option>
                                <option :value="1">随机</option>
                                <option :value="2">浏览</option>
                            </select>
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">小工具</div>
                        <div class="list-c">
                            <select class="select d" v-model.number="value.Widget.Enabled">
                                <option :value="0">关闭</option>
                                <option :value="1">排行榜</option>
                                <option :value="2">轮播图</option>
                                <option :value="3">HTML代码</option>
                            </select>
                        </div>
                    </li>
                    <template v-if="value.Widget.Enabled >0">
                        <template v-if="value.Widget.Enabled == 3">
                            <li class="list-list">
                                <div class="list-f">小工具HTML代码</div>
                                <div class="list-c">
                                    <textarea placeholder="此代码将显示在该模块右侧" class="textarea a" v-model.trim="value.Widget.Code"></textarea>
                                </div>
                            </li>
                        </template>
                        <template v-else>
                            <li class="list-list">
                                <div class="list-f">小工具名称</div>
                                <div class="list-c">
                                    <input class="input d" type="text" placeholder="显示名称" v-model.trim="value.Widget.Title" />
                                </div>
                            </li>
                            <li class="list-list">
                                <div class="list-f">小工具数量</div>
                                <div class="list-c">
                                    <input class="input d" type="number" min="0" placeholder="显示数量" v-model.number="value.Widget.Number" />
                                </div>
                            </li>
                            <li class="list-list">
                                <div class="list-f">小工具排序</div>
                                <div class="list-c">
                                    <select class="select d" v-model.number="value.Widget.Orderby">
                                        <option :value="0">默认</option>
                                        <option :value="1">随机</option>
                                        <option :value="2">浏览</option>
                                    </select>
                                </div>
                            </li>
                        </template>
                    </template>
                    <li class="list-list">
                        <div class="list-f">分类隐藏</div>
                        <div class="list-c">
                            <ul class="tabs-panel">
                                <template v-for="c in Category">
                                    <label :key="c.Id">
                                        <input type="checkbox" :value="c.Id" v-model.number="value.Category" />
                                        {{c.Name}}
                                    </label>
                                </template>
                            </ul>
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f">自定义代码</div>
                        <div class="list-c">
                            <textarea placeholder="此代码将显示在该模块下方" class="textarea a" v-model.trim="value.Code"></textarea>
                        </div>
                    </li>
                    <li class="list-list">
                        <div class="list-f"></div>
                        <div class="list-c">
                            <i @click="BoxItems.splice(i+1 ,0, AddBox())" class="add" title="增加">+</i>
                            <i @click="BoxItems.splice(i,1)" class="back" title="删除">-</i>
                        </div>
                    </li>
                </ul>
            </template>
            <div class="box-button">
                <div class="plus-btn" @click="BoxItems.push(AddBox()) ">+</div>
                <button class="button" type="submit">保存</button>
            </div>
        </form>
    </div>
</template>
<script>
export default {
    data: function () {
        return {
            Category: [],
            BoxItems: [],
        }
    },
    created: function () {
        this.$axios.all([
            this.$axios.get('/admin/box'),
            this.$axios.get('/admin/category')
        ]).then(this.$axios.spread((box, category) => {
            if (box.data.data != null) {
                this.BoxItems = box.data.data;
            }
            if (category.data.data != null) {
                this.Category = category.data.data
            }
        }));
    },
    mounted: function () {

    },
    methods: {
        AddBox: function () {
            return {
                Title: '',
                Icon: '',
                Url: '',
                Module: 0,
                Orderby: 0,
                Category: [],
                Widget: {
                    Enabled: 0,
                    Title: '',
                    Number: 0,
                    Orderby: 0,
                    Code: "",
                },
                Number: 0,
                Code: '',
            }
        },
        Update: function () {
            this.$axios.post('/admin/box', {
                data: JSON.stringify(this.BoxItems),
            }).then(res => {
                this.$store.commit("msg", res.data.message);
            });
        }
    }
}
</script>
<style scoped>
form.row {
    width: 100%;
    margin-bottom: 150px;
}
.box-button {
    display: flex;
    flex-wrap: nowrap;
    width: 100%;
    height: 25px;
    line-height: 25px;
}

.box-button > * {
    -ms-flex-positive: 1;
    -webkit-box-flex: 1;
    flex-grow: 1;
    margin: 0 10px;
}

.plus-btn {
    background: #00a206;
    cursor: pointer;
    color: #fff;
    border-radius: 4px;
    text-align: center;
    font-size: 2rem;
}

button.button {
    cursor: pointer;
    background: #008ec2;
    color: #fff;
    border: 0;
    border-radius: 4px;
}

.row ul.row-list {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 100%;
    border-bottom: solid 1px #86828224;
    border: 1px solid rgba(255, 255, 255, 0.29);
    box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
    border-radius: 1rem;
    margin-bottom: 2rem;
}
.row ul.row-list li.list-list {
    display: flex;
    flex-wrap: wrap;
    padding: 0 10px;
    margin-bottom: 6px;
    width: 100%;
    position: relative;
    min-height: 25px;
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
ul.tabs-panel {
    min-height: 42px;
    max-height: 200px;
    overflow: auto;
    border: solid 1px #ddd;
}
.list-c input,
.list-c select,
.list-c textarea {
    padding: 5px;
    border: 1px solid #ddd;
    background: rgba(255, 255, 255, 0);
    border: 1px solid rgb(191, 191, 191);
}
textarea.textarea.a {
    overflow: auto;
    padding: 2px 6px;
    line-height: 1.4;
    resize: vertical;
    height: 100px;
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
.list-c i:nth-child(2) {
    background: #f44336;
}
li.list-list i:hover,
.plus-btn:hover {
    opacity: 0.9;
}
.list-c i {
    -webkit-box-flex: 0;
    flex: 0 0 100%;
    max-width: 25%;
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
</style>
<style scoped>
@media (max-width: 600px) {
	.list-f {
		    max-width: 100%;
	}
}
</style>