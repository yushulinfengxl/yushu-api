package file

func (files *File) Write(b []byte) (n int, err error) {
	return files.File.Write(b)
}

func (files *File) WriteString(s string) (n int, err error) {
	return files.File.WriteString(s)
}

func (files *File) WriteAt(b []byte, off int64) (n int, err error) {
	return files.File.WriteAt(b, off)
}
