package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogap/errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	log "github.com/sirupsen/logrus"
	"mime/multipart"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Upload struct {
}

func UploadInstance() *Upload {
	return &Upload{}
}

// UploadFile 文件上传
func (u *Upload) UploadFile(ctx *gin.Context, file *multipart.FileHeader) (string, error) {
	cfg := config.GetConfig().Qiniu
	if cfg.SavePath == "" {
		log.Errorf("UploadFile savepath nil")
		return "", errors.New("未配置文件保存路径")
	}
	if err := utils.CreateDir(cfg.SavePath); err != nil {
		log.Errorf("create dir err=%s", err.Error())
		return "", err
	}
	savePath := cfg.SavePath + "/" + file.Filename
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		log.Errorf("SaveUploadedFile err=%s", err.Error())
		return "", err
	}

	ext := strings.TrimLeft(filepath.Ext(file.Filename), ".")
	qiNiuSaveName := cfg.Path + utils.Md5(file.Filename) + "." + ext

	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)

	putPolicy := storage.PutPolicy{
		Scope: cfg.Bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	storagecfg := new(storage.Config)
	storagecfg.UseHTTPS = false
	storagecfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(storagecfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
		OnProgress: func(fsize, uploaded int64) {
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiNiuSaveName, savePath, &putExtra)
	if err != nil {
		log.Errorf("uploadfile PutFile err=%s", err.Error())
		return "", err
	}
	if err := os.Remove(savePath); err != nil {
		log.Errorf("UploadFile os.Remove err %s", err.Error())
		return "", err
	}

	// 七牛访问地址
	cdnUrl := cfg.Domain + qiNiuSaveName

	return cdnUrl, nil
}

// UploadLocalFile 上传本地文件到七牛云
func (u *Upload) UploadLocalFile(localFilePath string) (string, error) {
	cfg := config.GetConfig().Qiniu
	if cfg.SavePath == "" {
		log.Errorf("UploadLocalFile savepath nil")
		return "", errors.New("未配置文件保存路径")
	}
	if err := utils.CreateDir(cfg.SavePath); err != nil {
		log.Errorf("create dir err=%s", err.Error())
		return "", err
	}

	// 构造在七牛云中的保存文件名
	fileName := filepath.Base(localFilePath)
	ext := strings.TrimLeft(filepath.Ext(fileName), ".")
	currentTime := time.Now().Unix()
	uniqueFileName := fmt.Sprintf("%s_%d", utils.Md5(fileName), currentTime)
	qiNiuSaveName := fmt.Sprintf("%s.%s", uniqueFileName, ext)

	mac := qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", cfg.Bucket, qiNiuSaveName),
	}
	upToken := putPolicy.UploadToken(mac)

	storagecfg := new(storage.Config)
	storagecfg.UseHTTPS = false
	storagecfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(storagecfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{},
		OnProgress: func(fsize, uploaded int64) {
		},
	}

	// 执行上传操作
	err := formUploader.PutFile(context.Background(), &ret, upToken, qiNiuSaveName, localFilePath, &putExtra)
	if err != nil {
		log.Errorf("UploadLocalFile PutFile err=%s", err.Error())
		return "", err
	}

	// 七牛云访问URL
	cdnUrl := cfg.Domain + qiNiuSaveName
	return cdnUrl, nil
}
