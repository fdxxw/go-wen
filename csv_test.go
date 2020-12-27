package wen

import (
	"os"
	"testing"
)

func TestCsvAppend(t *testing.T) {
	e := Example{
		Id:   "1",
		Name: "admin",
	}
	CsvWrite("test.csv", []Example{e})
	CsvAppend("test.csv", []Example{e})
	os.Remove("test.csv")
}
