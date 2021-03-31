package wen

import (
	"fmt"
	"log"
	"testing"
)

func TestGet(t *testing.T) {
	m := Map{}
	m["1"] = "2"
	log.Println(m.Get("1"))
}
func TestSet(t *testing.T) {
	m := Map{}
	m.Set("a.v.b", "1")
	fmt.Println(m)
}
