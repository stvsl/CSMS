package cos

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
	"stvsljl.com/CSMS/utils"
)

var CosClient *cos.Client
var BucketName = ""

func Init() {
	CosClient = NewCos()
	BucketName = utils.GetCosConfig().Bucket
}

func NewCos() *cos.Client {
	u, _ := url.Parse(utils.GetCosConfig().Domain)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  utils.GetCosConfig().SecretId,
			SecretKey: utils.GetCosConfig().SecretKey,
		},
	})
	return client
}

func GetCosClient() *cos.Client {
	return CosClient
}

// 上传文件
func UploadFile(file *multipart.FileHeader) (string, error) {
	// 打开文件
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	name, _ := utils.GetRandomString(15)
	md5, _ := utils.FileMd5(file)
	// 创建上传路径XXX.文件后缀
	key := "uploads/" + name + md5 + file.Filename[len(file.Filename)-4:]
	// 上传文件
	_, err = GetCosClient().Object.Put(context.Background(), key, src, nil)
	if err != nil {
		return "", err
	}
	// 返回文件在 COS 中的 URL
	url := fmt.Sprintf("http://cdn.stvsljl.com/%s", key)
	fmt.Println(url)
	return url, nil
}

// 删除文件
// string key 文件路径
func DeleteFile(key string) error {
	// 对key进行处理，只保留key后面的部分
	// 例如https://ssimp-1316672330.cos.ap-beijing.myqcloud.com/2023/03/02/AvCbB2APsng%3D.png 转换为 /2023/03/02/AvCbB2APsng%3D.png
	key = key[len(utils.GetCosConfig().Domain):]
	_, err := CosClient.Object.Delete(context.Background(), key, nil)
	return err
}
