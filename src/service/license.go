package service

import (
	"encoding/hex"
)

const key = "indinsindinsindi"

type Params struct {
	Model  string `json:"model"`
	Number string `json:"factoryNumber"`
	Term   string `json:"licenseTerm"`
	Key    string `json:"licenseKey"`
}

func encrypt(s string) (encryptedString string) {
	return hex.EncodeToString([]byte(s))
}
