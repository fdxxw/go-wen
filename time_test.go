package wen

import (
	"log"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	log.Println(Timestamp(time.Now()))
}
