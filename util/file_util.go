package util

import (
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
