package api

import (
	"fmt"
	"go.x2ox.com/THz"
	"os"
	"sync"
	"video-server/api/result"
	"video-server/model"
)

const tempPath = "D:/video/"

var (
	hashRecord = make(map[string]uint64)
	md5Lock    *sync.Mutex
)

func UploadMediaAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var media model.Media
	if err := c.Bind(&media); err != nil {
		r.BadRequest()
		return
	}

	md5Lock.Lock()
	defer md5Lock.Unlock()
	//todo 初始化map值操作
	if s, ok := hashRecord[media.MD5]; ok && s == media.Size {
		r.Set(-9, "请勿重复上传")
		return
	}

	hashRecord[media.MD5] = media.Size

	filepath := fmt.Sprintf("%s%s", tempPath, media.MD5)
	if err := os.Mkdir(filepath, 0777); err != nil {
		r.Set(-9, err.Error())
		return
	}

	r.Data = filepath
}
