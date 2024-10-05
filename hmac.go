package main

import (
    "crypto/hmac"
    "crypto/sha256"
)

// Pre-shared HMAC key
var preSharedHMACKey = []byte("SuperSecretHMACKey")

func GenerateHMAC(data []byte) []byte {
    mac := hmac.New(sha256.New, preSharedHMACKey)
    mac.Write(data)
    return mac.Sum(nil)
}

func VerifyHMAC(data, receivedMac []byte) bool {
    expectedMac := GenerateHMAC(data)
    return hmac.Equal(expectedMac, receivedMac)
}
