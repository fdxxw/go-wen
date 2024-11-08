package wen

import (
	"errors"
	"io"
	"io/fs"
	"net/http"
)

type RootFileSystem interface {
	Open(name string) (http.File, error)
}

type ioFS struct {
	// 根文件夹路径
	rootUrl string
	fsys    fs.FS
}

type ioFile struct {
	file fs.File
}

func (f ioFile) Close() error               { return f.file.Close() }
func (f ioFile) Read(b []byte) (int, error) { return f.file.Read(b) }
func (f ioFile) Stat() (fs.FileInfo, error) { return f.file.Stat() }

var errMissingSeek = errors.New("io.File missing Seek method")
var errMissingReadDir = errors.New("io.File directory missing ReadDir method")

func (f ioFile) Seek(offset int64, whence int) (int64, error) {
	s, ok := f.file.(io.Seeker)
	if !ok {
		return 0, errMissingSeek
	}
	return s.Seek(offset, whence)
}

func (f ioFile) Readdir(count int) ([]fs.FileInfo, error) {
	d, ok := f.file.(fs.ReadDirFile)
	if !ok {
		return nil, errMissingReadDir
	}
	var list []fs.FileInfo
	for {
		dirs, err := d.ReadDir(count - len(list))
		for _, dir := range dirs {
			info, err := dir.Info()
			if err != nil {
				// Pretend it doesn't exist, like (*os.File).Readdir does.
				continue
			}
			list = append(list, info)
		}
		if err != nil {
			return list, err
		}
		if count < 0 || len(list) >= count {
			break
		}
	}
	return list, nil
}

func (f ioFS) Open(name string) (http.File, error) {
	// 拼接根文件夹前缀
	if name == "/" {
		name = f.rootUrl
	} else {
		name = f.rootUrl + name
	}
	file, err := f.fsys.Open(name)
	if err != nil {
		return nil, err
	}
	return ioFile{file}, nil
}

// 拼接指定路径的文件系统，在请求时可省略指定url
func RootFs(fsys fs.FS, url string) RootFileSystem {
	return ioFS{rootUrl: url, fsys: fsys}
}
