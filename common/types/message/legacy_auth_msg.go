package message

import (
	"crypto/ecdsa"

	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/common/hexutil"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/rlp"
)

// LegacyAuthMsg is the old auth message exchanged from the Prover to the Sequencer.
// It effectively acts as a registration, and makes the Prover identification
// known to the Sequencer.
type LegacyAuthMsg struct {
	// Message fields
	Identity *LegacyIdentity `json:"message"`
	// Prover signature
	Signature string `json:"signature"`
}

// LegacyIdentity contains all the fields to be signed by the prover.
type LegacyIdentity struct {
	// ProverName the prover name
	ProverName string `json:"prover_name"`
	// ProverVersion the prover version
	ProverVersion string `json:"prover_version"`
	// Challenge unique challenge generated by manager
	Challenge string `json:"challenge"`
}

// SignWithKey auth message with private key and set public key in auth message's Identity
func (a *LegacyAuthMsg) SignWithKey(priv *ecdsa.PrivateKey) error {
	// Hash identity content
	hash, err := a.Identity.Hash()
	if err != nil {
		return err
	}

	// Sign register message
	sig, err := crypto.Sign(hash, priv)
	if err != nil {
		return err
	}
	a.Signature = hexutil.Encode(sig)

	return nil
}

// Verify verifies the message of auth.
func (a *LegacyAuthMsg) Verify() (bool, error) {
	hash, err := a.Identity.Hash()
	if err != nil {
		return false, err
	}
	sig := common.FromHex(a.Signature)

	pk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return false, err
	}
	return crypto.VerifySignature(crypto.CompressPubkey(pk), hash, sig[:len(sig)-1]), nil
}

// PublicKey return public key from signature
func (a *LegacyAuthMsg) PublicKey() (string, error) {
	hash, err := a.Identity.Hash()
	if err != nil {
		return "", err
	}
	sig := common.FromHex(a.Signature)
	// recover public key
	pk, err := crypto.SigToPub(hash, sig)
	if err != nil {
		return "", err
	}
	return common.Bytes2Hex(crypto.CompressPubkey(pk)), nil
}

// Hash returns the hash of the auth message, which should be the message used
// to construct the Signature.
func (i *LegacyIdentity) Hash() ([]byte, error) {
	byt, err := rlp.EncodeToBytes(i)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256Hash(byt)
	return hash[:], nil
}
