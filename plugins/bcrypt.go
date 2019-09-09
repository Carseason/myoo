package plugins

import (
	"golang.org/x/crypto/bcrypt"
)

func BcryotSet(s, s1 string) bool { //密码比较
	if err := bcrypt.CompareHashAndPassword([]byte(s1), []byte(s)); err == nil { //s原密码,S1为数据库密码
		return true
	}
	return false
}

func BcryotPut(s string) string { //密码生成
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost); err == nil { //原文,加密的等级
		return string(hashedPassword)
	} else {
		return ""
	}

}
