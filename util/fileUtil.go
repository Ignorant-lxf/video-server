package util

import (
	"mime/multipart"
	"strings"
)

func GetFileType(file *multipart.FileHeader) string {
	if file == nil {
		return ""
	}

	filename := file.Filename
	index := strings.LastIndex(filename, ".")

	if index == -1 {
		return ""
	}

	return filename[index+1:]
}
