package main

import (
    "crypto/elliptic"
    "math/big"
)

// Share represents a share in Shamir's Secret Sharing
type Share struct {
    X *big.Int
    Y *big.Int
}

// Device represents an IoT device in the network
type Device struct {
    ID         string
    PrivateKey *big.Int
    PublicKeyX *big.Int
    PublicKeyY *big.Int
    Share      *Share
    HMAC       []byte
    Peers      map[string]*Device
    GroupKey   []byte
    Threshold  int
    Curve      elliptic.Curve
}
