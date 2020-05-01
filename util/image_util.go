package util

import (
	"apiproject/log"
	"github.com/disintegration/imaging"
	"go.uber.org/zap"
	"image"
	"os"
)

/**
调整图片分辨率: 同时压缩图片质量
*/
func ResizeImage(srcPath string, dstPath string, width int, height int) error {
	img, err := imaging.Open(srcPath)
	if err != nil {
		log.Logger.Error("调整图片分辨率, 异常", zap.Error(err))
		return err
	}

	img = imaging.Resize(img, width, height, imaging.Linear)

	err = imaging.Save(img, dstPath, imaging.JPEGQuality(30))
	if err != nil {
		log.Logger.Error("调整图片分辨率, 保存图片异常", zap.Error(err))
		return err
	}

	return nil
}

/**
从图片中央按指定分辨率切图: 同时压缩图片质量
*/
func CropImageCenter(srcPath string, dstPath string, width int, height int) error {
	img, err := imaging.Open(srcPath)
	if err != nil {
		log.Logger.Error("从图片中央按指定分辨率切图, 异常", zap.Error(err))
		return err
	}

	//切图
	img = imaging.CropAnchor(img, width, height, imaging.Center)

	err = imaging.Save(img, dstPath, imaging.JPEGQuality(50))
	if err != nil {
		log.Logger.Error("从图片中央按指定分辨率切图, 保存图片异常", zap.Error(err))
		return err
	}

	return nil
}

/**
获取图片信息
*/
func GetImgInfo(imgPath string) (width int, height int, formatName string, err error) {
	file, err := os.Open(imgPath)
	defer file.Close()
	if err != nil {
		log.Logger.Error("获取图片信息, 打开文件, 异常", zap.Error(err))
		return
	}

	config, formatName, err := image.DecodeConfig(file)
	if err != nil {
		log.Logger.Error("获取图片信息, 解码文件, 异常", zap.Error(err))
		return
	}

	return config.Width, config.Height, formatName, nil
}

/**
下载图片到本地
*/
func DownloadImg(imgUrl string, localFullPath string) (width int, height int, formatName string, err error) {
	//下载图片
	if err = DownloadFileByLocalPath(imgUrl, localFullPath); err != nil {
		log.Logger.Error("下载图片到本地, 异常", zap.Error(err))
		return 0, 0, "", err
	}

	//获取图片信息
	return GetImgInfo(localFullPath)
}
