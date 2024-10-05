package main

import (
    "crypto/elliptic"
    "crypto/sha256"
    "fmt"
    "math/big"
)

func main() {
    // Parameters
    numDevices := 5
    threshold := 3

    // Initialize devices
    devices := make([]*Device, numDevices)
    for i := 0; i < numDevices; i++ {
        devices[i] = &Device{
            ID:        fmt.Sprintf("Device%d", i+1),
            Peers:     make(map[string]*Device),
            Threshold: threshold,
            Curve:     elliptic.P256(),
        }
        GenerateECCKeys(devices[i])
    }

    // Group Key Generation by Initiator (Device1)
    groupKey := GenerateGroupKey(devices[0], devices[1:])

    // Distribute Shares
    DistributeShares(devices, groupKey)

    // Reconstruct Group Key
    ReconstructGroupKey(devices)

    // Simulate a device leaving
    fmt.Println("\nSimulating Device3 leaving...")
    devices = append(devices[:2], devices[3:]...) // Remove Device3

    // Reconstruct Group Key after Device3 leaves
    ReconstructGroupKey(devices)
}

func GenerateGroupKey(initiator *Device, peers []*Device) []byte {
    sharedSecrets := make([][]byte, len(peers))
    for i, peer := range peers {
        secret := ComputeSharedSecret(initiator.PrivateKey, peer.PublicKeyX, peer.PublicKeyY, initiator.Curve)
        sharedSecrets[i] = secret
    }

    // Combine shared secrets to form group key
    groupKey := sharedSecrets[0]
    for _, secret := range sharedSecrets[1:] {
        groupKey = xorBytes(groupKey, secret)
    }

    // Set group key for initiator
    initiator.GroupKey = groupKey
    return groupKey
}

func DistributeShares(devices []*Device, groupKey []byte) {
    // Convert groupKey to big.Int for SSS
    secret := new(big.Int).SetBytes(groupKey)
    coefficients, _ := GeneratePolynomial(secret, devices[0].Threshold-1)
    shares, _ := GenerateShares(coefficients, len(devices))

    for i, device := range devices {
        device.Share = shares[i]
        // Generate HMAC for the share using pre-shared key
        shareBytes := append(device.Share.X.Bytes(), device.Share.Y.Bytes()...)
        device.HMAC = GenerateHMAC(shareBytes)
    }

    // Simulate P2P Share Distribution
    for _, sender := range devices {
        for _, receiver := range devices {
            if sender.ID != receiver.ID {
                SendShare(sender, receiver)
            }
        }
    }
}

func ReconstructGroupKey(devices []*Device) {
    fmt.Println("\nReconstructing Group Key with available devices...")
    // Collect shares from at least threshold devices
    var collectedShares []*Share
    for _, device := range devices {
        collectedShares = append(collectedShares, device.Share)
        if len(collectedShares) >= device.Threshold {
            break
        }
    }
    reconstructedSecret := ReconstructSecret(collectedShares)
    reconstructedKey := reconstructedSecret.Bytes()
    // Verify the reconstructed key
    originalKey := devices[0].GroupKey
    if sha256.Sum256(reconstructedKey) == sha256.Sum256(originalKey) {
        fmt.Println("Group Key successfully reconstructed!")
    } else {
        fmt.Println("Failed to reconstruct Group Key.")
    }
}

func xorBytes(a, b []byte) []byte {
    n := len(a)
    if len(b) < n {
        n = len(b)
    }
    result := make([]byte, n)
    for i := 0; i < n; i++ {
        result[i] = a[i] ^ b[i]
    }
    return result
}
