<template>
    <div class="body">
        <form class="row" method="post" @submit.prevent="Update()">
            <ul class="row-list">
                <h2 class="list-title">网站配置</h2>
                <li class="list-list">
                    <div class="list-f">站点标题</div>
                    <div class="list-c">
                        <input class="input b" type="text" placeholder="标题" v-model.trim="Options.Sitename" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">副标题</div>
                    <div class="list-c">
                        <input class="input b" type="text" placeholder="副标题" v-model.trim="Options.Sitedescription" />
                    </div>
                </li>
            </ul>
            <ul class="row-list">
                <h2 class="list-title">占位图片</h2>
                <li class="list-list">
                    <div class="list-f">站点图标</div>
                    <div class="list-c">前往 /src/img/ 文件夹修改 icon.png 图像文件</div>
                    <div class="list-r">
                        <img src="/src/img/icon.png" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">站点LOGO</div>
                    <div class="list-c">前往 /src/img/ 文件夹修改 logo.png 图像文件</div>
                    <div class="list-r">
                        <img src="/src/img/logo.png" />
                    </div>
                </li>

                <li class="list-list">
                    <div class="list-f">默认特色图</div>
                    <div class="list-c">前往 /src/img/ 文件夹修改 default.jpg 图像文件</div>
                    <div class="list-r">
                        <img src="/src/img/default.jpg" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">懒加载图片</div>
                    <div class="list-c">前往 /src/img/ 文件夹修改 lazyload.jpg 图像文件</div>
                    <div class="list-r">
                        <img src="/src/img/lazyload.jpg" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">默认头像</div>
                    <div class="list-c">前往 /src/avatar/ 文件夹修改 default.jpg 图像文件</div>
                    <div class="list-r">
                        <img src="/src/avatar/default.jpg" />
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">注册</h2>
                <li class="list-list">
                    <div class="list-f">注册功能</div>
                    <div class="list-c">
                        <select class="select a" v-model="Options.Registered">
                            <option :value="true">开启</option>
                            <option :value="false">关闭</option>
                        </select>
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">默认等级</div>
                    <div class="list-c">
                        <input type="number" name="defaultLv" placeholder="默认用户等级" min="1" v-model.number="Options.DefaultLv" required="required" />
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">下载</h2>
                <li class="list-list">
                    <div class="list-f">下载权限</div>
                    <div class="list-c">
                        <select class="select a" v-model="Options.Download">
                            <option :value="false">所有人</option>
                            <option :value="true">用户组</option>
                        </select>
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">搜索配置</h2>
                <li class="list-list">
                    <div class="list-f"></div>
                    <div class="list-c">
                        <code style="font-size:1.5rem">使用关键词 @Myoo@ 替换</code>
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">搜索URL地址</div>
                    <div class="list-c">
                        <input type="url" placeholder="使用第三方索引URL" class="input a" v-model.trim="Options.Search" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">移动端搜索地址</div>
                    <div class="list-c">
                        <input type="url" placeholder="使用第三方索引URL" class="input a" v-model.trim="Options.MobileSearch" />
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">邮箱配置</h2>
                <li class="list-list">
                    <div class="list-f">站点链接</div>
                    <div class="list-c">
                        <input type="url" placeholder="例如：https://www.google.com 结尾不要带/" class="input a" v-model.trim="Options.Email.Url" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">发件邮箱</div>
                    <div class="list-c">
                        <input type="email" placeholder="例如：admin@admin.com" class="input a" v-model.trim="Options.Email.FromEmail" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">发件人名称</div>
                    <div class="list-c">
                        <input type="text" placeholder="例如：当前网站名称" class="input a" v-model.trim="Options.Email.FromUname" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">SMTP主机地址</div>
                    <div class="list-c">
                        <input type="text" placeholder="例如：smtp.gmail.com" class="input a" v-model.trim="Options.Email.EmailHost" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">端口</div>
                    <div class="list-c">
                        <input type="number" min="0" placeholder="例如：25,ssl加密为465" class="input a" v-model.number="Options.Email.EmailPort" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">SMTP主机账号</div>
                    <div class="list-c">
                        <input type="email" placeholder="SMTP账号" class="input a" v-model.trim="Options.Email.EmailUsername" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">SMTP主机密码</div>
                    <div class="list-c">
                        <input type="password" placeholder="SMTP密码" class="input a" v-model.trim="Options.Email.EmailPassword" />
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">评论配置</h2>
                <li class="list-list">
                    <div class="list-f">关键词审核</div>
                    <div class="list-c">
                        <input type="text" placeholder="正则规则,匹配到的关键词将无法发布" class="input a" v-model.trim="Options.Comments.KeywordsFilter" />
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">关键词过滤</div>
                    <div class="list-c">
                        <input type="text" placeholder="正则匹配" class="input b" v-model.trim="Options.Comments.SensitivewordBefore" />
                        <input type="text" placeholder="替换后的敏感词" class="input b" v-model.trim="Options.Comments.SensitivewordRear" />
                    </div>
                </li>
            </ul>

            <ul class="row-list">
                <h2 class="list-title">广告代码</h2>
                <li class="list-list">
                    <div class="list-f">导航HTML代码</div>
                    <div class="list-c">
                        <textarea placeholder="此代码将显示在导航下方" class="textarea a" v-model.trim="Options.Ad.Header"></textarea>
                    </div>
                </li>
                <li class="list-list">
                    <div class="list-f">底部HTML代码</div>
                    <div class="list-c">
                        <textarea placeholder="此代码将显示在底部友情链接上方" class="textarea a" v-model.trim="Options.Ad.Footer"></textarea>
                    </div>
                </li>
            </ul>
            <ul class="row-list">
                <h2 class="list-title">公告配置</h2>
                <li class="list-list">
                    <div class="list-f">全站HTML代码公告</div>
                    <div class="list-c">
                        <textarea placeholder="此代码将以弹窗显示" class="textarea a" v-model.trim="Options.Message"></textarea>
                    </div>
                </li>
            </ul>

            <!-- <ul class="row-list">
                <h2 class="list-title">自定义代码</h2>
                <li class="list-list">
                    <div class="list-f"></div>
                    <div class="list-c">
                        <input type="url" placeholder="Style资源URL引用" class="input a" v-model.trim="Options.Code.Style">
                    </div>
                </li>

                <li class="list-list">
                    <div class="list-f"></div>
                    <div class="list-c">
                        <textarea placeholder="自定义脚本代码代码" class="textarea a" v-model.trim="Options.Code.Script"></textarea>
                    </div>
                </li>
            
            </ul>-->
            <button class="button" type="submit">保存</button>
        </form>
    </div>
</template>
<script>
export default {
    data() {
        return {
            Options: {
                Sitename: "",
                DefaultLv: 1,
                Sitedescription: "",
                Registered: false,   //注册功能
                Download: false,
                Search: "",
                MobileSearch: "",
                Email: {
                    Url: "",
                    FromEmail: "",
                    FromUname: "",
                    EmailUsername: "",
                    EmailPassword: "",
                    EmailHost: "",
                    EmailPort: 0,
                },
                Comments: {
                    KeywordsFilter: "",
                    SensitivewordBefore: "",
                    SensitivewordRear: "",
                },
                DefaultCategory: 0,
                Ad: {
                    Header: "",
                    Footer: "",
                },
                Message: "",
                Code: {
                    Style: "",
                    Script: "",
                }
            },
        }
    },
    created() {
        this.$axios.all([
            this.$axios.get('/admin/options'),
        ]).then(this.$axios.spread((options) => {
            if (options.data.data != null) {
                this.Options = options.data.data;
            }
        }));
    },
    methods: {
        Update() {
            this.$store.commit("load", true);
            this.$axios.post('/admin/options', {
                data: JSON.stringify(this.Options)
            }).then(res => {
                this.$store.commit("load", false);
                this.$store.commit("msg", res.data.message)
            }).catch(error => {
                console.log(error);
            });
            return;
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
    max-width: 90px;
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
</style>
