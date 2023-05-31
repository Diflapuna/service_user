package service

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func shaHashing(input string) string {

	hash := []byte(input)
	sha256Hash := sha256.Sum256(hash)

	return hex.EncodeToString(sha256Hash[:])
}

func (s *Service) encrypt(value []byte, secret string) []byte {

	hashedSecret := shaHashing(secret)
	aesBlock, err := aes.NewCipher([]byte(hashedSecret))
	if err != nil {
		s.Log.Errorf("Can't create block of cipher %w", err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		s.Log.Errorf("Can't create cipher %w", err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		s.Log.Errorf("Can't count size of ciphered value %w", err)
	}

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)
	return cipheredText
}

func (s *Service) decrypt(ciphered []byte, secret string) []byte {

	hashSecret := shaHashing(secret)
	aesBlock, err := aes.NewCipher([]byte(hashSecret))
	if err != nil {
		s.Log.Errorf("Can't create block of cipher to decrypt %w", err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		s.Log.Errorf("Can't create new cipher to decrypt %w", err)
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, ciphered := ciphered[:nonceSize], ciphered[nonceSize:]

	decryptedValue, err := gcmInstance.Open(nil, nonce, ciphered, nil)
	if err != nil {
		s.Log.Errorf("Can't decrypt value %w", err)
	}

	return decryptedValue
}
