package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
)

const key = "indinsindinsindi"

type LicenseService struct {
	Model  string `json:"model"`
	Number string `json:"factoryNumber"`
	Term   string `json:"licenseTerm"`
	Key    string `json:"licenseKey"`
}

func (l *LicenseService) Encrypt() (string, error) {
	js, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	return encrypt(string(js)), nil
}

func encrypt(stringToEncrypt string) (encryptedString string) {
	plaintext := []byte(stringToEncrypt)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)
}
