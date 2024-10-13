package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

// Prime number for finite field operations
var prime, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)

// Share struct representing a share in Shamir's Secret Sharing
type Share struct {
	X *big.Int
	Y *big.Int
}

// Generate and share the polynomial
func GenerateAndSharePolynomial(device *Device, degree int) ([]*big.Int, error) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	coefficients := make([]*big.Int, degree+1)
	secret := big.NewInt(int64(rand.Intn(1000))) // Each device generates its own secret
	coefficients[0] = secret                     // Constant term is the device's "secret"

	// Print the secret (constant term of the polynomial) for debugging
	fmt.Println("Device", device.ID, "generated secret:", secret)

	for i := 1; i <= degree; i++ {
		coeff := big.NewInt(int64(rand.Intn(1000))) // Random coefficients for the polynomial
		coefficients[i] = coeff
	}

	ShareCoefficientsWithPeers(device, coefficients)
	return coefficients, nil
}

// Share coefficients with peers (simulated network communication)
func ShareCoefficientsWithPeers(device *Device, coefficients []*big.Int) {
	for _, peer := range device.Peers {
		// Simulate sending coefficients to peers
		fmt.Printf("Device %s sends coefficients to Device %s\n", device.ID, peer.ID)
		peer.ReceiveCoefficients(device.ID, coefficients)
	}
}

// Generate shares by evaluating the polynomial at distinct x-values
func GenerateShares(coefficients []*big.Int, n int) []*Share {
	shares := make([]*Share, n)
	for i := 1; i <= n; i++ {
		x := big.NewInt(int64(i))
		y := EvaluatePolynomial(coefficients, x)
		shares[i-1] = &Share{X: x, Y: y}
	}
	return shares
}

// Evaluate the polynomial at a specific x-value
func EvaluatePolynomial(coefficients []*big.Int, x *big.Int) *big.Int {
	result := big.NewInt(0)
	for i := len(coefficients) - 1; i >= 0; i-- {
		result.Mul(result, x)
		result.Add(result, coefficients[i])
		result.Mod(result, prime)
	}
	return result
}
