package function

import (
	//"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//获取文件名
func Basename(filepath string) string {
	return filepath.Base(string(filepath))
}

// 相当于PHP的 dirname()
func Dirname(dirctory string) string {
	return Substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// 获取当前执行的绝对路径 相当于PHP的__DIR__
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
