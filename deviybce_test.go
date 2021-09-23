package wen

import (
	"log"
	"testing"
	"time"
)

func TestDebounce(t *testing.T) {
	debouncer := DebounceNew(10 * time.Millisecond)
	f := func() {
		log.Println("1")
	}
	for i := 0; i < 1000; i++ {
		debouncer(f)
		time.Sleep(time.Millisecond * 5)
	}
	time.Sleep(time.Second)
}
