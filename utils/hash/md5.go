package hash

import (
	"crypto/md5"
	"fmt"
	"io"
	"mime/multipart"
)

func MD5(byteData []byte) string {
	hash := md5.New()
	hash.Write(byteData)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}

func FileMd5(file multipart.File) string {
	hash := md5.New()
	io.Copy(hash, file)
	hashByteData := hash.Sum(nil)
	return fmt.Sprintf("%x", hashByteData)
}
