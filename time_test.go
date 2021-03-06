package wen

import (
	"log"
	"testing"
	"time"
)

func TestTimestamp(t *testing.T) {
	log.Println(Timestamp(time.Now()))
}
func TestTimeCost(t *testing.T) {
	defer TimeCost()()
	log.Println("test")
}
