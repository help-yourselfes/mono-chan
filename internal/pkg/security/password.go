package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	pattern = "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s"
	time    = 2
	memory  = 19 * 1024
	threads = 1
	keyLen  = 32
	saltLen = 16
)

func Hash(password string) (string, error) {
	salt := make([]byte, saltLen)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf(pattern, memory, time, threads, encodedSalt, encodedHash), nil
}

func Verify(password, storedHash string) (bool, error) {
	var memory, time uint32
	var parallel uint8
	var salt, hash []byte

	_, err := fmt.Scanf(storedHash, pattern, &memory, &time, &parallel, &salt, &hash)
	if err != nil {
		return false, err
	}

	decodedSalt, err := base64.RawStdEncoding.DecodeString(string(salt))
	if err != nil {
		return false, err
	}
	decodedHash, err := base64.RawStdEncoding.DecodeString(string(hash))
	if err != nil {
		return false, err
	}

	comparisonHash := argon2.IDKey([]byte(password), decodedSalt, time, memory, threads, uint32(len(decodedHash)))

	return string(comparisonHash) == string(decodedHash), nil
}
