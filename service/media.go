package service

import "video-server/model"

var Media mediaSrv

type mediaSrv struct{}

func (mediaSrv) Save(fileEntity *model.FileMetadata) error {
	return db.Create(fileEntity).Error
}

func (s mediaSrv) Exist(md5 string, size uint64) bool {
	var count int64
	db.Model(&model.FileMetadata{}).Where("md5= ? and size= ?", md5, size).Count(&count)
	return count > 0
}
