package request

type Email struct {
	To      string `json:"to"`      // 邮件的收件地址
	Subject string `json:"subject"` // 邮件主题
	Body    string `json:"body"`    // 邮件内容
}
