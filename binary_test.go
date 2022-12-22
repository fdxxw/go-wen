package wen

import (
	"fmt"
	"log"
	"testing"
)

func TestHex2Byte(t *testing.T) {
	b := Hex2Byte("00010000")
	fmt.Printf("bcd: %x\n", b)
}

func TestByteRev(t *testing.T) {
	bs := Hex2Byte("2382")
	fmt.Println(fmt.Sprintf("%x", ByteRev(bs)))
}

func TestByte2Hex(t *testing.T) {
	fmt.Println(Byte2Hex([]byte{0x20, 0x40, 0x0f, 0xf3}))
}

func TestByteSplit(t *testing.T) {
	bs := []byte{1, 2, 3, 4, 5, 6, 7}
	log.Println(ByteSplit(bs, 7))

}
