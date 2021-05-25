package wen

import (
	"log"
	"testing"
)

func TestIsGBK(t *testing.T) {
	s := []byte("附件四福建师范搜福建省")
	if IsUtf8(s) {
		log.Println(Utf8ToGbk([]byte(s)))
	}
}
