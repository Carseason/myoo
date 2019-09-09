package sign

import (
	"encoding/json"
	"myoo/models"
	"myoo/plugins"
	"myoo/redis"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
)

type Forgetpwd struct{}

func (this *Forgetpwd) Post(router *gin.Context) {
	email := router.PostForm("email")
	if !plugins.An_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "邮箱错误", "data": nil})
		return
	}
	user := new(models.User)
	if !user.Is_Email(email) {
		router.JSON(200, gin.H{"success": 403, "message": "该邮箱尚未注册", "data": nil})
		return
	}
	if this.email(router) {
		router.JSON(200, gin.H{"success": 200, "message": "已向该邮箱发送验证邮件...", "data": nil})
		return
	}
	router.JSON(200, gin.H{"success": 403, "message": "邮件发送失败...", "data": nil})
	return
}

func (*Forgetpwd) email(router *gin.Context) bool {
	email := router.PostForm("email")
	optionsEmail := redis.GetOptionsEmail()

	user := new(models.User).Query_Email(email)
	sitename := redis.GetOptionsSitename()
	date := (time.Now().Unix())
	fo := ResetPwder{
		Id:         user.Id,
		IssuedTime: date,
		OutTime:    date + 600, //十分钟过期时间
	}
	jsonData, err := json.Marshal(fo)
	if err != nil {
		return false
	}
	token := optionsEmail.Url + "/sign/resetpwd/" + plugins.Aes_Encode(string(jsonData))
	links := `<a href="` + token + `">` + token + `</a>`
	conntent := `
		Hi	` + user.Name + `,
		<br>
		您正在修改密码,请点击下方链接去修改:
		<br>
		` + links + `
		<br>
		<br>
		如果这个请求不是由你发起的，那没问题，你不用担心，你可以安全地忽略这封邮件。
		<br>
		该链接有效期为 10 分钟。
		<br>
		` + sitename + `
		<br>
		<br>
		p.s.作为安全备注，这次密码找回请求是由 IP 地址 ` + router.ClientIP() + ` 使用浏览器 ` + router.GetHeader("User-Agent") + ` 在 ` + time.Now().Format("2006-01-02 15:04:05") + ` 发起的
	`
	m := gomail.NewMessage()
	m.SetAddressHeader("From", optionsEmail.FromEmail, optionsEmail.FromUname)                                                         // 发件人
	m.SetHeader("To", m.FormatAddress(user.Email, ""))                                                                                 // 收件人
	m.SetHeader("Subject", sitename+" 重置密码")                                                                                           // 主题
	m.SetBody("text/html", conntent)                                                                                                   // 正文
	d := gomail.NewPlainDialer(optionsEmail.EmailHost, optionsEmail.EmailPort, optionsEmail.EmailUsername, optionsEmail.EmailPassword) // 发送邮件服务器、端口、发件人账号、发件人密码
	return d.DialAndSend(m) == nil
}
