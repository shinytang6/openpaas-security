package models

import "fmt"

type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
	fmt.Println(fileMetas)
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}