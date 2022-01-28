package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
)

const key = "indinsindinsindi"

type License struct {
	Model  string `json:"model"`
	Number string `json:"factoryNumber"`
	Term   string `json:"licenseTerm"`
	Key    string `json:"licenseKey"`
}

func (l *License) Encrypt() (string, error) {
	js, err := json.Marshal(l)
	if err != nil {
		return "", err
	}
	return encrypt(string(js)), nil
}

func Decrypt(cipher string) (License, error) {
	var l License
	js := decrypt(cipher)
	fmt.Println(js)
	err := json.Unmarshal([]byte(js), &l)
	return l, err
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

func decrypt(encryptedString string) (decryptedString string) {
	enc, _ := hex.DecodeString(encryptedString)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
