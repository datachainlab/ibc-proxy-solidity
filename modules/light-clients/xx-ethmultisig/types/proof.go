package types

import (
	fmt "fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(addresses []common.Address, multiSig *MultiSignature, signBytes []byte) error {
	if len(addresses) != len(multiSig.Signatures) {
		return ErrInvalidSignatureCount
	}
	h := crypto.Keccak256(signBytes)
	for i, sig := range multiSig.Signatures {
		if len(sig) != crypto.SignatureLength {
			return fmt.Errorf("signature must be 65 bytes long")
		}
		if sig[crypto.RecoveryIDOffset] != 27 && sig[crypto.RecoveryIDOffset] != 28 {
			return fmt.Errorf("invalid Ethereum signature (V is not 27 or 28)")
		}
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1

		rpk, err := crypto.SigToPub(h, sig)
		if err != nil {
			return err
		}
		signer := crypto.PubkeyToAddress(*rpk)
		if addresses[i] != signer {
			return fmt.Errorf("signature does not match signer")
		}
	}
	return nil
}

// ClientStateSignBytes returns the sign bytes for verification of the
// client state.
func ClientStateSignBytes(
	cdc codec.BinaryCodec,
	height clienttypes.Height, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	clientState exported.ClientState,
) ([]byte, error) {
	clientStateBz, err := cdc.MarshalInterface(clientState)
	if err != nil {
		return nil, err
	}
	data := StateData{
		Path:  []byte(path.String()),
		Value: clientStateBz,
	}
	dataBz, err := cdc.Marshal(&data)
	if err != nil {
		return nil, err
	}
	signBytes := &SignBytes{
		Height:      Height(height),
		Timestamp:   timestamp,
		Diversifier: diversifier,
		DataType:    CLIENT,
		Data:        dataBz,
	}
	return cdc.Marshal(signBytes)
}

// ConsensusStateSignBytes returns the sign bytes for verification of the
// consensus state.
func ConsensusStateSignBytes(
	cdc codec.BinaryCodec,
	height clienttypes.Height, timestamp uint64,
	diversifier string,
	path commitmenttypes.MerklePath,
	consensusState exported.ConsensusState,
) ([]byte, error) {
	consensusStateBz, err := cdc.MarshalInterface(consensusState)
	if err != nil {
		return nil, err
	}
	data := StateData{
		Path:  []byte(path.String()),
		Value: consensusStateBz,
	}
	dataBz, err := cdc.Marshal(&data)
	if err != nil {
		return nil, err
	}
	signBytes := &SignBytes{
		Height:      Height(height),
		Timestamp:   timestamp,
		Diversifier: diversifier,
		DataType:    CONSENSUS,
		Data:        dataBz,
	}
	return cdc.Marshal(signBytes)
}
