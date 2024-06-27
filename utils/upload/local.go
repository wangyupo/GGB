package upload

import (
	"errors"
	"github.com/wangyupo/GGB/global"
	"github.com/wangyupo/GGB/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Local struct{}

func (*Local) UploadFile(file *multipart.FileHeader) (filePath string, fileName string, err error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	fileName = name + "_" + time.Now().Format("20060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.GGB_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GGB_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GGB_CONFIG.Local.StorePath + "/" + fileName
	filePath = global.GGB_CONFIG.Local.Path + "/" + fileName

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GGB_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GGB_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))

		return "", "", errors.New("function os.Create() failed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GGB_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filePath, fileName, nil
}

// DeleteFile 删除文件
func (l *Local) DeleteFile(key string) (err error) {
	p := global.GGB_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.GGB_CONFIG.Local.StorePath) {
		err = os.Remove(p)
		if err != nil {
			return errors.New("本地文件删除失败，err:" + err.Error())
		}
	}
	return nil
}
