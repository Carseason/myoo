<template>
	<div class="row">
		<form class="row" method="post" @submit.prevent="Update()">
			<template v-for=" (value,i) in Group">
				<ul class="row-list" :key="i">
					<h2 class="list-title">用户组 - {{i}}</h2>
					<li class="list-list">
						<div class="list-f">所属等级</div>
						<div class="list-c">
							<input class="input" min="1" required="required" type="number" placeholder="用户等级" v-model.number="value.Lv" />
						</div>
					</li>
					<li class="list-list">
						<div class="list-f">用户组名称</div>
						<div class="list-c">
							<input class="input" required="required" type="text" placeholder="用户组等级别名" v-model.trim="value.NickName" />
						</div>
					</li>
					<!-- <li class="list-list">
                        <div class="list-f">用户组勋章</div>
                        <div class="list-c"></div>
                        <div class="list-r">
                            <i class="add" title="增加" @click="value.Medal.push('')">+</i>
                        </div>
					</li>-->
					<!-- <template v-for="(item,j) in value.Medal">
                        <li class="list-list" :key="j">
                            <div class="list-f"></div>
                            <div class="list-c">
                                <input class="input" required="required" type="url" placeholder="勋章图标URL" v-model.trim="value.Medal[j]" />
                            </div>
                            <div class="list-r">
                                <i class="back" title="删除" @click="value.Medal.splice(j,1)">-</i>
                            </div>
                        </li>
					</template>-->

					<li class="list-list">
						<div class="list-f">文章投稿</div>
						<div class="list-c">
							<select class="select" v-model="value.NewPosts">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">文章评论</div>
						<div class="list-c">
							<select class="select" v-model="value.Comments">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">文章点赞</div>
						<div class="list-c">
							<select class="select" v-model="value.Supported">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">文章收藏</div>
						<div class="list-c">
							<select class="select" v-model="value.Favorite">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">用户关注</div>
						<div class="list-c">
							<select class="select" v-model="value.Followers">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">资源下载</div>
						<div class="list-c">
							<select class="select" v-model="value.Download">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">上传头像</div>
						<div class="list-c">
							<select class="select" v-model="value.UpdateAvatar">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">
							用户签到
							(暂不可用)
						</div>
						<div class="list-c">
							<select class="select" v-model="value.Checkin">
								<option :value="true">允许</option>
								<option :value="false">禁止</option>
							</select>
						</div>
					</li>

					<li class="list-list">
						<div class="list-f">
							分类权限
							(暂不可用)
						</div>
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
				</ul>
			</template>

			<div class="box-button">
				<div class="plus-btn" @click="Group.push(Add()) ">+</div>
				<button class="button" type="submit">保存</button>
			</div>
		</form>
	</div>
</template>
<script>
export default {
	data() {
		return {
			Category: [],
			Group: [],
		}
	},
	created: function () {
		this.$axios.all([
			this.$axios.get('/admin/category'),
			this.$axios.get('/admin/group')
		]).then(this.$axios.spread((category, group) => {
			if (category.data.data != null) {
				this.Category = category.data.data;
			}
			if (group.data.data != null) {
				this.Group = group.data.data;
			}
		}));
	},
	methods: {
		Add() {
			return {
				Lv: 0,
				Nickname: "",
				//    Medal: [],
				Category: [],
				NewPosts: true,
				Comments: true,
				Supported: true,
				Favorite: true,
				Followers: true,
				Download: true,
				Checkin: true,
				UpdateAvatar: true,
			}
		},
		Update() {
			this.$axios.post('/admin/group', {
				data: JSON.stringify(this.Group),
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
ul.row-list {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	border-bottom: solid 1px #86828224;
	border: 1px solid rgba(255, 255, 255, 0.29);
	box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
	border-radius: 1rem;
	margin-bottom: 2rem;
}
li.list-list {
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
input,
select,
textarea {
	padding: 5px 10px;
	background: rgba(255, 255, 255, 0);
	border: 1px solid rgb(191, 191, 191);
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
li.list-list i.back {
	background: #f44336;
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
li.list-list i:hover,
.plus-btn:hover {
	opacity: 0.9;
}
ul.tabs-panel {
	min-height: 42px;
	max-height: 200px;
	overflow: auto;
	border: solid 1px #ddd;
}
</style>
<style scoped>
@media (max-width: 600px) {
	.list-f {
		max-width: 100%;
	}
}
</style>