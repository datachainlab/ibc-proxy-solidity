package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "0x361552A65C96621003C62C5971b910a1fdC9ba78"
	ProxyMultisigClientAddress = "0x9eBF3956EE45B2b9F1fC85FB8990ce6be52F47a6"
	IBCHostAddress = "0x727A5648832D2b317925CE043eA9b7fE04B4CD55"
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
