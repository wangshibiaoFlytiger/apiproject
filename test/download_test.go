package test

import (
	"apiproject/util"
	"testing"
)

/**
测试文件下载
*/
func TestDownloadFile(t *testing.T) {
	rootDir := "/data/workspace/github/go/private/apiproject/rootDir"
	util.DownloadFileByNetPath("http://gallery.opgirl.cn/2d69ab21707511e98da0759aa36307f9/04c37b81f31683ce109bd4e98beb01bf/88548f2b0f37a45fda202e40c061ebbe/88548f2b0f37a45fda202e40c061ebbe_vertical.jpg", rootDir)
}
