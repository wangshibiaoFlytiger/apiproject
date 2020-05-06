package c_upload

import "mime/multipart"

type ReqUploadFile struct {
	File *multipart.FileHeader `binding:"required"`
}
