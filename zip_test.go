package wen

import "testing"

func TestUnzip(t *testing.T) {
	zipFile := "C:/Users/fd/Downloads/pcp_0.3.3_windows_amd64.zip"
	destDir := "C:/Users/fd/Downloads/pcp"
	if err := Unzip(zipFile, destDir); err != nil {
		t.Error(err)
	}
}
