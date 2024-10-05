package main

import (
    "crypto/elliptic"
    "crypto/rand"
    "math/big"
)

func GenerateECCKeys(device *Device) error {
    curve := device.Curve
    privateKey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
    if err != nil {
        return err
    }
    device.PrivateKey = new(big.Int).SetBytes(privateKey)
    device.PublicKeyX = x
    device.PublicKeyY = y
    return nil
}

func ComputeSharedSecret(ownPrivKey *big.Int, peerPubX, peerPubY *big.Int, curve elliptic.Curve) []byte {
    sharedX, _ := curve.ScalarMult(peerPubX, peerPubY, ownPrivKey.Bytes())
    return sharedX.Bytes()
}
