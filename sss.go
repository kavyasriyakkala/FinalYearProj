package main

import (
    "crypto/rand"
    "math/big"
)

// Prime number for finite field operations
var prime, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEFFFFFC2F", 16)

func GeneratePolynomial(secret *big.Int, degree int) ([]*big.Int, error) {
    coefficients := make([]*big.Int, degree+1)
    coefficients[0] = secret
    for i := 1; i <= degree; i++ {
        coeff, err := rand.Int(rand.Reader, prime)
        if err != nil {
            return nil, err
        }
        coefficients[i] = coeff
    }
    return coefficients, nil
}

func GenerateShares(coefficients []*big.Int, n int) ([]*Share, error) {
    shares := make([]*Share, n)
    for i := 1; i <= n; i++ {
        x := big.NewInt(int64(i))
        y := EvaluatePolynomial(coefficients, x)
        shares[i-1] = &Share{X: x, Y: y}
    }
    return shares, nil
}

func EvaluatePolynomial(coefficients []*big.Int, x *big.Int) *big.Int {
    result := big.NewInt(0)
    for i := len(coefficients) - 1; i >= 0; i-- {
        result.Mul(result, x)
        result.Add(result, coefficients[i])
        result.Mod(result, prime)
    }
    return result
}

func ReconstructSecret(shares []*Share) *big.Int {
    secret := big.NewInt(0)
    for i, share := range shares {
        numerator := big.NewInt(1)
        denominator := big.NewInt(1)
        for j, otherShare := range shares {
            if i != j {
                numerator.Mul(numerator, new(big.Int).Neg(otherShare.X))
                numerator.Mod(numerator, prime)
                diff := new(big.Int).Sub(share.X, otherShare.X)
                denominator.Mul(denominator, diff)
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
    return secret
}

func ModInverse(a, p *big.Int) *big.Int {
    return new(big.Int).ModInverse(a, p)
}
