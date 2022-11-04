package service

import "video-server/model"

func AutoMigrate() error { return db.AutoMigrate(model.Models...) }
