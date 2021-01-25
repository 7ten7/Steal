package upload

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

func Upload(file string, AccessKeyId string, AccessKeySecret string, Endpoint string, Bucket string) error {
	// 创建OSSClient实例。
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(Bucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	var ObjectName []string
	if strings.Contains(file, "/") {
		ObjectName = strings.Split(file, "/")
		ObjectName = ObjectName[len(ObjectName)-1:]
	} else if strings.Contains(file, "\\") {
		ObjectName = strings.Split(file, "\\")
		ObjectName = ObjectName[len(ObjectName)-1:]
	} else {
		ObjectName = append(ObjectName, file)
	}

	// 上传本地文件。
	err = bucket.PutObjectFromFile(ObjectName[0], file)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}
