package util

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"
)



func GetFileSize(filename string) int64 {
	var result int64
	filepath.Walk(filename, func(path string, info os.FileInfo, err error) error {
		result = info.Size()
		return nil
	})
	return result
}

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}