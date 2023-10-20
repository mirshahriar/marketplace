package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

var initialVector = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

func Encrypt(key, plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	bytes := []byte(plaintext)
	cfb := cipher.NewCFBEncrypter(block, initialVector)
	cipherText := make([]byte, len(bytes))
	cfb.XORKeyStream(cipherText, bytes)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
