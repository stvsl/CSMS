package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

// 异或运算
func Xor(a, b string) string {
	// 保证a,b长度相同
	if len(a) > len(b) {
		b = b + a[len(b):]
	}
	if len(a) < len(b) {
		a = a + b[len(a):]
	}
	// 异或运算
	var result string
	for i := 0; i < len(a); i++ {
		result = result + string(a[i]^b[i])
	}
	return result
}

// MD5
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 计算文件的MD5
func FileMd5(file *multipart.FileHeader) (string, error) {
	h := md5.New()
	f, err := file.Open()
	defer f.Close()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(h, f)
	return hex.EncodeToString(h.Sum(nil)), err
}
