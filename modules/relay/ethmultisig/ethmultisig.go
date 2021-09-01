package ethmultisig

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"

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

// GetCurrentTimestamp returns current time
func (m ETHMultisig) GetCurrentTimestamp() uint64 {
	return uint64(time.Now().UnixNano())
}

func (m ETHMultisig) SignConsensusState(height clienttypes.Height, clientID string, dstClientConsHeight ibcexported.Height, consensusState exported.ConsensusState) (*ethmultisigtypes.MultiSignature, error) {
	bz, err := m.cdc.MarshalInterface(consensusState)
	if err != nil {
		return nil, err
	}
	path, err := ConsensusCommitmentKey(m.prefix, clientID, dstClientConsHeight.GetRevisionHeight())
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.CONSENSUS, path, bz)
}

func (m ETHMultisig) SignClientState(height clienttypes.Height, clientID string, clientState exported.ClientState) (*ethmultisigtypes.MultiSignature, error) {
	bz, err := m.cdc.MarshalInterface(clientState)
	if err != nil {
		return nil, err
	}
	path, err := ClientCommitmentKey(m.prefix, clientID)
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.CLIENT, path, bz)
}

func (m ETHMultisig) SignConnectionState(height clienttypes.Height, connectionID string, connection conntypes.ConnectionEnd) (*ethmultisigtypes.MultiSignature, error) {
	bz, err := m.cdc.Marshal(&connection)
	if err != nil {
		return nil, err
	}
	path, err := ConnectionCommitmentKey(m.prefix, connectionID)
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.CONNECTION, path, bz)
}

func (m ETHMultisig) SignChannelState(height clienttypes.Height, portID, channelID string, channel chantypes.Channel) (*ethmultisigtypes.MultiSignature, error) {
	bz, err := m.cdc.Marshal(&channel)
	if err != nil {
		return nil, err
	}
	path, err := ChannelCommitmentKey(m.prefix, portID, channelID)
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.CHANNEL, path, bz)
}

func (m ETHMultisig) SignPacketState(height clienttypes.Height, portID, channelID string, sequence uint64, packetCommitment []byte) (*ethmultisigtypes.MultiSignature, error) {
	if len(packetCommitment) != 32 {
		return nil, fmt.Errorf("packetCommitment length must be 32")
	}
	path, err := PacketCommitmentKey(m.prefix, portID, channelID, sequence)
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.PACKETCOMMITMENT, path, packetCommitment)
}

func (m ETHMultisig) SignPacketAcknowledgementState(height clienttypes.Height, portID, channelID string, sequence uint64, acknowledgementCommitment []byte) (*ethmultisigtypes.MultiSignature, error) {
	if len(acknowledgementCommitment) != 32 {
		return nil, fmt.Errorf("acknowledgementCommitment length must be 32")
	}
	path, err := PacketAcknowledgementCommitmentKey(m.prefix, portID, channelID, sequence)
	if err != nil {
		return nil, err
	}
	return m.SignState(ethmultisigtypes.Height(height), ethmultisigtypes.PACKETACKNOWLEDGEMENT, path, acknowledgementCommitment)
}

func (m ETHMultisig) SignState(height ethmultisigtypes.Height, dtp ethmultisigtypes.SignBytes_DataType, path, value []byte) (*ethmultisigtypes.MultiSignature, error) {
	data, err := m.cdc.Marshal(&ethmultisigtypes.StateData{
		Path:  path,
		Value: value,
	})
	if err != nil {
		return nil, err
	}
	ts := m.GetCurrentTimestamp()
	signBytes, err := m.cdc.Marshal(&ethmultisigtypes.SignBytes{
		Height:      height,
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
