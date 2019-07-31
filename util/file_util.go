package util

import (
	"apiproject/log"
	"github.com/imroc/req"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

/**
获取程序的可执行文件的绝对路径
*/
func GetExePath() string {
	exePath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	return exePath
}

/**
下载文件到本地, 按网络路径规则自动创建本地目录
*/
func DownloadFileByNetPath(url string, rootDir string) {
	uri := ParsePath(url)

	//创建本地目录
	localFullPath := rootDir + uri
	DownloadFileByLocalPath(url, localFullPath)
}

/**
下载文件到本地
*/
func DownloadFileByLocalPath(url string, localPath string) {
	//创建本地目录
	err := os.MkdirAll(filepath.Dir(localPath), 0777)
	if err != nil {
		panic(err)
	}

	resp, err := req.Get(url)
	if err != nil {
		panic(err)
	}

	err = resp.ToFile(localPath)
	if err != nil {
		panic(err)
	}

	log.Logger.Info("下载文件到本地完成", zap.Any("url", url), zap.Any("path", localPath))
}
