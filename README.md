# Steal
渗透时遇到需要下载一整个文件夹的内容时比较麻烦，故写了个小工具实现文件打包，并上传至阿里云OSS，方便信息的获取。

```
   _____  __                __
  / ___/ / /_ ___   ____ _ / /
  \__ \ / __// _ \ / __ `// /
 ___/ // /_ /  __// /_/ // /
/____/ \__/ \___/ \__,_//_/
                                    By @ 7TEN7
Steal version: Steal/0.0.2
Usage: Steal [-k AccessKeyId] [-s AccessKeySecret] [-e Endpoint] [-b Bucket] [-f file]

Options:

  -b string
        Bucket
  -e string
        Endpoint
  -f string
        local file path
  -h    help
  -k string
        AccessKeyId
  -s string
        AccessKeySecret
```
