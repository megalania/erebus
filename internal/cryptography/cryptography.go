package cryptography

import (
	"crypto/rand"
	"crypto/rsa"
)

func GeneratePrivateRSAKey(bitSize int) (*rsa.PrivateKey, error) {
	privateRSAKey, err := rsa.GenerateKey(
		rand.Reader,
		bitSize,
	)
	if err != nil {
		return nil, err
	}
	err = privateRSAKey.Validate()
	if err != nil {
		return nil, err
	}
	return privateRSAKey, nil
}

func GetPublicRSAKey(privateRSAKey *rsa.PrivateKey) *rsa.PublicKey {
	return privateRSAKey.Public().(*rsa.PublicKey)
}
