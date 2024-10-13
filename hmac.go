package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

// Generate HMAC (without needing GroupKey during distribution phase)
func GenerateHMAC(data []byte) []byte {
	h := hmac.New(sha256.New, []byte("static-secret-hmac-key")) // Use a static key for integrity check
	h.Write(data)
	return h.Sum(nil)
}

// Verify HMAC (without needing GroupKey during distribution phase)
func VerifyHMAC(data, receivedMac []byte) bool {
	expectedMac := GenerateHMAC(data)
	return hmac.Equal(expectedMac, receivedMac)
}
