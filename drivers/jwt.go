package drivers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

// Structs must have exported fields (capitalized) for JSON
type header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type payload struct {
	Sub        int    `json:"sub"`
	Name       string `json:"name"`
	IsVerified bool   `json:"is_verified"`
}

func CreateJWT(sub int, name string, isVerified bool) (string, error) {
	// Step 1: Header
	header := header{
		Alg: "HS256",
		Typ: "JWT",
	}
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}
	headerBase64 := base64.RawURLEncoding.EncodeToString(headerJSON)

	// Step 2: Payload
	payload := payload{
		Sub:        sub,
		Name:       name,
		IsVerified: isVerified,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadBase64 := base64.RawURLEncoding.EncodeToString(payloadJSON)

	// Step 3: Signature
	data := headerBase64 + "." + payloadBase64
	secret := []byte("my-secret-key")
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	signature := h.Sum(nil)
	signatureBase64 := base64.RawURLEncoding.EncodeToString(signature)

	// Step 4: JWT token
	jwt := data + "." + signatureBase64
	return jwt, nil
}
func ValidateJWT(jwt string) (bool, error) {
	parts := strings.Split(jwt, ".")
	if len(parts) != 3 {
		return false, errors.New("invalid token format")
	}

	headerBase64 := parts[0]
	payloadBase64 := parts[1]
	signatureBase64 := parts[2]

	// Recompute signature
	data := headerBase64 + "." + payloadBase64
	secret := []byte("my-secret-key")
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(data))
	expectedSignature := h.Sum(nil)
	expectedSignatureBase64 := base64.RawURLEncoding.EncodeToString(expectedSignature)

	// Compare signatures
	if !hmac.Equal([]byte(signatureBase64), []byte(expectedSignatureBase64)) {
		return false, errors.New("invalid signature")
	}

	return true, nil
}
