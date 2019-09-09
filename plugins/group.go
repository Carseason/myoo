/**************************************************
@DCommit
@Description: 用户组功能
@MethodAuthor:  Carseason
@Date: 2019-09-04 21:59:44*
**************************************************/
package plugins

import (
	"encoding/json"
	"myoo/models"
	"myoo/redis"
)

type Group struct {
	Lv       int64
	NickName string
	//	Medal        []string
	Category     []int64
	NewPosts     bool
	Comments     bool
	Supported    bool
	Favorite     bool
	Followers    bool
	Download     bool
	Checkin      bool
	UpdateAvatar bool
}

func CacheGroup() map[int64]Group {
	cachePath := "Myoo/Group"
	groups := make(map[int64]Group)
	if redis.Is(cachePath) {
		redis.GetStruct(cachePath, &groups)

	} else {
		template := []Group{}
		options := new(models.Options)
		db := options.Query("Group")
		json.Unmarshal([]byte(db["Group"]), &template)
		for v := range template {
			groups[template[v].Lv] = template[v]
		}
		redis.PutStruct(cachePath, groups, 3600)
	}
	return groups
}
func CheckGroup(lv int64) Group {
	groups := CacheGroup()
	if value, ok := groups[lv]; ok {
		return value
	}
	return Group{}
}

/***********关注权限**************/
func Group_Followers(lv int64) bool {
	return CheckGroup(lv).Followers
}

/***********投稿权限**************/
func Group_NewPosts(lv int64) bool {
	return CheckGroup(lv).NewPosts
}

/***********评论权限**************/
func Group_Comments(lv int64) bool {
	return CheckGroup(lv).Comments
}

/***********点赞权限**************/
func Group_Supported(lv int64) bool {
	return CheckGroup(lv).Supported
}

/***********收藏权限**************/
func Group_Favorite(lv int64) bool {
	return CheckGroup(lv).Favorite
}

/***********下载权限**************/
func Group_Download(lv int64) bool {
	return CheckGroup(lv).Download
}

/***********签到权限**************/
func Group_Checkin(lv int64) bool {
	return CheckGroup(lv).Checkin
}

/**********更新头像权限*************/
func Group_UpdateAvatar(lv int64) bool {
	return CheckGroup(lv).UpdateAvatar
}

/***********返回等级**************/
func Group_Lv(lv int64) int64 {
	return CheckGroup(lv).Lv
}

/***********返回别名**************/
func Group_NickName(lv int64) string {
	return CheckGroup(lv).NickName
}

/***********返回勋章**************/
// func Group_Medal(lv int64) []string {
// 	return CheckGroup(lv).Medal
// }

/***********返回**************/
func Group_Category(lv int64) []int64 {
	return CheckGroup(lv).Category
}

/***********分类权限**************/
func Group_Category_Check(lv, id int64) bool {
	value := CheckGroup(lv).Category
	for v := range value {
		if value[v] == id {
			return true
		}
	}
	return false
}
