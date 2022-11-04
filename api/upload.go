package api

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"go.uber.org/zap"
	"go.x2ox.com/THz"
	"io"
	"os"
	"sync"
	"video-server/api/result"
	"video-server/model"
	"video-server/util"
)

const tempPath = "D:/video/"

var (
	hashRecord = make(map[string]uint64)
	mediaLock  *sync.Mutex
	chunkLock  *sync.Mutex
)

//todo 文件名和md5的对应保存?

func UploadMediaAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var media model.Media
	if err := c.Bind(&media); err != nil {
		r.BadRequest()
		return
	}

	mediaLock.Lock()
	//todo 1.初始化map值操作 2.hash碰撞问题
	if s, ok := hashRecord[media.MD5]; ok && s == media.Size {
		r.Set(-9, "请勿重复上传")
		mediaLock.Unlock()
		return
	}

	hashRecord[media.MD5] = media.Size
	mediaLock.Unlock()

	filepath := fmt.Sprintf("%s%s", tempPath, media.MD5)
	_ = os.Mkdir(filepath, 0777)

	r.Data = media.MD5
}

func UploadChunkAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var chunk model.Chunk
	if err := c.Bind(&chunk); err != nil {
		r.BadRequest()
		return
	}

	hash := md5.New()
	file, err := chunk.File.Open()
	defer file.Close()
	if err != nil {
		r.Set(-9, err.Error())
		return
	}

	size, _ := io.Copy(hash, file)
	if hex.EncodeToString(hash.Sum(nil)) != chunk.MD5 || chunk.Size != uint64(size) {
		r.Set(-9, fmt.Sprintf("切片%d有错误", chunk.ChunkID))
		return
	}

	filepath := fmt.Sprintf("%s%s/%d", tempPath, chunk.ID, chunk.ChunkID)

	chunkLock.Lock()
	if util.FileExist(filepath) {
		chunkLock.Unlock()
		return // 重复切片 啥也不做
	}

	des, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0766)
	defer des.Close()
	chunkLock.Unlock()
	if err != nil {
		r.Set(-9, err.Error())
		return
	}

	buf := new(bytes.Buffer)
	_, _ = io.Copy(buf, file)
	_, _ = des.Write(buf.Bytes())

	r.Data = chunk.ChunkID
}

type Tmp struct {
	ID    string `json:"id"`
	Count int    `json:"count"`
}

func MergeChunkAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var arg Tmp
	if err := c.Bind(&arg); err != nil {
		r.BadRequest()
		return
	}
	//todo 文件ID不存在 与重复 判断

	des, _ := os.OpenFile(fmt.Sprintf("%s%s_all", tempPath, arg.ID), os.O_RDWR|os.O_CREATE, 0766)
	defer des.Close()

	for i := 0; i < arg.Count; i++ {
		indexFile, err := os.OpenFile(fmt.Sprintf("%s%s/%d", tempPath, arg.ID, i), os.O_RDWR, 0766)
		if err != nil {
			zap.L().Debug(err.Error())
			r.Set(-9, "出现错误，请重新上传")
			return
		}

		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, indexFile)
		_, _ = des.Write(buf.Bytes())

		indexFile.Close()
	}

	_ = os.RemoveAll(fmt.Sprintf("%s%s/", tempPath, arg.ID))
}
