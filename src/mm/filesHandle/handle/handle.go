package handle

import "os"

//FileHandle
type FileHandle struct {
}

func (FileHandle) Close() error {
	panic("implement me")
}

func (FileHandle) Read(p []byte) (n int, err error) {
	panic("implement me")
}

func (FileHandle) Seek(offset int64, whence int) (int64, error) {
	panic("implement me")
}

func (FileHandle) Readdir(count int) ([]os.FileInfo, error) {
	panic("implement me")
}

func (FileHandle) Stat() (os.FileInfo, error) {
	panic("implement me")
}
