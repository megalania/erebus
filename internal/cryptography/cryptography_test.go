package cryptography_test

import (
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/megalania/erebus/internal/cryptography"
)

const bitSize = 4096

func TestGeneratePrivateRSAKey(t *testing.T) {
	privateRSAKey, err := cryptography.GeneratePrivateRSAKey(
		bitSize,
	)
	if err != nil {
		t.Fatal(
			err.Error(),
		)
	}
	privateASNKey, err := x509.MarshalPKCS8PrivateKey(
		privateRSAKey,
	)
	if err != nil {
		t.Fatal(
			err.Error(),
		)
	}
	privatePEMKey := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PRIVATE KEY",
			Bytes: privateASNKey,
		},
	)
	t.Logf(
		"privatePEMKey: %s",
		privatePEMKey,
	)
}

func TestGetPublicRSAKey(t *testing.T) {
	privateRSAKey, err := cryptography.GeneratePrivateRSAKey(
		bitSize,
	)
	if err != nil {
		t.Fatal(
			err.Error(),
		)
	}
	publicRSAKey := cryptography.GetPublicRSAKey(
		privateRSAKey,
	)
	publicASNKey, err := x509.MarshalPKIXPublicKey(
		publicRSAKey,
	)
	if err != nil {
		t.Fatal(
			err.Error(),
		)
	}
	publicPEMKey := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PUBLIC KEY",
			Bytes: publicASNKey,
		},
	)
	t.Logf(
		"publicPEMKey: %s",
		publicPEMKey,
	)
}
