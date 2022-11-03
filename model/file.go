package model

type Media struct {
	MD5      string `json:"md5"`
	Size     uint64 `json:"size"`
	Filename string `json:"filename"`
}
