package chatglm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func rsaEncode(text []byte, publicKey string) (string, error) {
	pubPEM, _ := base64.StdEncoding.DecodeString(publicKey)
	_, block := pem.Decode(pubPEM)
	if block == nil {
		return "", fmt.Errorf("failed to decode PEM block containing public key")
	}

	pub, err := x509.ParsePKIXPublicKey(block)
	if err != nil {
		return "", err
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("cannot assert type: publicKey is not of type *rsa.PublicKey")
	}

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, text)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}
