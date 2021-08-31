package ethmultisig

import (
	"crypto/ecdsa"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"

	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	ethmultisigtypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
)

type ETHMultisig struct {
	cdc         codec.ProtoCodecMarshaler
	diversifier string
	keys        []*ecdsa.PrivateKey
	prefix      []byte
}

func NewETHMultisig(cdc codec.ProtoCodecMarshaler, diversifier string, keys []*ecdsa.PrivateKey, prefix []byte) ETHMultisig {
	return ETHMultisig{cdc: cdc, diversifier: diversifier, keys: keys, prefix: prefix}
}

func (m ETHMultisig) Addresses() []common.Address {
	var addresses []common.Address
	for _, key := range m.keys {
		addr := crypto.PubkeyToAddress(key.PublicKey)
		addresses = append(addresses, addr)
	}
	return addresses
}

func (m ETHMultisig) GetTimestamp() uint64 {
	return uint64(time.Now().UnixNano())
}

func (m ETHMultisig) SignClientState(clientID string, height clienttypes.Height, clientState exported.ClientState) (*ethmultisigtypes.MultiSignature, error) {
	bz, err := m.cdc.MarshalInterface(clientState)
	if err != nil {
		return nil, err
	}
	path, err := ClientCommitmentKey(m.prefix, clientID)
	if err != nil {
		return nil, err
	}
	data, err := m.cdc.Marshal(&ethmultisigtypes.StateData{
		Path:  path,
		Value: bz,
	})
	if err != nil {
		return nil, err
	}

	ts := m.GetTimestamp()
	signBytes, err := m.cdc.Marshal(&ethmultisigtypes.SignBytes{
		Height:      ethmultisigtypes.Height(height),
		Timestamp:   ts,
		Diversifier: m.diversifier,
		DataType:    ethmultisigtypes.CLIENT,
		Data:        data,
	})
	if err != nil {
		return nil, err
	}
	signHash := gethcrypto.Keccak256(signBytes)

	proof := ethmultisigtypes.MultiSignature{Timestamp: ts}
	for _, key := range m.keys {
		sig, err := gethcrypto.Sign(signHash, key)
		if err != nil {
			return nil, err
		}
		proof.Signatures = append(proof.Signatures, sig)
	}
	return &proof, nil
}
