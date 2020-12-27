package wen

import (
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
)

type Example struct {
	Id   string `csv:"id"`
	Name string `csv:"name"`
}

func CsvWrite(path string, data interface{}) error {
	dir := filepath.Dir(path)
	if dir != "" {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return gocsv.MarshalFile(data, f)
}

func CsvRead(path string, i interface{}) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return gocsv.UnmarshalFile(f, i)
}

func CsvAppend(path string, data interface{}) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	return gocsv.MarshalWithoutHeaders(data, f)
}

// func CsvFind(path string, args map[string]string, r interface{}) error {
// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	gocsv.UnmarshalToCallback(f, func(t interface{}) {

// 	})
// }
