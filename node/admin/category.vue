<template>
	<div class="row">
		<ul class="row-list">
			<h2 class="list-title">分类配置</h2>
			<li class="list-list">
				<div class="list-f"></div>
				<div class="list-c">
					<input class="input a" type="text" placeholder="分类名称" v-model.trim="CategoryName" />
				</div>
				<div class="list-r">
					<button class="button1" @click="NewCategory()">新增分类</button>
				</div>
			</li>
		</ul>
		<ul class="row-list">
			<h2 class="list-title">分类管理</h2>
			<template v-for="(value,index) in CategoryAll">
				<li class="list-list" :key="value.Id">
					<div class="list-c">
						<input class="input c" type="text" placeholder="分类名称" v-model.trim="value.Name" />
						<input class="input c" type="text" placeholder="分类ID" disabled="disabled" :value="value.Id" style="background: rgba(175, 175, 175, 0.48);" />
						<div style="margin:0 0 0 auto;text-align: right;">
							<a class="button-a" @click="Update(index,'update')">更新</a>
							<a class="button-a" @click="Update(index,'delete')">删除</a>
							<a class="button-a" :href="'/category/' + value.Id" target="_blank">查看</a>
						</div>
					</div>
				</li>
			</template>
		</ul>
	</div>
</template>

<script>
export default {
	data: function () {
		return {
			CategoryName: "",
			Category: { Id: null, Name: null },
			CategoryAll: []
		};
	},
	mounted: function () { },
	created: function () {
		this.$axios.get("/admin/category").then(res => {
			if (res.data.data != null) {
				this.CategoryAll = res.data.data;
			}
		}).catch(error => {
			console.log(error);
		});
	},
	methods: {
		NewCategory: function () {
			if (this.CategoryName.length == 0) {
				this.$store.commit("msg", "分类名称不能为空");
				return;
			}
			this.$axios
				.post("/admin/category", {
					type: "insert",
					name: this.CategoryName
				})
				.then(res => {
					this.$store.commit("msg", res.data.message);
				})
				.catch(error => {
					console.log(error);
				});
			return;
		},
		Update: function (index, types) {
			this.$axios
				.post("/admin/category", {
					type: types,
					id: this.CategoryAll[index].Id,
					name: this.CategoryAll[index].Name
				})
				.then(res => {
					this.$store.commit("msg", res.data.message);
				})
				.catch(error => {
					console.log(error);
				});
			return;
		}
	}
};
</script>

<style scoped>
ul.row-list {
	border: 1px solid rgba(255, 255, 255, 0.29);
	box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
	border-radius: 1rem;
	margin: 10px 0;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
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

.list-r {
	flex: 0 0 100%;
	max-width: 90px;
	overflow: hidden;
	text-align: left;
}

button.button1 {
	height: 100%;
	cursor: pointer;
	background: #008ec2;
	color: #fff;
	border: 0;
	border-radius: 4px;
	margin-left: 5px;
	padding: 5px 10px;
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
.list-c > * {
	-ms-flex-positive: 1;
	-webkit-box-flex: 1;
	flex-grow: 1;
	margin: 0 5px;
}

a.button-a {
	padding: 0 10px;
	display: inline-block;
	height: 100%;
	background: #fff;
	border-radius: 2px;
}
a.button-a:hover {
	background: #008ec2;
	color: #fff;
}
</style>
<style scoped>
@media (max-width: 600px) {
	.list-f {
		max-width: 100%;
	}
}
</style>