package main

import (
	"fmt"
	"math/big"
)

// Reconstruct the group key using Lagrange interpolation
func ReconstructGroupKey(shares []*Share) *big.Int {
	secret := big.NewInt(0)
	for i, share := range shares {
		numerator := big.NewInt(1)
		denominator := big.NewInt(1)
		for j, otherShare := range shares {
			if i != j {
				numerator.Mul(numerator, new(big.Int).Neg(otherShare.X)) // -xj
				numerator.Mod(numerator, prime)
				diff := new(big.Int).Sub(share.X, otherShare.X)
				denominator.Mul(denominator, diff) // xi - xj
				denominator.Mod(denominator, prime)
			}
		}
		lagrange := new(big.Int).Mul(numerator, ModInverse(denominator, prime))
		lagrange.Mod(lagrange, prime)
		temp := new(big.Int).Mul(share.Y, lagrange)
		temp.Mod(temp, prime)
		secret.Add(secret, temp)
		secret.Mod(secret, prime)
	}
	fmt.Println("Reconstructed Group Key:", secret)
	return secret // Reconstructed group key (constant term of the polynomial)
}

// Compute modular inverse (used in Lagrange interpolation)
func ModInverse(a, p *big.Int) *big.Int {
	return new(big.Int).ModInverse(a, p)
}

// Simulate device failure and still reconstruct the group key
func HandleDeviceFailure(devices []*Device, threshold int) *big.Int {
	validShares := []*Share{}
	for _, device := range devices {
		if device.Share != nil {
			validShares = append(validShares, device.Share)
		}
		if len(validShares) >= threshold {
			break
		}
	}

	if len(validShares) < threshold {
		fmt.Println("Not enough valid shares to reconstruct the key!")
		return big.NewInt(0)
	}

	// Reconstruct the group key from the available shares
	return ReconstructGroupKey(validShares)
}
