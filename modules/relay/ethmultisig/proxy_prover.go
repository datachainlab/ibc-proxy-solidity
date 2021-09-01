package ethmultisig

import (
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/datachainlab/ibc-proxy-prover/pkg/proxy"
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

func (pr *ProxyChainProver) QueryProxyClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	res, err := pr.proxyChain.QueryProxyClientState(height)
	if err != nil {
		return nil, err
	}
	return pr.SignClientStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	res, err := pr.QueryClientConsensusStateWithProof(height, dstClientConsHeight)
	if err != nil {
		return nil, err
	}
	return pr.SignConsensusStateResponse(res, dstClientConsHeight)
}

func (pr *ProxyChainProver) QueryProxyConnectionStateWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	res, err := pr.QueryConnectionWithProof(height)
	if err != nil {
		return nil, err
	}
	return pr.SignConnectionStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	res, err := pr.QueryChannelWithProof(height)
	if err != nil {
		return nil, err
	}
	return pr.SignChannelStateResponse(res)
}

func (pr *ProxyChainProver) QueryProxyPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	res, err := pr.QueryPacketCommitmentWithProof(height, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignPacketStateResponse(res, seq)
}

func (pr *ProxyChainProver) QueryProxyPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	res, err := pr.QueryPacketAcknowledgementCommitmentWithProof(height, seq)
	if err != nil {
		return nil, err
	}
	return pr.SignAcknowledgementStateResponse(res, seq)
}
