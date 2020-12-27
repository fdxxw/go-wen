package wen

import (
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	m := Map{}
	m["1"] = "2"
	log.Println(m.Get("1"))
}
