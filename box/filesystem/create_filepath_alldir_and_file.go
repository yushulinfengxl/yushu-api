package filesystem

import (
	"os"
	"path/filepath"
)

// Create CreateFilepathAllDirAndFile 创建文件路径
func Create(pathname string) (file *os.File, err error) {
	// 使用 ConvertPath 将路径转换成标准形式
	pathname = ConvertPath(pathname)

	// 获取文件的父级目录路径
	dirPath := filepath.Dir(pathname)

	// 创建父级目录，0777 表示目录权限
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return nil, err
	}

	// 创建文件，如果文件已经存在，则会打开它，os.O_CREATE 表示文件不存在时会创建
	file, err = os.OpenFile(pathname, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	// 返回文件对象和可能的错误
	return
}
