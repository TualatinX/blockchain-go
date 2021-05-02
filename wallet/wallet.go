package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

const (
	checksumLength = 4

	//hexadecimal representation of 0
	version = byte(0x00)
)

type Wallet struct {
	//ecdsa = elliptic curve digital signature algorithm
	PrivateKey ecdsa.PrivateKey

	PublicKey []byte
}

func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}

	pub := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pub
}
