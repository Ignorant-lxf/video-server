package util

import (
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
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

// GenerateVideo generate video by file which consists of chunks
func GenerateVideo(path string, arg ...string) {
	// -s 视频缩放比 -c:a 保证音频 -aspect 4:3 (更改长宽比)
	// 增加字幕 ffmpeg -i video.mp4 -i subtitles.srt -c:v copy -c:a copy -preset veryfast -c:s mov_text -map 0 -map 1 output.mp4
	// 添加字幕 ffmpeg -i out.mp4 -vf subtitles=out.srt output.mp4
	command := exec.Command("ffmpeg", "-i", path, "-s", "1024x576", path+arg[0])
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}
}

func GetVideoInfo(path string) {
	command := exec.Command("ffmpeg", "-i", path) // , "-hide_banner"
	if err := command.Run(); err != nil {
		fmt.Println(err)
	}
}

func GetImage(path string) {
	if err := exec.Command("ffmpeg", "-i", path, "-r", "1", "-f", "image2", "image-%2d.png").Run(); err != nil {
		fmt.Println(err)
	}
}
