常用指令：

```bash
# 生成/更新openAPI文档
swag init

# 格式化注释
swag fmt

# 显示帮助信息
swag help
```

本机openAPI地址：

```bash
# 注：需要先启动本地服务后访问
http://localhost:5312/swagger/index.html
```

描述API的注释（用于生成接口的openAPI文档）

```bash
# 通用API
// FuncName   功能说明（如：校验图形验证码）
// @Tags      归属类别（如：Base）
// @Summary   功能说明（如：校验图形验证码）
// @Security  ApiKeyAuth（是否需要权鉴，不需要去掉这行）
// @accept    application/json（接收参数类型）
// @Produce   application/json（返回数据类型）
// @Param     data  body      request.Captcha           true  "Captcha模型"   （入参）
// @Success   200   {object}  response.MsgResponse  	"返回验证码校验成功提示" （出参）
// @Router    /captcha/verify [POST]

# 上传文件API
// UploadFile 上传文件
// @Tags      CommonUploadFile
// @Summary   上传文件
// @Security  ApiKeyAuth
// @accept    multipart/form-data
// @Produce   application/json
// @Param     file  formData  file        true              "上传文件"
// @Success   200   {object}  response.UploadFileResponse   "返回包括文件路径，文件名称"
// @Router    /common/upload [POST]
```