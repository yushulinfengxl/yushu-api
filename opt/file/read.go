package file

func (files *File) Read(b []byte) (n int, err error) {
	return files.File.Read(b)
}
