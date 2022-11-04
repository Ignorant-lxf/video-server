package util

import (
	"mime/multipart"
	"os"
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

func FileExist(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !stat.IsDir()
}
