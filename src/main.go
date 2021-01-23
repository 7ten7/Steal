package main

import (
	"flag"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strings"
)

func main() {
	// 创建OSSClient实例。
	client, err := oss.New("oss-cn-hangzhou.aliyuncs.com", "LTAI4G4EWr5HPGj754qTuNWx", "Jas8Fgg0nVz5RMASJBD6P9IVwjz40N")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("srccodebak")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 指定存储类型为较少访问
	storageType := oss.ObjectStorageClass(oss.StorageIA)

	// 指定访问权限为公共读，缺省为继承bucket的权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 上传字符串。
	err = bucket.PutObject("test.txt", strings.NewReader("hello world"), storageType, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}

func init() {
	var (
		AccessKeyId     = flag.String("k", "", "AccessKeyId")
		AccessKeySecret = flag.String("s", "", "AccessKeySecret")
		file            = flag.String("f", "", "local file path")
		h               = flag.Bool("h", true, "help")
	)
	Banner := " ____                              __    ___               ______              ___      \n/\\  _`\\                           /\\ \\__/\\_ \\             /\\__  _\\            /\\_ \\     \n\\ \\,\\L\\_\\     __    ___  _ __   __\\ \\ ,_\\//\\ \\   __  __   \\/_/\\ \\/   ___    __\\//\\ \\    \n \\/_\\__ \\   /'__`\\ /'___/\\`'__/'__`\\ \\ \\/ \\ \\ \\ /\\ \\/\\ \\     \\ \\ \\  / __`\\ / __`\\ \\ \\   \n   /\\ \\L\\ \\/\\  __//\\ \\__\\ \\ \\/\\  __/\\ \\ \\_ \\_\\ \\\\ \\ \\_\\ \\     \\ \\ \\/\\ \\L\\ /\\ \\L\\ \\_\\ \\_ \n   \\ `\\____\\ \\____\\ \\____\\ \\_\\ \\____\\\\ \\__\\/\\____\\/`____ \\     \\ \\_\\ \\____\\ \\____/\\____\\\n    \\/_____/\\/____/\\/____/\\/_/\\/____/ \\/__/\\/____/`/___/> \\     \\/_/\\/___/ \\/___/\\/____/\n                                                     /\\___/                             \n                                                     \\/__/      By @ 7TEN7               "
	flag.Parse()
	fmt.Println(Banner)
	if *h == true {
		fmt.Println("SecretlyTool version: SecretlyTool/0.0.1\nUsage: SecretlyTool [-k AccessKeyId] [-s AccessKeySecret] [-f file]\n\nOptions:\n")
		flag.PrintDefaults()
		return
	}
	fmt.Println(*AccessKeyId)
	fmt.Println(*AccessKeySecret)
	fmt.Println(*file)

}
