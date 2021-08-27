package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "<%= MultisigClientAddress; %>"
	ProxyMultisigClientAddress = "<%= ProxyMultisigClientAddress; %>"
	IBCHostAddress = "<%= IBCHostAddress; %>"
)

type contractConfig struct{}

var Contract contractConfig

func (contractConfig) GetMultisigClientAddress() common.Address {
	return common.HexToAddress(MultisigClientAddress)
}

func (contractConfig) GetProxyMultisigClientAddress() common.Address {
	return common.HexToAddress(ProxyMultisigClientAddress)
}

func (contractConfig) GetIBCHostAddress() common.Address {
	return common.HexToAddress(IBCHostAddress)
}
