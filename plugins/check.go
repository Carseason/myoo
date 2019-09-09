package plugins

import (
	"regexp"
	"strings"
)

/******************正则验证*****************/
func RegexpMatch(content, rule string) bool {
	if r1, err := regexp.MatchString(rule, content); err == nil {
		return r1
	}
	return false
}

/******************正则字符串查找字符串,返回数组*****************/
func RegexpFindAll(content, rule string) []string { //Regexp0
	r1 := regexp.MustCompile(rule)
	r2 := r1.FindAllString(content, -1)
	return r2
}

/******************正则查找匹配内容*****************/
func RegexpSubmatch(content []byte, rule string) []byte { //Regexp3,s2为正文，s1为规则
	r1 := regexp.MustCompile(rule).FindSubmatch(content)
	if len(r1) == 0 {
		return nil
	}
	return r1[1]
}

/******************正则查找内容，返回第一个匹配到的*****************/
func RegexFindString(content, rule string) string { //Regexp2
	r1 := regexp.MustCompile(rule)
	r2 := r1.FindString(content)
	return r2
}

/******************正则替换*****************/
func RegexpReplaceAllString(content, rule, key string) string { //Regexp1
	r1 := regexp.MustCompile(rule)
	r2 := r1.ReplaceAllString(content, key)
	return r2
}
func Replace(content, rule string) string {
	return strings.Replace(content, rule, "", -1)
}

/**************验证用户名正则**************/
func An_Name(name string) bool {
	return RegexpMatch(`(\n|\s|\r|\f|\t|@|'|")`, name)
}

/**************验证用户密码正则**************/
func An_Pwd(pwd string) bool {
	return RegexpMatch(pwd, `[^0-9a-zA-Z\._]`)
}

/**************验证用户邮箱正则**************/
func An_Email(email string) bool {
	return RegexpMatch(email, `^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,6}$`)
}

/**************验证是否图像URL**************/
func An_Image(img string) bool {
	return RegexpMatch(img, `^https?:\/\/([^\/]+\/){1,}[^\.]+\.(jpg|png|jpge|jpeg|gif|webp)$`)
}
