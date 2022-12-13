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

	_, _ = f.Seek(0, io.SeekStart)

	if err != nil {
		return
	}

	fileType := GetFileType(file)
	fmt.Println(fileType)
}

func TestGenerateVideo(t *testing.T) {
	path := "D:/video/in1_all"
	GenerateVideo(path)
}

func TestGetVideoInfo(t *testing.T) {
	path := "D:/video/in1_all.mp4"
	GetVideoInfo(path)
}

func TestGetImage(t *testing.T) {
	path := "D:/video/in1_all.mp4"
	GetImage(path)
}

func TestDefer(t *testing.T) {
	i := func() (result int) {
		defer func() {
			result++
		}()
		return 2
	}()
	fmt.Println(i)

	r := func() int {
		r := 5
		defer func() {
			r++
		}()
		return r
	}()
	fmt.Println(r)
}
