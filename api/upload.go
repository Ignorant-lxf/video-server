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
	"video-server/service"
	"video-server/util"
)

const tempPath = "D:/video/"

var (
	fileRecord = make(map[string]*model.FileMetadata)
	mediaLock  *sync.Mutex
	chunkLock  *sync.Mutex
	mergeLock  *sync.Mutex
)

func UploadMediaAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var media model.Media
	if err := c.Bind(&media); err != nil {
		r.BadRequest()
		return
	}

	mediaLock.Lock()

	if s, ok := fileRecord[media.MD5]; ok && s.Size == media.Size {
		if fileRecord[media.MD5].Status == 1 {
			r.Set(-9, "视频已经上传过了")
		} else {
			r.Set(-9, "进行中，请勿重复上传")
		}

		mediaLock.Unlock()
		return
	}

	entity := &model.FileMetadata{
		Filename: media.Filename,
		MD5:      media.MD5,
		Size:     media.Size,
	}
	fileRecord[media.MD5] = entity

	mediaLock.Unlock()

	if service.Media.Exist(media.MD5, media.Size) {
		entity.Status = 1
		r.Set(-9, "视频已经上传过了")
		return
	}

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

	entity := fileRecord[chunk.ID]
	if entity == nil {
		r.Set(-9, "切片对应的文件不存在")
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

	filepath := fmt.Sprintf("%s%s/%s_%d", tempPath, chunk.ID, entity.Filename, chunk.ChunkID)

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
	ID    string `json:"id"`    // 文件ID
	Count int    `json:"count"` // 切片总数
}

func MergeChunkAction(c *THz.Context) {
	r := result.New[any]()
	defer c.JSON(r)

	var arg Tmp
	if err := c.Bind(&arg); err != nil {
		r.BadRequest()
		return
	}

	entity := fileRecord[arg.ID]
	if entity == nil {
		r.Set(-9, "输入的文件不存在")
		return
	}

	mergeLock.Lock()

	switch entity.Status {
	case 2:
		r.Set(-9, "正在合并中，请勿重复操作")
		return
	case 1:
		r.Set(-9, "文件早已存在")
		return
	}
	entity.Status = 2

	mergeLock.Unlock()

	filepath := fmt.Sprintf("%s%s_all", tempPath, arg.ID)
	file, _ := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0766)
	defer file.Close()

	for i := 0; i < arg.Count; i++ {
		indexFile, err := os.OpenFile(fmt.Sprintf("%s%s/%s_%d", tempPath, arg.ID, entity.Filename, i), os.O_RDWR, 0766)
		if err != nil {
			zap.L().Debug(err.Error())
			r.Set(-9, "出现错误，请重新上传")
			entity.Status = 0
			return
		}

		buf := new(bytes.Buffer)
		_, _ = io.Copy(buf, indexFile)
		_, _ = file.Write(buf.Bytes())

		_ = indexFile.Close()
	}

	entity.Path = filepath
	entity.Status = 1

	if err := service.Media.Save(entity); err != nil {
		r.Set(-9, err.Error())
		entity.Status = 0
		return
	}

	_ = os.RemoveAll(fmt.Sprintf("%s%s/", tempPath, arg.ID))
}
