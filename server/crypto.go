package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

// Encrypt takes plaintext, generates a key,
// encrypts and authenticates the plaintext,
// then returns the encrypted message and key
func Encrypt(plainText []byte) (cipherTextString string, keyString string, err error) {

	key := make([]byte, 64)
	if _, err := rand.Read(key); err != nil {
		return "", "", err
	}

	hashedKey := hashTo32Bytes([]byte(base64.URLEncoding.EncodeToString(key)))

	encrypted, err := encryptAES(hashedKey, plainText)
	if err != nil {
		return "", "", err
	}

	return base64.URLEncoding.EncodeToString(encrypted), base64.URLEncoding.EncodeToString(key), nil
}

func encryptAES(key, data []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// create two 'windows' in to the ciphertext slice.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	encrypted := ciphertext[aes.BlockSize:]

	// populate the IV slice with random data.
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted, data)

	// note that encrypted is still a window in to the ciphertext slice
	return ciphertext, nil
}

// Decrypt takes the ecrypted message and key,
// then it decrypts and authenticates,
// then returns the original plaintext
func Decrypt(cryptoText string, key []byte) (plainText string, err error) {

	encrypted, err := base64.URLEncoding.DecodeString(cryptoText)

	if len(encrypted) < aes.BlockSize {
		return "", fmt.Errorf("cipherText too short. It decodes to %v bytes but the minimum length is 16", len(encrypted))
	}

	hashedKey := hashTo32Bytes(key)

	decrypted, err := decryptAES(hashedKey, encrypted)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

func decryptAES(key, data []byte) ([]byte, error) {

	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(data, data)

	return data, nil
}

func hashTo32Bytes(input []byte) []byte {

	data := sha256.Sum256(input)
	return data[0:]

}
