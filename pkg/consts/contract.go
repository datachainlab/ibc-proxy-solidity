package consts

import (
	"github.com/ethereum/go-ethereum/common"
)

const (
	MultisigClientAddress = "0x6dE8E613E024De9FEE0759A016ff9BA80b668bFf"
	ProxyMultisigClientAddress = "0x4dc66A31643AA3f20E0D36CA21D7E0f24FbB63ca"
	IBCHostAddress = "0xdBB3621D708E159995754D3816f19757a8Cc945F"
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
