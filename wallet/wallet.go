package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey        *ecdsa.PrivateKey `json:"private_key"`
	PublicKey         *ecdsa.PublicKey  `json:"public_key"`
	BlockchainAddress string            `json:"blockchainaddress"` //address of the user in the blockchain network

}

func NewWallet() (*Wallet, error) {
	w := new(Wallet)
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return &Wallet{}, err
	}
	w.PrivateKey = privateKey
	w.PublicKey = &privateKey.PublicKey
	w.createWalletAddress()

	return w, nil
}

// create wallet address
func (w *Wallet) createWalletAddress() {
	h2 := sha256.New()
	h2.Write(w.PublicKey.X.Bytes())
	h2.Write(w.PublicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	chsum := digest6[:4]

	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], chsum[:])

	walletAddress := base58.Encode(dc8)
	w.BlockchainAddress = walletAddress

}

func (w *Wallet) WalletAddress() string {
	return w.BlockchainAddress
}

func (w *Wallet) WalletPrivateKey() *ecdsa.PrivateKey {
	return w.PrivateKey
}

func (w *Wallet) WalletPrivateKeyToString() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) WalletPublicKey() *ecdsa.PublicKey {
	return w.PublicKey
}

func (w *Wallet) WalletPublicKeyToString() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}
