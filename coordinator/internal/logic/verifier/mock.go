//go:build mock_verifier

package verifier

import (
	"scroll-tech/common/types/message"

	"scroll-tech/coordinator/internal/config"
)

// NewVerifier Sets up a mock verifier.
func NewVerifier(cfg *config.VerifierConfig) (*Verifier, error) {
	batchVKMap := map[string]string{
		"shanghai":  "",
		"bernoulli": "",
		"london":    "",
		"istanbul":  "",
		"homestead": "",
		"eip155":    "",
	}
	chunkVKMap := map[string]string{
		"shanghai":  "",
		"bernoulli": "",
		"london":    "",
		"istanbul":  "",
		"homestead": "",
		"eip155":    "",
	}
	batchVKMap[cfg.ForkName] = ""
	chunkVKMap[cfg.ForkName] = ""
	return &Verifier{cfg: cfg, ChunkVKMap: chunkVKMap, BatchVKMap: batchVKMap}, nil
}

// VerifyChunkProof return a mock verification result for a ChunkProof.
func (v *Verifier) VerifyChunkProof(proof *message.ChunkProof) (bool, error) {
	if string(proof.Proof) == InvalidTestProof {
		return false, nil
	}
	return true, nil
}

// VerifyBatchProof return a mock verification result for a BatchProof.
func (v *Verifier) VerifyBatchProof(proof *message.BatchProof, forkName string) (bool, error) {
	if string(proof.Proof) == InvalidTestProof {
		return false, nil
	}
	return true, nil
}
