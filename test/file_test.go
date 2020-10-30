package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试复制文件
*/
func TestCopyFile(t *testing.T) {
	writeByteCount, err := util.CopyFile("/data/workspace/github/go/private/apiproject/schema.sql", "/data/workspace/github/go/private/apiproject/testcopy/schema_dst.sql")
	fmt.Println(writeByteCount, err)
}
