package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "0x906083bF79339343Ba9Da6Ee6FD4c446720Daf0C"
	ProxyMultisigClientAddress = "0x98Bb2d64238474552ABe3181E31C9CC4eCAeBeA3"
	IBCHostAddress = "0x180B6C325525dB54C0DA871F1fd924a0bcf06397"
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
