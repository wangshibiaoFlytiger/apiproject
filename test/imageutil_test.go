package test

import (
	"apiproject/util"
	"testing"
)

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

/**
测试图片处理
*/
func TestImagingImageUtil(t *testing.T) {
	util.ResizeImage("/home/wangshibiao/test/testimg/src.jpg", "/home/wangshibiao/test/testimg/out_resize.jpg", 640, 960)
	util.CropImageCenter("/home/wangshibiao/test/testimg/out_resize.jpg", "/home/wangshibiao/test/testimg/out_crop.jpg", 640, 362)

	width, height, formatName, err := util.GetImgInfo("/home/wangshibiao/test/testimg/src.jpg")
	println(width, height, formatName, err)
}
