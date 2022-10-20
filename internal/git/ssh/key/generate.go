package key

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/leftbin/go-util/pkg/file"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	SshKeyFilePerm = 0400
)

func newKey(keyPath, keyName string) error {
	if err := os.MkdirAll(keyPath, 0744); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", keyPath)
	}
	pvtKeyPath := filepath.Join(keyPath, keyName)
	pubKeyPath := fmt.Sprintf("%s.pub", pvtKeyPath)
	bitSize := 4096

	pvtKey, err := generatePrivateKey(bitSize)
	if err != nil {
		return errors.Wrap(err, "failed to cre pvt key")
	}

	pubKeyBytes, err := generatePubKey(&pvtKey.PublicKey)
	if err != nil {
		return errors.Wrap(err, "failed to cre pub key")
	}

	pvtKeyBytes := encodePvtKeyToPEM(pvtKey)

	if err := writeKeyToFile(pvtKeyBytes, pvtKeyPath); err != nil {
		return errors.Wrapf(err, "failed to write pvt key to %s file", pvtKeyPath)
	}
	if err := writeKeyToFile(pubKeyBytes, pubKeyPath); err != nil {
		return errors.Wrapf(err, "failed to write pub key to %s file", pubKeyPath)
	}
	return nil
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func generatePrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	// Private Key generation
	pvtKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate key")
	}
	// Validate Private Key
	err = pvtKey.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "failed to validate key")
	}
	return pvtKey, nil
}

// encodePvtKeyToPEM encodes Private Key from RSA to PEM format
func encodePvtKeyToPEM(pvtKey *rsa.PrivateKey) []byte {
	// Get ASN.1 DER format
	provider := x509.MarshalPKCS1PrivateKey(pvtKey)
	// pem.Block
	pvtBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   provider,
	}
	// Private key in PEM format
	pvtPem := pem.EncodeToMemory(&pvtBlock)
	return pvtPem
}

// generatePubKey take a rsa.PublicKey and return bytes suitable for writing to .pub file
// returns in the format "ssh-rsa ..."
func generatePubKey(pvtKey *rsa.PublicKey) ([]byte, error) {
	pubRsaKey, err := ssh.NewPublicKey(pvtKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get new pub key")
	}
	pubKeyBytes := ssh.MarshalAuthorizedKey(pubRsaKey)
	return pubKeyBytes, nil
}

// writePemToFile writes keys to a file
func writeKeyToFile(keyBytes []byte, saveFileTo string) error {
	if file.IsFileExists(saveFileTo) {
		if err := os.Remove(saveFileTo); err != nil {
			return errors.Wrapf(err, "failed to remove %s file", saveFileTo)
		}
	}
	if err := ioutil.WriteFile(saveFileTo, keyBytes, SshKeyFilePerm); err != nil {
		return errors.Wrapf(err, "failed to write file %s", saveFileTo)
	}
	return nil
}
