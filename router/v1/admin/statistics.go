package admin

import (
	"myoo/models"

	"github.com/gin-gonic/gin"
)

type Statistics struct{}

type Statisticser struct {
	User     statisticsUser
	Posts    statisticsPosts
	Comments statisticsComments
}

type statisticsUser struct {
	AllRegistered   int64
	TodayRegistered int64
	WeekRegistered  int64
	MonthRegistered int64
	YearRegistered  int64
}

type statisticsPosts struct {
	AllPosts   int64
	TodayPosts int64
	WeekPosts  int64
	MonthPosts int64
	YearPosts  int64
}

type statisticsComments struct {
	AllComments   int64
	TodayComments int64
	WeekComments  int64
	MonthComments int64
	YearComments  int64
}

func (*Statistics) Get(router *gin.Context) {
	template := new(Statisticser)
	posts := new(models.Posts)
	users := new(models.User)
	comments := new(models.Comments)
	template.User = statisticsUser{
		AllRegistered:   users.Count(),
		TodayRegistered: users.Count_Date("day"),
		WeekRegistered:  users.Count_Date("week"),
		MonthRegistered: users.Count_Date("month"),
		YearRegistered:  users.Count_Date("year"),
	}
	template.Posts = statisticsPosts{
		AllPosts:   posts.Count(),
		TodayPosts: posts.Count_Date("day"),
		WeekPosts:  posts.Count_Date("week"),
		MonthPosts: posts.Count_Date("month"),
		YearPosts:  posts.Count_Date("year"),
	}

	template.Comments = statisticsComments{
		AllComments:   comments.Count_All(),
		TodayComments: comments.Count_Date("day"),
		WeekComments:  comments.Count_Date("week"),
		MonthComments: comments.Count_Date("month"),
		YearComments:  comments.Count_Date("year"),
	}
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": template})
	return
}
