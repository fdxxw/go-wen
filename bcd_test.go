package wen

import (
	"fmt"
	"log"
	"testing"
)

func TestBCD8421(t *testing.T) {
	// fmt.Println(fmt.Sprintf("%d%d", 0x2, 0x3))
	bcd := BCD8421
	fmt.Println(bcd.ToNumber(Hex2Byte("2382"), -1))
	fmt.Println(bcd.ToNumber(Hex2Byte("2382"), 3))
}

func TestFromValue(t *testing.T) {
	bcd := BCD8421
	log.Println(Byte2Hex(bcd.FromNumber("10.2", 4, 2)))
	log.Println(Byte2Hex(bcd.FromNumber("10.2", 4, 1)))
}
