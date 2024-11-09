package filesystem

import (
	"os"
	"path/filepath"
	"strings"
	"yushu/box/logs"
)

// Exists 判断文件是否存在 存在则为true 不存在则为false
func Exists(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return err == nil
}

// ExecPath 获取当前执行路径
func ExecPath(name ...string) string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 转义
	dir = filepath.ToSlash(dir)
	return JoinPathname(dir, name...)
}

// AppPath 获取当前app目录
func AppPath(name ...string) string {
	// 获取当前执行文件的路径
	exePath, err := os.Executable()
	if err != nil {
		logs.Print("App dir path error:", err)
		return ""
	}

	// 获取绝对路径
	absPath, err := filepath.Abs(exePath)
	if err != nil {
		logs.Print("App dir path error:", err)
		return ""
	}

	// 转义路径
	absPath = filepath.ToSlash(absPath)

	// 拿到最后一次/之前的路径
	absPath = absPath[:strings.LastIndex(absPath, "/")]

	return JoinPathname(absPath, name...)
}

// JoinPathname 拼接路径
func JoinPathname(rootName string, endName ...string) string {
	name := strings.Join(endName, "")
	s2 := ConvertPath(name)
	if len(s2) != 0 {
		if s2[0:1] != "/" {
			s2 = "/" + s2
		}
	}
	return rootName + s2
}

// ConvertPath 转义路径
func ConvertPath(a string) string {
	if len(a) > 0 && len(a) >= 2 {
		if a[0:2] == "./" || a[0:2] == ".\\" {
			a = a[1:]
		}
	}

	// 转义
	a = filepath.ToSlash(a)
	return a
}
