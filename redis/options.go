/*#Comment
 *@Description: 网站基本信息缓存
 *@MethodAuthor:  Carseason
 *@Date: 2019-08-12 19:38:21
 */
package redis

import (
	"encoding/json"
	"myoo/models"
	"strconv"
)

type Options struct {
	Sitename        string
	Sitedescription string
	Registered      bool
	Download        bool
	Search          string
	Email           Email
	Comments        Comments
	DefaultCategory int64
	Ad              Ad
	Message         string
	Code            Code
	DefaultLv       int64
}

type Email struct {
	Url           string
	FromEmail     string
	FromUname     string
	EmailUsername string
	EmailPassword string
	EmailHost     string
	EmailPort     int
}
type Comments struct {
	KeywordsFilter      string
	SensitivewordBefore string
	SensitivewordRear   string
}
type Ad struct {
	Header string
	Footer string
}
type Code struct {
	Style  string
	Script string
}

func GetOptionsValue(key string) string {
	path := "Myoo/Options/" + key
	if Is(path) {
		return GetString(path)
	} else {
		options := new(models.Options)
		MapData := options.Query(key)
		Put(path, MapData[key], 7200)
		return MapData[key]
	}
	return ""
}
func GetDefaultLv() int64 {
	var lv int64
	value := GetOptionsValue("DefaultLv")
	if json.Unmarshal([]byte(value), &lv) != nil {
		return lv
	}
	return 1
}

func GetOptionsSitename() string {
	return GetOptionsValue("Sitename")
}
func GetOptionsSitedescription() string {
	return GetOptionsValue("Sitedescription")
}
func GetOptionsRegistered() bool {
	value := GetOptionsValue("Registered")
	return value == "true" || value == ""
}
func GetOptionsDownload() bool {
	value := GetOptionsValue("Download")
	return value == "true"
}
func GetOptionsSearch() string {
	return GetOptionsValue("Search")
}
func GetOptionsMobileSearch() string {
	return GetOptionsValue("MobileSearch")
}
func GetOptionsMessage() string {
	return GetOptionsValue("Message")
}

func GetOptionsDefaultCategory() int64 {
	value, err := strconv.Atoi(GetOptionsValue("DefaultCategory"))
	if err != nil {
		return 0
	}
	return int64(value)
}

func GetOptionsEmail() *Email {
	this := new(Email)
	value := GetOptionsValue("Email")
	json.Unmarshal([]byte(value), this)
	return this
}

func GetOptionsComments() *Comments {
	this := new(Comments)
	value := GetOptionsValue("Comments")
	json.Unmarshal([]byte(value), this)
	return this
}

func GetOptionsAd() Ad {
	this := Ad{}
	value := GetOptionsValue("Ad")
	json.Unmarshal([]byte(value), &this)
	return this
}

func GetOptionsCode() *Code {
	this := new(Code)
	value := GetOptionsValue("Code")
	json.Unmarshal([]byte(value), this)
	return this
}
