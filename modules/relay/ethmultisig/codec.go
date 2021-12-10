package ethmultisig

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/datachainlab/ibc-proxy-relay/pkg/proxy"
	"github.com/hyperledger-labs/yui-relayer/core"

	ethmultisigtypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
)

// RegisterInterfaces register the module interfaces to protobuf
// Any.
func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	ethmultisigtypes.RegisterInterfaces(registry)
	registry.RegisterImplementations(
		(*core.ProverConfigI)(nil),
		&ProverConfig{},
	)
	registry.RegisterImplementations(
		(*proxy.ProxyChainProverConfigI)(nil),
		&ProxyChainProverConfig{},
	)
}
