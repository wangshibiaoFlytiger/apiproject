package test

import (
	"apiproject/util"
	"fmt"
	"testing"
)

/**
测试复制文件
*/
func TestCopyFile(t *testing.T) {
	writeByteCount, err := util.CopyFile("/data/workspace/github/go/private/apiproject/schema.sql", "/data/workspace/github/go/private/apiproject/testcopy/schema_dst.sql")
	fmt.Println(writeByteCount, err)
}
