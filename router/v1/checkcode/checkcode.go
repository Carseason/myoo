package checkcode

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

type CheckCode struct{}

type CheckCodeer struct {
	Key      string
	BaseCode string
}

func (*CheckCode) Get(router *gin.Context) {
	code := new(CheckCodeer)
	code.Key, code.BaseCode = CodeCaptchaCreate()
	router.JSON(200, gin.H{"success": 200, "message": nil, "data": code})
	return
}

func CodeCaptchaCreate() (string, string) {
	//创建数字验证码.
	// idKey, capD := base64Captcha.GenerateCaptcha("", base64Captcha.ConfigDigit{
	// 	Height:     80,
	// 	Width:      240,
	// 	MaxSkew:    0.7,
	// 	DotCount:   80,
	// 	CaptchaLen: 5,
	// })
	//创建字符串验证码.
	idKey, capD := base64Captcha.GenerateCaptcha("", base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeNumberAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	})

	//以base64编码
	base64stringD := base64Captcha.CaptchaWriteToBase64Encoding(capD)
	return idKey, base64stringD
}

func VerfiyCaptcha(idkey, verifyValue string) bool {
	return base64Captcha.VerifyCaptcha(idkey, verifyValue)
}
