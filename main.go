package main

import (
	"myoo/jwt"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"myoo/router/v1/account"
	"myoo/router/v1/admin"
	"myoo/router/v1/auth"
	"myoo/router/v1/author"
	"myoo/router/v1/category"
	"myoo/router/v1/checkcode"
	"myoo/router/v1/comments"
	"myoo/router/v1/download"
	"myoo/router/v1/followers"
	"myoo/router/v1/inc"
	"myoo/router/v1/index"
	"myoo/router/v1/menus"
	"myoo/router/v1/posts"
	"myoo/router/v1/sign"

	"github.com/gin-gonic/gin"
)

func init() {
	conf := plugins.NewConfig()
	jwt.NewJwtKey(conf.GetKey())
	plugins.NewAesKey(conf.GetKey())
	models.RegisterDB(conf.GetDbAddr())
	redis.New(conf.GetRedisAddr(), conf.GetRedisPass())
}

func main() {
	/*******************Gin初始化*******************/
	r := gin.New()
	r.Use(FilterXsrf(), FilterJwt())
	/*******************v1路由组*******************/
	rest := r.Group("/rest/v1")
	{
		//	RouterCheckCode(rest) //验证码
		RouterAuth(rest)
		RouterInc(rest)       //信息
		RouterMenus(rest)     //菜单
		RouterIndex(rest)     //首页
		RouterFollowers(rest) //关注事件
		RouterCategory(rest)  //分类
		RouterAuthor(rest)    //用户空间
		RouterSign(rest)      //登录
		RouterAccount(rest)   //个人中心
		RouterPosts(rest)     //文章
		RouterDownload(rest)  //下载
		RouterComments(rest)  //评论
		RouterAdmin(rest)     //管理页面
	}

	/*******************模板输出*******************/
	if plugins.GetConfig().GetServerRender() {
		r.LoadHTMLGlob("views/*")
		r.NoRoute(RouterViews) //未知路由处理
	}
	/*******************静态文件夹*******************/
	r.Static("/src", "src")
	r.Run(plugins.GetConfig().GetPort())
}

/**************************************
				路由事件
**************************************/
/*******************验证码*******************/
func RouterCheckCode(rest *gin.RouterGroup) {
	rest.Use(FilterCheckCode())
	rest.GET("/checkcode", new(checkcode.CheckCode).Get)
	return
}

/*******************auth*******************/
func RouterAuth(rest *gin.RouterGroup) {
	rest.GET("/auth", new(auth.Auth).Get)
	return
}

/*******************信息*******************/
func RouterInc(rest *gin.RouterGroup) {
	rest.GET("/inc", new(inc.Inc).Get)
	return
}

/*******************菜单*******************/
func RouterMenus(rest *gin.RouterGroup) {
	rest.GET("/menus", new(menus.Menus).Get)
	return
}

/*******************首页*******************/
func RouterIndex(rest *gin.RouterGroup) {
	rest.GET("/index", new(index.Index).Get)
	return
}

/*******************关注*******************/
func RouterFollowers(rest *gin.RouterGroup) {
	rest.POST("/followers", new(followers.Followers).Post)
	return
}

/*******************分类*******************/
func RouterCategory(rest *gin.RouterGroup) {
	rest.GET("/category/:id/:page", new(category.Category).Get)
	return
}

/*******************空间*******************/
func RouterAuthor(rest *gin.RouterGroup) {
	router := rest.Group("/author")
	{
		router.GET("/:id", new(author.Author).Get)
		router.GET("/:id/posts/:page", new(author.Posts).Get)
		router.GET("/:id/followers/:page", new(author.Followers).Get)
		router.GET("/:id/fans/:page", new(author.Fans).Get)
	}
	return
}

/*******************登录*******************/
func RouterSign(rest *gin.RouterGroup) {
	router := rest.Group("/sign")
	{
		router.POST("/login", new(sign.Login).Post)
		router.GET("/registered", new(sign.Registered).Get)
		router.POST("/registered", new(sign.Registered).Post)
		router.POST("/forgetpwd", new(sign.Forgetpwd).Post)
		router.POST("/check/:path", new(sign.ResetPwd).Post)
	}
	return
}

/*******************文章*******************/
func RouterPosts(rest *gin.RouterGroup) {
	router := rest.Group("/posts")
	{
		router.GET("/:id", new(posts.Posts).Get)
		router.POST("/:id", new(posts.Posts).Post)
	}
	return
}

/*******************下载*******************/
func RouterDownload(rest *gin.RouterGroup) {
	router := rest.Group("/download")
	{
		router.GET("/:path", new(download.Download).Get)
	}
	return
}

/*******************评论*******************/
func RouterComments(rest *gin.RouterGroup) {
	router := rest.Group("/comments")
	{
		router.GET("/:id/:page", new(comments.Comments).Get)
		router.POST("/:id", new(comments.Comments).Post)
	}
	return
}

/*******************个人中心*******************/
func RouterAccount(rest *gin.RouterGroup) {
	router := rest.Group("/account")
	router.Use(FilterAuth())
	{
		router.GET("", new(account.Account).Get)
		router.POST("/info", new(account.Info).Post)
		router.GET("/newposts", new(account.NewPosts).Get)
		router.POST("/newposts", new(account.NewPosts).Post)
		router.GET("/posts/:page", new(account.Posts).Get)
		router.GET("/favorite/:page", new(account.Favorite).Get)
		router.GET("/comments/:page", new(account.Comments).Get)
	}
	return
}

/*******************管理*******************/
func RouterAdmin(rest *gin.RouterGroup) {
	router := rest.Group("/admin")
	router.Use(FilterAdmin())
	{
		router.GET("", new(admin.Admin).Get)
		router.GET("/statistics", new(admin.Statistics).Get)
		router.GET("/options", new(admin.Options).Get)
		router.POST("/options", new(admin.Options).Post)
		router.GET("/category", new(admin.Category).Get)
		router.POST("/category", new(admin.Category).Post)
		router.GET("/menus", new(admin.Menus).Get)
		router.POST("/menus", new(admin.Menus).Post)
		router.GET("/box", new(admin.Box).Get)
		router.POST("/box", new(admin.Box).Post)
		router.GET("/posts/:page", new(admin.Posts).Get)
		router.GET("/user/:id", new(admin.User).Get)
		router.POST("/user/:id", new(admin.User).Post)
		router.GET("/users/:page", new(admin.Users).Get)
		router.GET("/group", new(admin.Group).Get)
		router.POST("/group", new(admin.Group).Post)
		router.GET("/modifyposts/:id", new(admin.ModifyPosts).Get)
		router.POST("/modifyposts/:id", new(admin.ModifyPosts).Post)
		router.POST("/cache", new(admin.Cache).Post)
	}
	return
}

/*******************模板*******************/
func RouterViews(rest *gin.Context) {
	rest.HTML(200, "index.html", gin.H{
		"title": "Main website",
	})
	return
}

/**************************************
				中间件
**************************************/
/*******************Xssrf处理*******************/
func FilterXsrf() gin.HandlerFunc {
	return func(router *gin.Context) {
		router.Next()
	}
}

/*******************验证码处理*******************/
func FilterCheckCode() gin.HandlerFunc {
	return func(router *gin.Context) {
		code := router.PostForm("checkcode")
		key := router.PostForm("checkkey")
		if len(code) != 6 || !checkcode.VerfiyCaptcha(key, code) {
			router.JSON(200, gin.H{"success": 403, "message": "验证码错误", "data": nil})
			router.Abort()
			return
		}
		router.Next()
		return
	}
}

/*******************JWT处理*******************/
func FilterJwt() gin.HandlerFunc {
	return func(router *gin.Context) {
		if id, ok := jwt.Decode(router.GetHeader("X-MYOO")); ok {
			router.Set("auth", id)
		} else {
			router.Set("auth", nil)
		}
		router.Next()
	}
}

/*******************限制登录*******************/
func FilterAuth() gin.HandlerFunc {
	return func(router *gin.Context) {
		if id := router.GetInt64("auth"); id == 0 {
			router.JSON(403, gin.H{"success": 403, "message": "请先登录", "data": nil})
			router.Abort()
			return
		}
		router.Next()
	}
}

/*******************是否管理员*******************/
func FilterAdmin() gin.HandlerFunc {
	return func(router *gin.Context) {
		id := router.GetInt64("auth")
		if auth := redis.Query_User(id); auth.Admin {
			router.Next()
			return
		}
		router.JSON(403, gin.H{"success": 403, "message": "请先登录", "data": nil})
		router.Abort()
		return
	}
}
