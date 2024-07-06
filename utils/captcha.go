package utils

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

var captchaStore = base64Captcha.DefaultMemStore

// CreateCaptcha 生成图形验证码
func CreateCaptcha(captchaType string) (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	var (
		width           int         = 240                                         // 图形宽度
		height          int         = 80                                          // 图形高度
		length          int         = 6                                           // 数字、字符生成长度（默认6位）
		chineseLength   int         = 2                                           // 中文生成长度（默认2位）
		dotCount        int         = 80                                          // 随机干扰点
		maxSkew         float64     = 0.7                                         // 最大倾斜度
		noiseCount      int         = 0                                           // 噪点数量（0-无噪点|10-少量噪点|50-大量噪点）
		showLineOptions int         = base64Captcha.OptionShowHollowLine          // 干扰线条类型（OptionShowHollowLine-空心线图|OptionShowSlimeLine-细线条|OptionShowSineLine-正弦线条）
		bgColor         *color.RGBA = &color.RGBA{R: 242, G: 242, B: 242, A: 255} // 背景颜色
		stringSource    string      = base64Captcha.TxtAlphabet                   // 字符串字符源
		chineseSource   string      = base64Captcha.TxtChineseCharaters           // 中文字符源
		fonts           []string    = []string{"wqy-microhei.ttc"}                // 文泉驿微米黑，开源中文字体，适用于中文简体和繁体字符、日文汉字、韩文字母等
	)

	switch captchaType {
	case "digit":
		driver = base64Captcha.NewDriverDigit(height, width, length, maxSkew, dotCount)
	case "string":
		driver = base64Captcha.NewDriverString(height, width, noiseCount, showLineOptions, length, stringSource, bgColor, nil, fonts)
	case "math":
		driver = base64Captcha.NewDriverMath(height, width, noiseCount, showLineOptions, bgColor, nil, fonts)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(height, width, noiseCount, showLineOptions, chineseLength, chineseSource, bgColor, nil, fonts)
	default:
		driver = base64Captcha.NewDriverDigit(height, width, length, maxSkew, dotCount)
	}

	cp := base64Captcha.NewCaptcha(driver, captchaStore) // 创建验证码对象
	id, b64s, _, err = cp.Generate()                     // 生成验证码图像及其对应的标识（b64s是图片的base64编码）

	return id, b64s, err
}

// VerifyCaptcha 校验图形验证码
func VerifyCaptcha(captchaId string, captcha string) bool {
	return captchaStore.Verify(captchaId, captcha, true)
}
