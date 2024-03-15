package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
)

func Encrypt(keyString string, stringToEncrypt string) (encryptedString string, err error) {
	// convert key to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}
	plaintext := []byte(stringToEncrypt)

	// pad plaintext to a multiple of the block size
	blockSize := aes.BlockSize
	padding := blockSize - len(plaintext)%blockSize
	paddedPlaintext := append(plaintext, bytes.Repeat([]byte{byte(padding)}, padding)...)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the cipherText.
	cipherText := make([]byte, aes.BlockSize+len(paddedPlaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], paddedPlaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(cipherText), nil
}

// decrypt from base64 to decrypted string
func Decrypt(keyString string, stringToDecrypt string) (decryptedString string, err error) {
	// convert key to bytes
	key, err := hex.DecodeString(keyString)
	if err != nil {
		return "", err
	}
	cipherText, err := base64.URLEncoding.DecodeString(stringToDecrypt)
	if err != nil {
		return "", err
	}

	blockSize := aes.BlockSize
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(cipherText) < blockSize {
		return "", errors.New("cipherText too short")
	}

	// The IV is included at the beginning of the cipherText.
	iv := cipherText[:blockSize]
	cipherText = cipherText[blockSize:]

	// create a new stream for decrypting
	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(cipherText, cipherText)

	// remove PKCS#7 padding from the decrypted plaintext
	padding := int(cipherText[len(cipherText)-1])
	if padding < 1 || padding > blockSize {
		return "", errors.New("invalid padding")
	}
	decryptedPlaintext := cipherText[:len(cipherText)-padding]

	return string(decryptedPlaintext), nil
}
