package file

import "os"

type File struct {
	File      *os.File
	_stat     os.FileInfo
	statError error
}

// Close 关闭文件
func (files *File) Close() error {
	return files.File.Close()
}

// Stat 获取文件信息
func (files *File) Stat() (os.FileInfo, error) {
	files._stat, files.statError = files.File.Stat()
	return files._stat, files.statError
}

func (files *File) Size() (size int64) {
	if files._stat == nil {
		_, err := files.Stat()
		if err != nil {
			return 0
		}
	}
	size = files._stat.Size()
	return
}

func (files *File) StatName() (name string) {
	if files._stat == nil {
		_, err := files.Stat()
		if err != nil {
			return ""
		}
	}
	name = files._stat.Name()
	return
}

func (files *File) Name() string {
	return files.File.Name()
}

// Seek 移动文件指针
func (files *File) Seek(offset int64, whence int) (int64, error) {
	return files.File.Seek(offset, whence)
}
