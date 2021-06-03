package wen

import (
	"log"
	"testing"
)

func TestID(t *testing.T) {
	for i := 0; i < 1*100; i++ {
		log.Println(ID())
	}
}
