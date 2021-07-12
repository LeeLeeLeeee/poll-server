package lib

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

func EncryptSha256(value interface{}) (string, error) {
	h := sha256.New()

	switch v := value.(type) {
	case string:
		h.Write([]byte(v))
	case []byte:
		h.Write(v)
	default:
		return "", errors.New("doesn't match the type")

	}

	return hex.EncodeToString(h.Sum(nil)), nil
}
