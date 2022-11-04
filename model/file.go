package model

import "mime/multipart"

type Media struct {
	MD5      string `json:"md5"`
	Size     uint64 `json:"size"`
	Filename string `json:"filename"`
}

type Chunk struct {
	ID      string                `json:"id"`       // 文件标识
	ChunkID uint64                `json:"chunk_id"` // 切片ID
	Size    uint64                `json:"size"`     // 切片大小
	MD5     string                `json:"md5"`
	File    *multipart.FileHeader `json:"file"`
}

type FileMetadata struct {
	Model

	Filename string `json:"filename"`
	MD5      string `json:"md5"`
	Status   int    `json:"status"` // 0 未合并 1 已存在 2 合并中
	Size     uint64 `json:"size"`
	Path     string `json:"path"` // 文件存储的路径
}
