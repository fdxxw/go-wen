package wen

import (
	"crypto/rand"
	"encoding/hex"
	"testing"
)

func TestEncryptAES(t *testing.T) {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}
	a := hex.EncodeToString(bytes)
	b, _ := hex.DecodeString(a)
	println(a)
	println(hex.EncodeToString(b))
	password := []byte("123456")
	d := EncryptAES(password, bytes)
	println(string(DecryptAES(d, bytes)))
}

func TestBcryptCompare(t *testing.T) {
	Bcrypt("1qaz@WSX")
	BcryptCompare("$2a$10$mP01Fk0gbMqCPipoFF7WX.IWcbwljKrs11KA2Von008JDjdafO4ea", "1qaz@WSX")
}
