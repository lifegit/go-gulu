package crypto

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// EncodeSha256 sha256 encryption
func EncodeSha256(value string) string {
	m := sha256.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeSha384 sha384 encryption
func EncodeSha384(value string) string {
	m := sha512.New384()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

// EncodeSha256 sha512 encryption
func EncodeSha512(value string) string {
	m := sha512.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
