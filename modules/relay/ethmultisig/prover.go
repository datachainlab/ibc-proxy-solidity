package ethmultisig

import (
	"crypto/ecdsa"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	ethmultisigclient "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
	"github.com/datachainlab/ibc-proxy-solidity/modules/relay/ethmultisig/wallet"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ core.ProverConfigI = (*ProverConfig)(nil)

func (pr ProverConfig) Build(chain core.ChainI) (core.ProverI, error) {
	if len(pr.Wallets) == 0 {
		return nil, fmt.Errorf("at least one wallet is needed")
	}
	var prover Prover
	prover.diversifier = pr.Diversifier
	for _, w := range pr.Wallets {
		prv, err := wallet.GetPrvKeyFromMnemonicAndHDWPath(w.HdwPath, w.Mnemonic)
		if err != nil {
			return nil, err
		}
		prover.keys = append(prover.keys, prv)
	}
	return &prover, nil
}

type Prover struct {
	chain core.ChainI

	diversifier string
	keys        []*ecdsa.PrivateKey
}

var _ core.ProverI = (*Prover)(nil)

// GetChainID returns the chain ID
func (pr *Prover) GetChainID() string {
	return pr.chain.ChainID()
}

// QueryLatestHeader returns the latest header from the chain
func (pr *Prover) QueryLatestHeader() (out core.HeaderI, err error) {
	return &ethmultisigclient.Header{Height: ethmultisigclient.Height{RevisionNumber: 0, RevisionHeight: 1}}, nil
}

// GetLatestLightHeight returns the latest height on the light client
func (pr *Prover) GetLatestLightHeight() (int64, error) {
	return 1, nil
}

// CreateMsgCreateClient creates a CreateClientMsg to this chain
func (pr *Prover) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	var addresses [][]byte
	for _, key := range pr.keys {
		addr := crypto.PubkeyToAddress(key.PublicKey)
		addresses = append(addresses, addr.Bytes())
	}
	clientState := &ethmultisigclient.ClientState{
		LatestHeight: ethmultisigclient.Height{RevisionNumber: 0, RevisionHeight: 1},
	}
	consensusState := &ethmultisigclient.ConsensusState{
		Addresses:   addresses,
		Diversifier: pr.diversifier,
		Timestamp:   0,
	}
	return clienttypes.NewMsgCreateClient(clientState, consensusState, signer.String())
}

// SetupHeader creates a new header based on a given header
func (pr *Prover) SetupHeader(dst core.LightClientIBCQueryierI, baseSrcHeader core.HeaderI) (core.HeaderI, error) {
	return nil, nil
}

// UpdateLightWithHeader updates a header on the light client and returns the header and height corresponding to the chain
func (pr *Prover) UpdateLightWithHeader() (header core.HeaderI, provableHeight int64, queryableHeight int64, err error) {
	h, err := pr.QueryLatestHeader()
	if err != nil {
		return nil, 0, 0, err
	}
	return h, 0, 0, nil
}

// QueryClientConsensusState returns the ClientConsensusState and its proof
func (pr *Prover) QueryClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryClientStateWithProof returns the ClientState and its proof
func (pr *Prover) QueryClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryConnectionWithProof returns the Connection and its proof
func (pr *Prover) QueryConnectionWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	panic("not implemented") // TODO: Implement
}

// QueryChannelWithProof returns the Channel and its proof
func (pr *Prover) QueryChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	panic("not implemented") // TODO: Implement
}

// QueryPacketCommitmentWithProof returns the packet commitment and its proof
func (pr *Prover) QueryPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	panic("not implemented") // TODO: Implement
}

// QueryPacketAcknowledgementCommitmentWithProof returns the packet acknowledgement commitment and its proof
func (pr *Prover) QueryPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	panic("not implemented") // TODO: Implement
}
