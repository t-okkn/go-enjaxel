package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)


func EncryptByCFB(key []byte, plainText string) (string, error) {
	if !(len(key) == 16 || len(key) == 24 || len(key) == 32) {
		e := errors.New("鍵長は128bit, 192bit, 256bitのいずれかでなくてはなりません")
		return "", e
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	btxt := []byte(plainText)
	cipherdata := make([]byte, aes.BlockSize + len(btxt))

	iv := cipherdata[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherdata[aes.BlockSize:], btxt)

	return base64.StdEncoding.EncodeToString(cipherdata), nil
}

func EasyEncryptByCFB(plainText string) (string, []byte, error) {
	key, err := GenerateSafeKey(32)
	if err != nil {
		return "", nil, err
	}

	base64Text, err := EncryptByCFB(key, plainText)
	if err != nil {
		return "", nil, err
	}

	return base64Text, key, nil
}

func DecryptByCFB(key []byte, base64Text string) (string, error) {
	if !(len(key) == 16 || len(key) == 24 || len(key) == 32) {
		e := errors.New("鍵長は128bit, 192bit, 256bitのいずれかでなくてはなりません")
		return "", e
	}

	data, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(data) < aes.BlockSize {
		return "", errors.New("暗号テキストが不正です")
	}

	iv := data[:aes.BlockSize]
	plaindata := data[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(plaindata, plaindata)

	return string(plaindata), nil
}

func EncryptByGCM(key []byte, plainText string) (string, error) {
	if !(len(key) == 16 || len(key) == 24 || len(key) == 32) {
		e := errors.New("鍵長は128bit, 192bit, 256bitのいずれかでなくてはなりません")
		return "", e
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Unique nonce is required(NonceSize 12byte)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	btxt := []byte(plainText)
	cipherdata := gcm.Seal(nil, nonce, btxt, nil)
	cipherdata = append(nonce, cipherdata...)

	return base64.StdEncoding.EncodeToString(cipherdata), nil
}

func EasyEncryptByGCM(plainText string) (string, []byte, error) {
	key, err := GenerateSafeKey(32)
	if err != nil {
		return "", nil, err
	}

	base64Text, err := EncryptByGCM(key, plainText)
	if err != nil {
		return "", nil, err
	}

	return base64Text, key, nil
}

func DecryptByGCM(key []byte, base64Text string) (string, error) {
	if !(len(key) == 16 || len(key) == 24 || len(key) == 32) {
		e := errors.New("鍵長は128bit, 192bit, 256bitのいずれかでなくてはなりません")
		return "", e
	}

	data, err := base64.StdEncoding.DecodeString(base64Text)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := data[:gcm.NonceSize()]
	plaindata, err := gcm.Open(nil, nonce, data[gcm.NonceSize():], nil)
	if err != nil {
		return "", err
	}

	return string(plaindata), nil
}

func GenerateSafeKey(length int) ([]byte, error) {
	key := make([]byte, length)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	return key, nil
}

