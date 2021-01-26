package wen

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

type SignType string

const (
	MD5    SignType = "MD5"
	SHA1   SignType = "SHA1"
	SHA256 SignType = "SHA256"
	SHA512 SignType = "SHA512"
)

// MD5加密
func EncryptMD5(message string) string {
	hash := md5.New()
	hash.Write([]byte(message))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// SHA1加密
func EncryptSHA1(message string) string {
	hash := sha1.New()
	hash.Write([]byte(message))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// SHA256加密
func EncryptSHA256(message string) string {
	hash := sha256.New()
	hash.Write([]byte(message))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// SHA512加密
func EncryptSHA512(message string) string {
	hash := sha512.New()
	hash.Write([]byte(message))
	bytes := hash.Sum(nil)
	hashCode := hex.EncodeToString(bytes)
	return hashCode
}

// BASE64编码
func EncryptBASE64(message []byte) string {
	return base64.StdEncoding.EncodeToString(message)
}

// AES 加密
func EncryptAES(data, key []byte) []byte {

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)
	return ciphertext
}

// AES 解密
func DecryptAES(data, key []byte) []byte {

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return plaintext
}

// BASE64解码
func DecryptBASE64(message string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(message)
}

// 生成RSA密钥对
func GenerateRSAKey(bits int, isPKCS8 bool) (string, string, error) {
	if bits < 512 || bits > 2048 {
		return "", "", errors.New("密钥位数需在512-2048之间")
	}
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	var privateDer []byte
	if isPKCS8 {
		privateDer, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return "", "", err
		}
	} else {
		privateDer = x509.MarshalPKCS1PrivateKey(privateKey)
	}
	// 生成公钥
	publicDer, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	return EncryptBASE64(privateDer), EncryptBASE64(publicDer), nil
}

// RSA公钥加密
func EncryptRSA(message, publicKey string) (string, error) {
	key, err := DecryptBASE64(publicKey)
	if err != nil {
		return "", err
	}
	pubKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return "", err
	}
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey.(*rsa.PublicKey), []byte(message))
	if err != nil {
		return "", err
	}
	return EncryptBASE64(encryptedData), nil
}

// RSA私钥解密
func DecryptRSA(message, privateKey string, isPKCS8 bool) (string, error) {
	messageBytes, err := DecryptBASE64(message)
	if err != nil {
		return "", err
	}
	key, err := DecryptBASE64(privateKey)
	if err != nil {
		return "", err
	}
	var priKey interface{}
	if isPKCS8 {
		priKey, err = x509.ParsePKCS8PrivateKey(key)
	} else {
		priKey, err = x509.ParsePKCS1PrivateKey(key)
	}
	if err != nil {
		return "", err
	}
	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), messageBytes)
	if err != nil {
		return "", err
	}
	return string(decryptedData), nil
}

// RSA私钥签名
func SignRSA(message, privateKey string, signType SignType, isPKCS8 bool) (string, error) {
	key, err := DecryptBASE64(privateKey)
	if err != nil {
		return "", err
	}
	var priKey interface{}
	if isPKCS8 {
		priKey, err = x509.ParsePKCS8PrivateKey(key)
	} else {
		priKey, err = x509.ParsePKCS1PrivateKey(key)
	}
	if err != nil {
		return "", err
	}
	var signature []byte
	switch signType {
	case MD5:
		h := md5.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		signature, err = rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.MD5, hash)
	case SHA1:
		h := sha1.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		signature, err = rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA1, hash)
	case SHA256:
		h := sha256.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		signature, err = rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA256, hash)
	case SHA512:
		h := sha512.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		signature, err = rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), crypto.SHA512, hash)
	default:
		return "", errors.New("不支持的签名类型")
	}
	if err != nil {
		return "", err
	}
	return EncryptBASE64(signature), nil
}

// RSA公钥验签
func VerifyRSA(message, publicKey, sign string, signType SignType) error {
	signBytes, err := DecryptBASE64(sign)
	if err != nil {
		return err
	}
	key, err := DecryptBASE64(publicKey)
	if err != nil {
		return err
	}
	pubKey, err := x509.ParsePKIXPublicKey(key)
	if err != nil {
		return err
	}
	switch signType {
	case MD5:
		h := md5.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.MD5, hash, signBytes)
	case SHA1:
		h := sha1.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA1, hash, signBytes)
	case SHA256:
		h := sha256.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hash, signBytes)
	case SHA512:
		h := sha512.New()
		h.Write([]byte(message))
		hash := h.Sum(nil)
		err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA512, hash, signBytes)
	default:
		return errors.New("不支持的签名类型")
	}
	if err != nil {
		return err
	}
	return nil
}
