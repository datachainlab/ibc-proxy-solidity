package ethmultisig

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	commitmenttypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	host "github.com/cosmos/ibc-go/modules/core/24-host"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/datachainlab/ibc-proxy-prover/pkg/proxy"
	proxytypes "github.com/datachainlab/ibc-proxy/modules/light-clients/xx-proxy/types"
	ibcproxytypes "github.com/datachainlab/ibc-proxy/modules/proxy/types"
	"github.com/hyperledger-labs/yui-relayer/core"
)

var _ proxy.ProxyChainProverConfigI = (*ProxyChainProverConfig)(nil)

func (c *ProxyChainProverConfig) Build(proxyChain proxy.ProxyChainI) (proxy.ProxyChainProverI, error) {
	return NewProxyChainProver(c, proxyChain)
}

type ProxyChainProver struct {
	proxyChain proxy.ProxyChainI
	*Prover
}

var _ proxy.ProxyChainProverI = (*ProxyChainProver)(nil)

func NewProxyChainProver(cfg *ProxyChainProverConfig, proxyChain proxy.ProxyChainI) (*ProxyChainProver, error) {
	pr, err := NewProver(*cfg.ProverConfig, proxyChain)
	if err != nil {
		return nil, err
	}
	return &ProxyChainProver{
		proxyChain: proxyChain,
		Prover:     pr,
	}, nil
}

func (pr *ProxyChainProver) CreateMsgCreateClient(clientID string, dstHeader core.HeaderI, signer sdk.AccAddress) (*clienttypes.MsgCreateClient, error) {
	msg, err := pr.Prover.CreateMsgCreateClient(clientID, dstHeader, signer)
	if err != nil {
		return nil, err
	}
	ibcPrefix := commitmenttypes.NewMerklePrefix([]byte(host.StoreKey))
	proxyPrefix := commitmenttypes.NewMerklePrefix([]byte(ibcproxytypes.StoreKey))
	clientState := &proxytypes.ClientState{
		UpstreamClientId: pr.proxyChain.ProxyPath().UpstreamClientID,
		ProxyClientState: msg.ClientState,
		IbcPrefix:        &ibcPrefix,
		ProxyPrefix:      &proxyPrefix,
		TrustedSetup:     true,
	}
	anyClientState, err := clienttypes.PackClientState(clientState)
	if err != nil {
		return nil, err
	}
	consensusState := proxytypes.NewConsensusState(msg.ConsensusState)
	anyConsensusState, err := clienttypes.PackConsensusState(consensusState)
	if err != nil {
		return nil, err
	}
	return &clienttypes.MsgCreateClient{
		ClientState:    anyClientState,
		ConsensusState: anyConsensusState,
		Signer:         msg.Signer,
	}, nil
}

func (pr *ProxyChainProver) QueryProxyClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	res, err := pr.proxyChain.QueryProxyClientState(height)
	if err != nil {
		return nil, err
	}
	return pr.SignClientStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	res, err := pr.proxyChain.QueryProxyClientConsensusState(height, dstClientConsHeight)
	if err != nil {
		return nil, err
	}
	return pr.SignConsensusStateResponse(res, dstClientConsHeight)
}

func (pr *ProxyChainProver) QueryProxyConnectionStateWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	res, err := pr.proxyChain.QueryProxyConnectionState(height)
	if err != nil {
		return nil, err
	}
	return pr.SignConnectionStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	res, err := pr.proxyChain.QueryProxyChannel(height)
	if err != nil {
		return nil, err
	}
	return pr.SignChannelStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	res, err := pr.proxyChain.QueryProxyPacketCommitment(height, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignPacketStateResponse(res, seq)
}

func (pr *ProxyChainProver) QueryProxyPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	res, err := pr.proxyChain.QueryProxyPacketAcknowledgementCommitment(height, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignAcknowledgementStateResponse(res, seq)
}
