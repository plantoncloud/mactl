package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

const (
	SshPrivateKeyFilePermissions = 0600
	SshPublicKeyFilePermissions  = 0644
)

func newKey(keyPath, keyName string) error {
	if err := os.MkdirAll(keyPath, 0744); err != nil {
		return errors.Wrapf(err, "failed to ensure %s dir", keyPath)
	}
	privateKeyPath := filepath.Join(keyPath, keyName)
	// ed25519 algorithm is chosen for reasons explained on
	// https://gist.github.com/brennanMKE/8e09593ca4064deab59da807077d8f53
	if err := generateSaveEd25519(privateKeyPath); err != nil {
		return err
	}
	return nil
}

// generateSaveEd25519 generates and saves ed25519 keys to disk after
// encoding into PEM format
// copied from https://gist.github.com/rorycl/d300f3ab942fd79e6cc1f37db0c6260f
func generateSaveEd25519(privateKeyPath string) error {

	var (
		err   error
		b     []byte
		block *pem.Block
		pub   ed25519.PublicKey
		priv  ed25519.PrivateKey
	)

	pub, priv, err = ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Printf("Generation error : %s", err)
		os.Exit(1)
	}

	b, err = x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: b,
	}

	err = os.WriteFile(privateKeyPath, pem.EncodeToMemory(block), SshPrivateKeyFilePermissions)
	if err != nil {
		return err
	}

	// public key
	b, err = x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return err
	}

	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: b,
	}

	fileName := privateKeyPath + ".pub"
	err = os.WriteFile(fileName, pem.EncodeToMemory(block), SshPublicKeyFilePermissions)
	return err

}
