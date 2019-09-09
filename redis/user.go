/*#Comment
 *@Description: Redis-User
 *@MethodAuthor:  Carseason
 *@Date: 2019-08-12 21:00:09
 */
package redis

import (
	"myoo/models"
	"strconv"
)

/**************int64 to string *******************/
func Int64ToString(id int64) string {
	return strconv.FormatInt(id, 10)
}

func Del_User(id int64) error {
	cachePath := "Myoo/User/" + Int64ToString(id)
	return Del(cachePath)
}

func Query_User(id int64) *models.User {
	cachePath := "Myoo/User/" + Int64ToString(id)
	user := models.NewUser()
	if Is(cachePath) {
		GetStruct(cachePath, user)
	} else {
		user = user.Query_One(id)
		PutStruct(cachePath, user, 3600)
	}
	return user
}
