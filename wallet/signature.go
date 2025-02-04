package wallet

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (t *Transaction) GenerateSignature() (*Signature, error) {
	byteSlice, err := json.Marshal(t)
	if err != nil {
		return &Signature{}, err
	}
	hash := sha256.Sum256(byteSlice)
	r, s, err := ecdsa.Sign(rand.Reader, t.SenderPrivateKey, hash[:])
	if err != nil {
		return &Signature{}, err
	}
	return &Signature{R: r, S: s}, nil

}

func (s *Signature) String() string {
	return fmt.Sprintf("%x%x", s.R, s.S)
}
