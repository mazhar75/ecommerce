package drivers

import (
	"crypto/sha256"
	"encoding/hex"
)

// HashSHA256 takes an input string and returns its SHA-256 hash as a hex string
func HashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:]) // convert []byte -> hex string
}
