package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"github.com/wangyupo/GGB/enums"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system/request"
	systemResponse "github.com/wangyupo/GGB/model/system/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"image/color"
	"time"
)

type SysLoginApi struct{}

var captchaStore = base64Captcha.DefaultMemStore

// 写入登入/登出日志
func setLoginLog(c *gin.Context, userId uint, loginType enums.LoginType) {
	// 写入登录日志
	clientIP := c.ClientIP()               // 获取客户端IP
	userAgent := c.GetHeader("User-Agent") // 获取浏览器信息

	loginLog := log.SysLogLogin{
		UserId:    userId,
		Type:      loginType,
		IP:        clientIP,
		UserAgent: userAgent,
	}

	err := global.GGB_DB.Create(&loginLog).Error
	if err != nil {
		global.GGB_LOG.Error("写入登录日志失败！", zap.Error(err))
	}
}

// Login 登录
func (s *SysUserApi) Login(c *gin.Context) {
	// 声明 loginForm 类型的变量以存储 JSON 数据
	var loginForm request.Login
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		// 错误处理
		utils.HandleValidatorError(err, c)
		return
	}

	user, err := sysLoginService.Login(loginForm)
	if err != nil {
		global.GGB_LOG.Error("登录失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 通过jwt生成token
	claims := utils.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		UserName: user.UserName,
		NickName: user.NickName,
	})
	token, err := utils.CreateToken(claims)
	if err != nil {
		global.GGB_LOG.Error("登录获取token失败！", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	// 设置cookie
	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))

	// 记录登录日志
	setLoginLog(c, user.ID, 1)

	response.SuccessWithDetailed(systemResponse.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}

// Logout 登出
func (s *SysUserApi) Logout(c *gin.Context) {
	userId, err := utils.GetUserID(c) // 从token获取用户id
	if err != nil {
		global.GGB_LOG.Error("获取用户id失败！", zap.Error(err))
	} else {
		setLoginLog(c, userId, 0)
	}
	utils.ClearToken(c)
	response.SuccessWithDefaultMessage(c)
}

// GetCaptcha 获取图片验证码
func (s *SysUserApi) GetCaptcha(c *gin.Context) {
	var req request.CaptchaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	var driver base64Captcha.Driver
	var (
		width           int         = 240
		height          int         = 80
		noiseCount      int         = 5
		showLineOptions int         = base64Captcha.OptionShowHollowLine
		bgColor         *color.RGBA = &color.RGBA{R: 242, G: 242, B: 242, A: 255}
		stringSource    string      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"
		chineseSource   string      = "猪猪侠,架构,真,的好,用,频输,高效,可扩展,后端,服务,架构,专为,现代,应用,设计"
	)

	switch req.CaptchaType {
	case "digit":
		driver = base64Captcha.NewDriverDigit(height, width, 5, 0.7, 80)
	case "string":
		driver = base64Captcha.NewDriverString(height, width, noiseCount, showLineOptions, 6, stringSource, bgColor, nil, nil)
	case "math":
		driver = base64Captcha.NewDriverMath(height, width, noiseCount, showLineOptions, bgColor, nil, nil)
	case "chinese":
		driver = base64Captcha.NewDriverChinese(height, width, noiseCount, showLineOptions, 2, chineseSource, bgColor, nil, nil)
	default:
		driver = base64Captcha.NewDriverDigit(height, width, 5, 0.7, 80)
	}

	cp := base64Captcha.NewCaptcha(driver, captchaStore) // 创建验证码对象
	id, b64s, _, err := cp.Generate()                    // 生成验证码图像及其对应的标识（b64s是图片的base64编码）
	if err != nil {
		global.GGB_LOG.Error("生成验证码错误", zap.Error(err))
		response.FailWithMessage("生成验证码错误", c)
		return
	}

	response.SuccessWithData(systemResponse.CaptchaResponse{
		ID:     id,
		Base64: b64s,
	}, c)
}

// VerifyCaptcha 校验图形验证码
func (s *SysUserApi) VerifyCaptcha(c *gin.Context) {
	var req request.Captcha
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	if !captchaStore.Verify(req.CaptchaId, req.Captcha, true) {
		response.FailWithMessage("验证码错误", c)
		return
	}
	response.SuccessWithMessage("验证码校验成功", c)
}
