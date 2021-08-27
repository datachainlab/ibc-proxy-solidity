package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "0x20889941d4ed2D08AEd24B6470eE0b38a06EC89C"
	ProxyMultisigClientAddress = "0xeeC744224f08b37E026458e4C7603577741CC93A"
	IBCHostAddress = "0xE6ea3b07bB7a205CF9fBf219B7C37dE74Ee81490"
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
