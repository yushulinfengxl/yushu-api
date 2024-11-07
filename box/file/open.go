package file

import (
	"os"
)

func Open(name string, args ...int) (*File, error) {
	osOpenFileFlag := os.O_RDWR | os.O_CREATE
	osOpenFilePerm := os.FileMode(0644)
	if len(args) > 0 {
		osOpenFileFlag = osOpenFileFlag | args[0]
		if len(args) > 1 {
			osOpenFilePerm = os.FileMode(args[1])
		}
	}
	return OsOpen(name, osOpenFileFlag, osOpenFilePerm)
}

func OsOpen(name string, flag int, perm os.FileMode) (*File, error) {
	files := &File{}
	// 获取当前执行文件路径
	filename := ExecPath(name)
	osFile, osErr := os.OpenFile(filename, flag, perm)
	if osErr != nil {
		return nil, osErr
	}
	files.File = osFile
	return files, nil
}
