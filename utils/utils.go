package utils

import (
    "log"
    "os"
    "path/filepath"
    "strings"
)
// 获取exe执行的目录
func GetCurrentPath() string{
    dir, err := filepath.Abs(filepath.Dir(os.Args[0]))  //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
    if err != nil {
        log.Fatal(err)
    }
    return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
// 拼接路径
func ConcatPath(path1 string, path2 string) string{
    return path1 + "/" + path2
}