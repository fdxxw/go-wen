package wen

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func Unzip(fileName, dir string) error {
	if !DirExists(dir) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	zr, err := zip.OpenReader(fileName)
	if err != nil {
		return err
	}
	for _, z := range zr.Reader.File {
		if z.FileInfo().IsDir() {
			err := os.MkdirAll(filepath.Join(dir, z.Name), os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		r, err := z.Open()
		if err != nil {
			return err
		}
		defer r.Close()
		newFile, err := os.Create(filepath.Join(dir, z.Name))
		if err != nil {
			return err
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, r)
		if err != nil {
			return err
		}
		newFile.Close()
		r.Close()
	}
	return nil
}
