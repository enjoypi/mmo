package ext

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomString(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err == nil {
		return hex.EncodeToString(b)
	}
	return ""
}
