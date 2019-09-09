package models

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func RegexpMatch(s1, s2 string) bool {
	if r1, err := regexp.MatchString(s1, s2); err == nil { //s1 为要匹配的规则, s2为内容
		return r1
	}
	return false
}

func RegexpSubmatch(s1, s2 string) string {
	r1 := regexp.MustCompile(s1).FindSubmatch([]byte(s2))
	if len(r1) == 0 {
		return ""
	}
	return string(r1[1])
}

func PwdSet(s, s1 string) bool { //密码比较
	if err := bcrypt.CompareHashAndPassword([]byte(s1), []byte(s)); err == nil { //s原密码,S1为数据库密码
		return true
	}
	return false
}

func PwdGet(s string) (string, error) { //密码生成
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost); err == nil { //原文,加密的等级
		return string(hashedPassword), nil
	} else {
		return "", err
	}

}
