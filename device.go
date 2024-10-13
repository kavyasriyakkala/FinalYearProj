package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

// Device struct representing each IoT device
type Device struct {
	ID         string
	PrivateKey *big.Int
	PublicKeyX *big.Int
	PublicKeyY *big.Int
	Share      *Share
	Peers      map[string]*Device
	Curve      elliptic.Curve
	Threshold  int
	GroupKey   *big.Int
}

// InitializeDevice creates a device with ECC keys and threshold parameters
func InitializeDevice(id string, threshold int) *Device {
	device := &Device{
		ID:        id,
		Peers:     make(map[string]*Device),
		Threshold: threshold,
		Curve:     elliptic.P256(),
	}
	// Generate ECC key pair for the device
	GenerateECCKeys(device)
	return device
}

// Generate ECC key pair
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

// ReceiveCoefficients allows a device to receive polynomial coefficients from another device
func (device *Device) ReceiveCoefficients(senderID string, coefficients []*big.Int) {
	// Simulate receiving the coefficients
	// Here, you can add logic to handle the coefficients if needed
	println("Device", device.ID, "received coefficients from", senderID)
}
