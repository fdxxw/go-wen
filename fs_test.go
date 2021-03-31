package wen

import (
	"fmt"
	"log"
	"testing"
)

func TestFileExists(t *testing.T) {
	if FileExists("../go-wen") {
		fmt.Println("Example file exists")
	} else {
		fmt.Println("Example file does not exist (or is a directory)")
	}
}
func TestDirExists(t *testing.T) {
	if DirExists("fs.go") {
		fmt.Println("exists")
	} else {
		fmt.Println("does not exist")
	}
}

func TestCopyFile(t *testing.T) {
	from := "fs.go"
	to := "C:/Users/fd/Downloads/fs.go"
	err := CopyFile(from, to)
	if err != nil {
		log.Fatal(err)
	}
}

func TestFileTemp(t *testing.T) {
	f, err := FileTemp()
	if err != nil {
		t.Error(err)
	}
	f.Close()
}
