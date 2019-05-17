package util

import (
	"github.com/imroc/req"
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
func DownloadFile(url string, rootDir string) {
	uri := ParsePath(url)

	//创建本地目录
	localFullPath := rootDir + uri
	err := os.MkdirAll(filepath.Dir(localFullPath), 0777)
	if err != nil {
		panic(err)
	}

	resp, err := req.Get(url)
	if err != nil {
		panic(err)
	}

	err = resp.ToFile(localFullPath)
	if err != nil {
		panic(err)
	}
}
