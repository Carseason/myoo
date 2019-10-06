<template>
	<div class="row">
		<ul class="statistics-item">
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.User.TodayRegistered,Statistics.User.AllRegistered)  + '%' }"></div>
					<span>最近注册: {{ Statistics.User.TodayRegistered }}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.User.WeekRegistered , Statistics.User.AllRegistered )  + '%' }"></div>
					<span>最近7天注册: {{ Statistics.User.WeekRegistered }}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.User.MonthRegistered , Statistics.User.AllRegistered)  + '%' }"></div>
					<span>最近30天注册: {{ Statistics.User.MonthRegistered }}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.User.YearRegistered ,Statistics.User.AllRegistered) + '%' }"></div>
					<span>最近365天注册: {{ Statistics.User.YearRegistered }}</span>
				</div>
			</li>

			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Posts.TodayPosts , Statistics.Posts.AllPosts)  + '%' }"></div>
					<span>最近投稿:{{Statistics.Posts.TodayPosts}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Posts.WeekPosts, Statistics.Posts.AllPosts)  + '%' }"></div>
					<span>最近7天投稿:{{Statistics.Posts.WeekPosts}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Posts.MonthPosts, Statistics.Posts.AllPosts)  + '%' }"></div>
					<span>最近30天投稿:{{Statistics.Posts.MonthPosts}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Posts.YearPosts, Statistics.Posts.AllPosts)  + '%' }"></div>
					<span>最近365天投稿:{{Statistics.Posts.YearPosts}}</span>
				</div>
			</li>

			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Comments.TodayComments , Statistics.Comments.AllComments)  + '%' }"></div>
					<span>最近评论:{{Statistics.Comments.TodayComments}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count( Statistics.Comments.WeekComments , Statistics.Comments.AllComments)  + '%' }"></div>
					<span>最近7天评论:{{Statistics.Comments.WeekComments}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count(Statistics.Comments.MonthComments , Statistics.Comments.AllComments)  + '%' }"></div>
					<span>最近30天评论:{{Statistics.Comments.MonthComments}}</span>
				</div>
			</li>
			<li class="statistics-list">
				<div class="statistics-schedule">
					<div :style="{ width : Count(Statistics.Comments.YearComments , Statistics.Comments.AllComments)  + '%' }"></div>
					<span>最近365天评论:{{Statistics.Comments.YearComments}}</span>
				</div>
			</li>
		</ul>
	</div>
</template>
                    
                    
                    
                    
                    
<script>
export default {
	data() {
		return {
			Statistics: {
				User: {
					AllRegistered: 0,
					TodayRegistered: 0,
					WeekRegistered: 0,
					MonthRegistered: 0,
					YearRegistered: 0,
				},
				Posts: {
					AllPosts: 0,
					TodayPosts: 0,
					WeekPosts: 0,
					MonthPosts: 0,
					YearPosts: 0,
				},
				Comments: {
					AllComments: 0,
					TodayComments: 0,
					WeekComments: 0,
					MonthComments: 0,
					YearComments: 0,
				}
			}
		}
	},
	mounted() {
	},
	methods: {
		Count(val, number) {
			return ((val / number) * 100)
		}

	},
	created() {
		this.$axios.get('/admin/statistics').then(res => {
			if (res.data.data != null) {
				this.Statistics = res.data.data
			}
		}).catch(error => {
			console.log(error);
		});
	},
}
</script>
<style scoped>
ul.statistics-item {
	margin: auto;
	border: 1px solid rgba(255, 255, 255, 0.29);
	box-shadow: 0 1px 5px rgba(111, 111, 111, 0.45);
	border-radius: 1rem;
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
}
li.statistics-list {
	padding: 0 20px;
	margin: 20px 0;
	height: 20px;
	line-height: 20px;
	width: 100%;
	display: flex;
	flex-wrap: wrap;
}

.statistics-schedule {
	-webkit-box-flex: 0;
	flex: 0 0 100%;
	max-width: 100%;
	box-shadow: 0 1px 5px rgb(161, 161, 161);
	border-radius: 1rem;
	position: relative;
	cursor: pointer;
}
.statistics-schedule > div {
	background: #44a7df;
	width: 0;
	height: 100%;
	border-radius: 1rem;
}
.statistics-schedule > span {
	position: absolute;
	top: 0;
	width: 100%;
	text-align: center;
	color: #ffffff;
	font-size: 12px;
}
</style>
