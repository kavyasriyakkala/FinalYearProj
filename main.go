package main

import (
	"fmt"
)

const (
	numDevices = 5 // Number of devices
	threshold  = 3 // Threshold for reconstruction
)

var devices []*Device

func main() {
	// Initialize devices
	for i := 1; i <= numDevices; i++ {
		deviceID := fmt.Sprintf("Device%d", i)
		device := InitializeDevice(deviceID, threshold)
		devices = append(devices, device)
	}

	// Establish peer relationships (each device knows all other devices)
	for i, device := range devices {
		for j, peer := range devices {
			if i != j {
				device.Peers[peer.ID] = peer
			}
		}
	}

	// Each device generates its own polynomial and shares the coefficients with peers
	for _, device := range devices {
		coefficients, _ := GenerateAndSharePolynomial(device, threshold-1)
		shares := GenerateShares(coefficients, numDevices)
		DistributeShares(device, shares) // No need to pass GroupKey here
	}

	// Simulate group key reconstruction with a threshold of shares
	reconstructedKey := HandleDeviceFailure(devices, threshold)
	fmt.Println("Final Reconstructed Group Key:", reconstructedKey)
}
