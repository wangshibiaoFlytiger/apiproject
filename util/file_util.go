package util

import (
	"apiproject/log"
	"github.com/imroc/req"
	"go.uber.org/zap"
	"io"
	"os"
	"path/filepath"
)

/**
获取程序的可执行文件的绝对路径
*/
func GetExePath() (exePath string, err error) {
	exePath, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return exePath, nil
}

/**
下载文件到本地, 按网络路径规则自动创建本地目录
*/
func DownloadFileByNetPath(url string, rootDir string) (localFullPath string, err error) {
	uri := ParsePath(url)

	//创建本地目录
	localFullPath = rootDir + uri
	if err = DownloadFileByLocalPath(url, localFullPath); err != nil {
		log.Logger.Error("下载文件到本地, 按网络路径规则自动创建本地目录, 异常", zap.Any("url", url), zap.Any("rootDir", rootDir), zap.Error(err))
		return "", err
	}

	return localFullPath, nil
}

/**
下载文件到本地
*/
func DownloadFileByLocalPath(url string, localPath string) (err error) {
	//创建本地目录
	if err = CreateFileDir(localPath); err != nil {
		return err
	}

	resp, err := req.Get(url, req.Header{
		"referer": url,
	})
	if err != nil {
		return err
	}

	if err = resp.ToFile(localPath); err != nil {
		return err
	}

	log.Logger.Info("下载文件到本地完成", zap.Any("url", url), zap.Any("path", localPath))
	return nil
}

/**
获取文件所属目录
*/
func GetFileDir(filePath string) string {
	return filepath.Dir(filePath)
}

/**
创建文件所属目录
*/
func CreateFileDir(filePath string) error {
	//创建本地目录
	err := os.MkdirAll(GetFileDir(filePath), 0777)
	if err != nil {
		log.Logger.Error("创建文件所属目录, 异常", zap.Any("filePath", filePath), zap.Error(err))
		return err
	}

	return nil
}

/**
复制文件
*/
func CopyFile(srcPath string, dstPath string) (writeByteCount int64, err error) {
	sourceFileStat, err := os.Stat(srcPath)
	if err != nil {
		log.Logger.Error("复制文件, 检测源文件状态异常", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		log.Logger.Error("复制文件, 源文件不是正常文件", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}

	srcFile, err := os.Open(srcPath)
	if err != nil {
		log.Logger.Error("复制文件, 打开源文件异常", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}
	defer srcFile.Close()

	if err = CreateFileDir(dstPath); err != nil {
		log.Logger.Error("复制文件, 创建目标文件所属目录异常", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}

	dstFile, err := os.Create(dstPath)
	if err != nil {
		log.Logger.Error("复制文件, 创建目标文件异常", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}

	defer dstFile.Close()
	nBytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Logger.Error("复制文件, 复制文件异常", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath), zap.Error(err))
		return 0, err
	}

	log.Logger.Info("复制文件, 完成", zap.Any("srcPath", srcPath), zap.Any("dstPath", dstPath))
	return nBytes, err
}
