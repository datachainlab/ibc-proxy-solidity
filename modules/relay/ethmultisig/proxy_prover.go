package ethmultisig

import (
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibcexported "github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/datachainlab/ibc-proxy-prover/pkg/proxy"
	tendermintproxy "github.com/datachainlab/ibc-proxy-prover/pkg/proxy/tendermint"
)

type ProxyChainProver struct {
	proxyChain *tendermintproxy.TendermintProxyChain
	*Prover
}

var _ proxy.ProxyChainProverI = (*ProxyChainProver)(nil)

func (pr *ProxyChainProver) QueryProxyClientStateWithProof(height int64) (*clienttypes.QueryClientStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (pr *ProxyChainProver) QueryProxyClientConsensusStateWithProof(height int64, dstClientConsHeight ibcexported.Height) (*clienttypes.QueryConsensusStateResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (pr *ProxyChainProver) QueryProxyConnectionStateWithProof(height int64) (*conntypes.QueryConnectionResponse, error) {
	panic("not implemented") // TODO: Implement
}

func (pr *ProxyChainProver) QueryProxyChannelWithProof(height int64) (chanRes *chantypes.QueryChannelResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (pr *ProxyChainProver) QueryProxyPacketCommitmentWithProof(height int64, seq uint64) (comRes *chantypes.QueryPacketCommitmentResponse, err error) {
	panic("not implemented") // TODO: Implement
}

func (pr *ProxyChainProver) QueryProxyPacketAcknowledgementCommitmentWithProof(height int64, seq uint64) (ackRes *chantypes.QueryPacketAcknowledgementResponse, err error) {
	panic("not implemented") // TODO: Implement
}
