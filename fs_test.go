package wen

import (
	"fmt"
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
