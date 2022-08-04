package encoded_assets

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
	"github.com/algorand/go-algorand-sdk/stateproofs/stateprooftypes"
	"github.com/algorand/go-algorand-sdk/types"
)

// These functions take the encoded assets, committed as examples, and parse them.

func decodeFromFile(encodedPath string, target interface{}) error {
	encodedData, err := os.ReadFile(encodedPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(encodedData, target)
	return err
}

func GetParsedTransactionVerificationData(transactionVerificationDataPath string) (types.Digest, types.Round, stateprooftypes.Seed, types.Digest, models.ProofResponse,
	models.LightBlockHeaderProof, error) {
	var genesisHash types.Digest
	err := decodeFromFile(filepath.Join(transactionVerificationDataPath, "genesis_hash.json"), &genesisHash)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	var round types.Round
	err = decodeFromFile(filepath.Join(transactionVerificationDataPath, "round.json"), &round)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	var seed stateprooftypes.Seed
	err = decodeFromFile(filepath.Join(transactionVerificationDataPath, "seed.json"), &seed)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	var transactionId types.Digest
	err = decodeFromFile(filepath.Join(transactionVerificationDataPath, "transaction_id.json"), &transactionId)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	var transactionProofResponse models.ProofResponse
	err = decodeFromFile(filepath.Join(transactionVerificationDataPath, "transaction_proof_response.json"), &transactionProofResponse)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	var lightBlockHeaderProof models.LightBlockHeaderProof
	err = decodeFromFile(filepath.Join(transactionVerificationDataPath, "light_block_header_proof_response.json"), &lightBlockHeaderProof)
	if err != nil {
		return types.Digest{}, 0, stateprooftypes.Seed{}, types.Digest{}, models.ProofResponse{}, models.LightBlockHeaderProof{}, err
	}

	return genesisHash, round, seed, transactionId, transactionProofResponse,
		lightBlockHeaderProof, nil
}

func GetParsedGenesisData(genesisDataPath string) (stateprooftypes.GenericDigest, uint64, error) {
	genesisVotersCommitment := stateprooftypes.GenericDigest{}
	err := decodeFromFile(filepath.Join(genesisDataPath, "genesis_voters_commitment.json"), &genesisVotersCommitment)
	if err != nil {
		return stateprooftypes.GenericDigest{}, 0, err
	}

	genesisVotersLnProvenWeight := uint64(0)
	err = decodeFromFile(filepath.Join(genesisDataPath, "genesis_voters_ln_proven_weight.json"), &genesisVotersLnProvenWeight)
	if err != nil {
		return stateprooftypes.GenericDigest{}, 0, err
	}

	return genesisVotersCommitment, genesisVotersLnProvenWeight, nil
}

func GetParsedStateProofAdvancmentData(stateProofVerificationDataPath string) (stateprooftypes.Message,
	*stateprooftypes.EncodedStateProof, error) {
	stateProofMessage := stateprooftypes.Message{}
	err := decodeFromFile(filepath.Join(stateProofVerificationDataPath, "state_proof_message.json"), &stateProofMessage)
	if err != nil {
		return stateprooftypes.Message{}, nil, err
	}

	var stateProof stateprooftypes.EncodedStateProof
	err = decodeFromFile(filepath.Join(stateProofVerificationDataPath, "state_proof.json"), &stateProof)
	if err != nil {
		return stateprooftypes.Message{}, nil, err
	}

	return stateProofMessage, &stateProof, nil
}
