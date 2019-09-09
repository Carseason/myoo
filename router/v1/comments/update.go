package comments

import (
	"html"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"strings"

	"github.com/gin-gonic/gin"
)

func (*Comments) Post(router *gin.Context) {
	id := plugins.StringtoInt64(router.Param("id"))
	if id == 0 {
		router.JSON(200, gin.H{"success": 403, "message": nil, "data": nil})
		return
	}

	uid := router.GetInt64("auth")
	if uid == 0 {
		router.JSON(200, gin.H{"success": 403, "message": "请先登录", "data": nil})
		return
	}
	auth := redis.Query_User(uid)
	if !plugins.Group_Comments(auth.Lv) {
		router.JSON(200, gin.H{"success": 403, "message": "当前用户组无法提交评论", "data": nil})
		return
	}
	content := router.PostForm("content")
	if content == "" {
		router.JSON(200, gin.H{"success": 403, "message": "说点儿什么?", "data": nil})
		return
	}
	DbComments := redis.GetOptionsComments()
	if DbComments.KeywordsFilter != "" && plugins.RegexpMatch(content, DbComments.KeywordsFilter) {
		router.JSON(200, gin.H{"success": 403, "message": "评论里存在敏感词汇,无法提交", "data": nil})
		return
	}
	posts := new(models.Posts)
	if !posts.Is(id) {
		router.JSON(200, gin.H{"success": 403, "message": "非法的评论", "data": nil})
		return
	}
	/*******评论关键词替换*******/
	if DbComments.SensitivewordBefore != "" {
		content = plugins.RegexpReplaceAllString(content, DbComments.SensitivewordBefore, DbComments.SensitivewordRear)
	}

	comments := new(models.Comments)

	if v := comments.Insert(&models.Comments{
		Posts: &models.Posts{
			Id: id,
		},
		Author: &models.User{
			Id: uid,
		},
		Content: ReplaceEmoji(content),
	}); v > 0 {
		defer redis.DelAll("Myoo/Comments/" + router.Param("id") + "/*")
		router.JSON(200, gin.H{"success": 200, "message": "评论成功", "data": v})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "评论失败", "data": nil})
	return
}

/****************评论过滤*******************/
func ReplaceEmoji(com string) string {
	com = html.EscapeString(com)
	com = strings.Replace(com, "[大笑]", `<img title="大笑" src="https://i.loli.net/2019/08/05/wbmvoD5ycFJ79KA.gif">`, -1)
	com = strings.Replace(com, "[喷血]", `<img title="喷血" src="https://i.loli.net/2019/08/05/EbY8FCLmKcQIWHz.gif">`, -1)
	com = strings.Replace(com, "[抠鼻]", `<img title="抠鼻" src="https://i.loli.net/2019/08/05/xfhvsjE284nPVip.gif">`, -1)
	com = strings.Replace(com, "[吐]", `<img title="吐" src="https://i.loli.net/2019/08/05/98nzCEk4pHriBV1.gif">`, -1)
	com = strings.Replace(com, "[微笑]", `<img  title="微笑" src="https://i.loli.net/2019/08/05/YoNItSivPL2ZJAn.gif">`, -1)
	com = strings.Replace(com, "[笑哭]", `<img  title="笑哭" src="https://i.loli.net/2019/08/05/PX6M1eFxrYVZNdC.gif">`, -1)
	com = strings.Replace(com, "[斜眼笑]", `<img  title="斜眼笑" src="https://i.loli.net/2019/08/05/vcDieIRrf25KlAz.gif">`, -1)
	com = strings.Replace(com, "[阴险]", `<img  title="阴险" src="https://i.loli.net/2019/08/05/tNqeoXlZcywU41R.gif">`, -1)
	com = strings.Replace(com, "[doge]", `<img  title="doge" src="https://i.loli.net/2019/08/05/sRAJywb56qphYK1.gif">`, -1)
	com = strings.Replace(com, "[狗头]", `<img  title="狗头" src="https://i.loli.net/2019/08/05/Ww6GxbhD1Vamq8A.gif">`, -1)
	com = strings.Replace(com, "[猪头]", `<img  title="猪头" src="https://i.loli.net/2019/08/05/jGbMDPfFphlrXuz.gif">`, -1)
	com = strings.Replace(com, "[马]", `<img  title="马" src="https://i.loli.net/2019/08/05/MnKwdZi4eBVj3LD.gif">`, -1)
	com = strings.Replace(com, "[牛]", `<img  title="牛" src="https://i.loli.net/2019/08/05/AguMSPy1DTN7BVt.gif">`, -1)
	com = strings.Replace(com, "[啤酒]", `<img  title="啤酒" src="https://i.loli.net/2019/08/05/q8i1MQUCE5YrOxF.gif">`, -1)
	return com
}
