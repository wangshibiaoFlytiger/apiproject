package c_upload

import "mime/multipart"

/**
@author 王世彪
	个人博客: https://sofineday.com?from=apiproject
	微信: 645102170
	QQ: 645102170
*/

type ReqUploadFile struct {
	File *multipart.FileHeader `binding:"required" json:"file"`
}
