package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"testing"
)

func TestGetFileType(t *testing.T) {
	file := &multipart.FileHeader{
		Filename: "test.jpg",
	}
	f, err := file.Open()

	f.Seek(0, io.SeekStart)

	if err != nil {
		return
	}

	fileType := GetFileType(file)
	fmt.Println(fileType)
}
