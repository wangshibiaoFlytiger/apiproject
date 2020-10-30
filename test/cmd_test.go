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
测试执行系统命令
*/
func TestCmd(t *testing.T) {
	stdOut, errOut, err := util.ExecCmd("ffmpeg", "-y", "-i", "/home/wangshibiao/test/test11.mp4", "-c", "copy", "/home/wangshibiao/test/go_ffmpeg.mp4")
	fmt.Println(stdOut)
	fmt.Println(errOut)
	fmt.Println(err)
}
