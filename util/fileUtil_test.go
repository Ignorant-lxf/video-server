package util

import (
	"fmt"
	"io"
	"mime/multipart"
	"sort"
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

func TestSortStrings(t *testing.T) {
	str := []string{"d:/testlxdaaa/teaafdsadad_5",
		"d:/testlxdaaa/teaafdsadad_2",
		"d:/testlxdaaa/teaafdsadad_4",
		"d:/testlxdaaa/teaafdsadad_1",
		"d:/testlxdaaa/teaafdsadad_3",
		"d:/testlxdaaa/teaafdsadad_9",
		"d:/testlxdaaa/teaafdsadad_8",
		"d:/testlxdaaa/teaafdsadad_6",
		"d:/testlxdaaa/teaafdsadad_7",
	}
	sort.Strings(str)

	t.Log(str)
}
