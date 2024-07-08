package system

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyupo/GGB/enums"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/model/common/response"
	"github.com/wangyupo/GGB/model/log"
	"github.com/wangyupo/GGB/model/system/request"
	systemResponse "github.com/wangyupo/GGB/model/system/response"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"time"
)

type SysLoginApi struct{}

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
// @Tags      Base
// @Summary   登录
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Login           true  "SysDictionary模型"
// @Success   200   {object}  response.Response{data=systemResponse.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router    /login [POST]
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
// @Tags      Base
// @Summary   登出
// @Security  ApiKeyAuth
// @Produce   application/json
// @Success   200   {object}  response.MsgResponse  "返回操作成功提示"
// @Router    /logout [POST]
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
// @Tags      Base
// @Summary   获取图片验证码
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.CaptchaRequest           true  "CaptchaRequest模型"
// @Success   200   {object}  systemResponse.CaptchaResponse  "返回base64和图形ID"
// @Router    /captcha [POST]
func (s *SysUserApi) GetCaptcha(c *gin.Context) {
	var req request.CaptchaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	id, b64s, err := utils.CreateCaptcha(req.CaptchaType)
	if err != nil {
		global.GGB_LOG.Error("生成验证码失败！", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.SuccessWithData(systemResponse.CaptchaResponse{
		ID:     id,
		Base64: b64s,
	}, c)
}

// VerifyCaptcha 校验图形验证码
// @Tags      Base
// @Summary   校验图形验证码
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.Captcha           true  "Captcha模型"
// @Success   200   {object}  response.MsgResponse  	"返回验证码校验成功提示"
// @Router    /captcha/verify [POST]
func (s *SysUserApi) VerifyCaptcha(c *gin.Context) {
	var req request.Captcha
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.HandleValidatorError(err, c)
		return
	}

	if ok := utils.VerifyCaptcha(req.CaptchaId, req.Captcha); !ok {
		response.FailWithMessage("验证码错误", c)
		return
	}

	response.SuccessWithMessage("验证码校验成功", c)
}
