package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

// HexToECDSAPrivateKey converts hexa string to ECDSA private key.
func HexToECDSAPrivateKey(hexKey string) (*ecdsa.PrivateKey, error) {
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
