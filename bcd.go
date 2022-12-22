package wen

import (
	"fmt"
	"strings"
)

// 8421 bcd 压缩的
type BCD struct {
	Map map[byte]byte
}

var BCD8421 = &BCD{
	Map: map[byte]byte{
		0: 0x0, 1: 0x1, 2: 0x2, 3: 0x3,
		4: 0x4, 5: 0x5, 6: 0x6, 7: 0x7,
		8: 0x8, 9: 0x9,
	},
}

func (a *BCD) ToNumber(bs []byte, dotIndex int) string {
	mapRev := map[byte]byte{}
	for k, v := range a.Map {
		mapRev[v] = k
	}
	s := ""
	for _, b := range bs {
		right := b & 0x0f
		left := b >> 4
		s += fmt.Sprintf("%d", mapRev[left])
		if len(s) == dotIndex {
			s += "."
		}
		s += fmt.Sprintf("%d", mapRev[right])
		if len(s) == dotIndex {
			s += "."
		}
	}
	return s
}

func (a *BCD) FromNumber(value string, byteLen uint8, dotLen uint8) []byte {
	intLen := int(byteLen*2 - dotLen) // 整数长度
	splits := strings.Split(value, ".")
	intValue := splits[0] // 整数值
	bs := make([]byte, byteLen)
	for i := len(intValue) - 1; i >= 0; i-- {
		v := intValue[i] - 48
		pos := intLen - ((len(intValue) - 1) - i)

		mod := pos % 2
		index := (pos/2 + mod) - 1
		b := bs[index]

		if mod == 0 { // 字节的右半部分 right
			b = (b & 0xf0) | (v & 0x0f)
		} else { // left
			b = (b & 0x0f) | ((v << 4) & 0xf0)
		}
		bs[index] = b
	}
	if len(splits) > 1 {
		dotValue := splits[1]
		for i := 0; i < len(dotValue); i++ {
			v := dotValue[i] - 48
			pos := intLen + 1 + i
			mod := pos % 2
			index := (pos/2 + mod) - 1
			b := bs[index]
			if mod == 0 { // 字节的右半部分 right
				b = (b & 0xf0) | (v & 0x0f)
			} else { // left
				b = (b & 0x0f) | ((v << 4) & 0xf0)
			}
			bs[index] = b
		}
	}
	return bs
}
