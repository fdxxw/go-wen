package wen

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	a := hex.EncodeToString(bytes)
	b, _ := hex.DecodeString("ddea76f09505ce6849e73c61d7399dd8")
	println(a)
	println(hex.EncodeToString(b))
	password := []byte("wxfa922f34e3eea513")
	d := EncryptAES(password, b)
	log.Println(hex.EncodeToString(d))
	println(string(DecryptAES(d, b)))
}

func TestBcryptCompare(t *testing.T) {
	Bcrypt("1qaz@WSX")
	BcryptCompare("$2a$10$mP01Fk0gbMqCPipoFF7WX.IWcbwljKrs11KA2Von008JDjdafO4ea", "1qaz@WSX")
}

func TestHexStringToBytes(t *testing.T) {
	b, err := hex.DecodeString("8AA8573565BFF96BE67C32142B41D0872CC674B4EC1DB341A3260F10152FEBA8")
	log.Println(b)
	// key := []byte("46EBA22EF5204DD5B110A1F730513RQW")
	key, err := hex.DecodeString("ddea76f09505ce6849e73c61d7399dd8")
	if err != nil {
		panic(err)
	}

	log.Println((DecryptAESECBPKCS5(b, key)))

	// 加密
	// ecrypt := NewECBEncrypter(block)

	// src := PKCS5Padding([]byte("wxfa922f34e3eea513"), ecrypt.BlockSize())
	// out = make([]byte, len(src))
	// ecrypt.CryptBlocks(out, src)
	// log.Println(strings.ToUpper(hex.EncodeToString(out)))
	//log.Println(string(DecryptAES(b, []byte("46EBA22EF5204DD5B110A1F730513RQW"))))
}
