package service

import "video-server/model"

var Media mediaSrv

type mediaSrv struct{}

func (mediaSrv) Save(fileEntity *model.FileEntity) error {
	return db.Create(fileEntity).Error
}

func (s mediaSrv) Exist(md5 string) bool {
	var count int64
	db.Model(&model.FileEntity{}).Where("md5=?", md5).Count(&count)
	return count > 0
}
