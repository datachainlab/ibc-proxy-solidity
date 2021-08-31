// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxymultisigclient

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type ClientStateData struct {
	LatestHeight HeightData
	FrozenHeight HeightData
}

// ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type ConsensusStateData struct {
	Addresses   [][]byte
	Diversifier string
	Timestamp   uint64
}

// GoogleProtobufAnyData is an auto generated low-level Go binding around an user-defined struct.
type GoogleProtobufAnyData struct {
	TypeUrl string
	Value   []byte
}

// HeightData is an auto generated low-level Go binding around an user-defined struct.
type HeightData struct {
	RevisionNumber uint64
	RevisionHeight uint64
}

// IbcCoreCommitmentV1MerklePrefixData is an auto generated low-level Go binding around an user-defined struct.
type IbcCoreCommitmentV1MerklePrefixData struct {
	KeyPrefix []byte
}

// IbcLightclientsProxyV1ClientStateData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsProxyV1ClientStateData struct {
	UpstreamClientId string
	ProxyClientState GoogleProtobufAnyData
	ProxyPrefix      IbcCoreCommitmentV1MerklePrefixData
	IbcPrefix        IbcCoreCommitmentV1MerklePrefixData
	TrustedSetup     bool
}

// IbcLightclientsProxyV1ConsensusStateData is an auto generated low-level Go binding around an user-defined struct.
type IbcLightclientsProxyV1ConsensusStateData struct {
	ProxyConsensusState GoogleProtobufAnyData
}

// MultiSignatureData is an auto generated low-level Go binding around an user-defined struct.
type MultiSignatureData struct {
	Signatures [][]byte
	Timestamp  uint64
}

// ProxymultisigclientABI is the input ABI used to generate the binding from.
const ProxymultisigclientABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getLatestHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"name\":\"getTimestampAtHeight\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"portID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channel\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeChannelStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeClientStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"connectionID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connection\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeConnectionStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"clientID\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"consensusState\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makeConsensusStateSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"path\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"packetAcknowledgement\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makePacketAcknowledgementSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"path\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"packetCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"}],\"name\":\"makePacketSignBytes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"addresses\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structMultiSignature.Data\",\"name\":\"multisig\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signBytes\",\"type\":\"bytes\"}],\"name\":\"verifySignature\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerBytes\",\"type\":\"bytes\"}],\"name\":\"checkHeaderAndUpdateState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"newClientStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newConsensusStateBytes\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"clientStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"string\",\"name\":\"counterpartyClientIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"consensusHeight\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"consensusStateBytes\",\"type\":\"bytes\"}],\"name\":\"verifyClientConsensusState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"connectionId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"connectionBytes\",\"type\":\"bytes\"}],\"name\":\"verifyConnectionState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"channelBytes\",\"type\":\"bytes\"}],\"name\":\"verifyChannelState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"commitmentBytes\",\"type\":\"bytes32\"}],\"name\":\"verifyPacketCommitment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"height\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"prefix\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"proof\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"portId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"channelId\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"sequence\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"acknowledgement\",\"type\":\"bytes\"}],\"name\":\"verifyPacketAcknowledgement\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getProxyClientState\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"upstream_client_id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"type_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structGoogleProtobufAny.Data\",\"name\":\"proxy_client_state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structIbcCoreCommitmentV1MerklePrefix.Data\",\"name\":\"proxy_prefix\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes\",\"name\":\"key_prefix\",\"type\":\"bytes\"}],\"internalType\":\"structIbcCoreCommitmentV1MerklePrefix.Data\",\"name\":\"ibc_prefix\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"trusted_setup\",\"type\":\"bool\"}],\"internalType\":\"structIbcLightclientsProxyV1ClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getProxyConsensusState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"type_url\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"internalType\":\"structGoogleProtobufAny.Data\",\"name\":\"proxy_consensus_state\",\"type\":\"tuple\"}],\"internalType\":\"structIbcLightclientsProxyV1ConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"}],\"name\":\"getClientState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"latest_height\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"frozen_height\",\"type\":\"tuple\"}],\"internalType\":\"structClientState.Data\",\"name\":\"clientState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"contractIBCHost\",\"name\":\"host\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"clientId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"revision_number\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"revision_height\",\"type\":\"uint64\"}],\"internalType\":\"structHeight.Data\",\"name\":\"height\",\"type\":\"tuple\"}],\"name\":\"getConsensusState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes[]\",\"name\":\"addresses\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"diversifier\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"timestamp\",\"type\":\"uint64\"}],\"internalType\":\"structConsensusState.Data\",\"name\":\"consensusState\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"proxyKeyPrefix\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"upstreamClientId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"counterpartyPrefix\",\"type\":\"bytes\"}],\"name\":\"makeProxyCommitmentPrefix\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

// Proxymultisigclient is an auto generated Go binding around an Ethereum contract.
type Proxymultisigclient struct {
	ProxymultisigclientCaller     // Read-only binding to the contract
	ProxymultisigclientTransactor // Write-only binding to the contract
	ProxymultisigclientFilterer   // Log filterer for contract events
}

// ProxymultisigclientCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxymultisigclientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxymultisigclientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxymultisigclientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxymultisigclientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxymultisigclientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxymultisigclientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxymultisigclientSession struct {
	Contract     *Proxymultisigclient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ProxymultisigclientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxymultisigclientCallerSession struct {
	Contract *ProxymultisigclientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ProxymultisigclientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxymultisigclientTransactorSession struct {
	Contract     *ProxymultisigclientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ProxymultisigclientRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxymultisigclientRaw struct {
	Contract *Proxymultisigclient // Generic contract binding to access the raw methods on
}

// ProxymultisigclientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxymultisigclientCallerRaw struct {
	Contract *ProxymultisigclientCaller // Generic read-only contract binding to access the raw methods on
}

// ProxymultisigclientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxymultisigclientTransactorRaw struct {
	Contract *ProxymultisigclientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxymultisigclient creates a new instance of Proxymultisigclient, bound to a specific deployed contract.
func NewProxymultisigclient(address common.Address, backend bind.ContractBackend) (*Proxymultisigclient, error) {
	contract, err := bindProxymultisigclient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxymultisigclient{ProxymultisigclientCaller: ProxymultisigclientCaller{contract: contract}, ProxymultisigclientTransactor: ProxymultisigclientTransactor{contract: contract}, ProxymultisigclientFilterer: ProxymultisigclientFilterer{contract: contract}}, nil
}

// NewProxymultisigclientCaller creates a new read-only instance of Proxymultisigclient, bound to a specific deployed contract.
func NewProxymultisigclientCaller(address common.Address, caller bind.ContractCaller) (*ProxymultisigclientCaller, error) {
	contract, err := bindProxymultisigclient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxymultisigclientCaller{contract: contract}, nil
}

// NewProxymultisigclientTransactor creates a new write-only instance of Proxymultisigclient, bound to a specific deployed contract.
func NewProxymultisigclientTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxymultisigclientTransactor, error) {
	contract, err := bindProxymultisigclient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxymultisigclientTransactor{contract: contract}, nil
}

// NewProxymultisigclientFilterer creates a new log filterer instance of Proxymultisigclient, bound to a specific deployed contract.
func NewProxymultisigclientFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxymultisigclientFilterer, error) {
	contract, err := bindProxymultisigclient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxymultisigclientFilterer{contract: contract}, nil
}

// bindProxymultisigclient binds a generic wrapper to an already deployed contract.
func bindProxymultisigclient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxymultisigclientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxymultisigclient *ProxymultisigclientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxymultisigclient.Contract.ProxymultisigclientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxymultisigclient *ProxymultisigclientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxymultisigclient.Contract.ProxymultisigclientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxymultisigclient *ProxymultisigclientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxymultisigclient.Contract.ProxymultisigclientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxymultisigclient *ProxymultisigclientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxymultisigclient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxymultisigclient *ProxymultisigclientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxymultisigclient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxymultisigclient *ProxymultisigclientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxymultisigclient.Contract.contract.Transact(opts, method, params...)
}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address host, string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Proxymultisigclient *ProxymultisigclientCaller) CheckHeaderAndUpdateState(opts *bind.CallOpts, host common.Address, clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "checkHeaderAndUpdateState", host, clientId, clientStateBytes, headerBytes)

	outstruct := new(struct {
		NewClientStateBytes    []byte
		NewConsensusStateBytes []byte
		Height                 uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.NewClientStateBytes = out[0].([]byte)
	outstruct.NewConsensusStateBytes = out[1].([]byte)
	outstruct.Height = out[2].(uint64)

	return *outstruct, err

}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address host, string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Proxymultisigclient *ProxymultisigclientSession) CheckHeaderAndUpdateState(host common.Address, clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	return _Proxymultisigclient.Contract.CheckHeaderAndUpdateState(&_Proxymultisigclient.CallOpts, host, clientId, clientStateBytes, headerBytes)
}

// CheckHeaderAndUpdateState is a free data retrieval call binding the contract method 0xfe4391fd.
//
// Solidity: function checkHeaderAndUpdateState(address host, string clientId, bytes clientStateBytes, bytes headerBytes) view returns(bytes newClientStateBytes, bytes newConsensusStateBytes, uint64 height)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) CheckHeaderAndUpdateState(host common.Address, clientId string, clientStateBytes []byte, headerBytes []byte) (struct {
	NewClientStateBytes    []byte
	NewConsensusStateBytes []byte
	Height                 uint64
}, error) {
	return _Proxymultisigclient.Contract.CheckHeaderAndUpdateState(&_Proxymultisigclient.CallOpts, host, clientId, clientStateBytes, headerBytes)
}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetClientState(opts *bind.CallOpts, host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getClientState", host, clientId)

	outstruct := new(struct {
		ClientState ClientStateData
		Found       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClientState = out[0].(ClientStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientSession) GetClientState(host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	return _Proxymultisigclient.Contract.GetClientState(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetClientState is a free data retrieval call binding the contract method 0xff33b3e6.
//
// Solidity: function getClientState(address host, string clientId) view returns(((uint64,uint64),(uint64,uint64)) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetClientState(host common.Address, clientId string) (struct {
	ClientState ClientStateData
	Found       bool
}, error) {
	return _Proxymultisigclient.Contract.GetClientState(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetConsensusState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getConsensusState", host, clientId, height)

	outstruct := new(struct {
		ConsensusState ConsensusStateData
		Found          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusState = out[0].(ConsensusStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientSession) GetConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	return _Proxymultisigclient.Contract.GetConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// GetConsensusState is a free data retrieval call binding the contract method 0xdd8dd65f.
//
// Solidity: function getConsensusState(address host, string clientId, (uint64,uint64) height) view returns((bytes[],string,uint64) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState ConsensusStateData
	Found          bool
}, error) {
	return _Proxymultisigclient.Contract.GetConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetLatestHeight(opts *bind.CallOpts, host common.Address, clientId string) (uint64, bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getLatestHeight", host, clientId)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientSession) GetLatestHeight(host common.Address, clientId string) (uint64, bool, error) {
	return _Proxymultisigclient.Contract.GetLatestHeight(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetLatestHeight is a free data retrieval call binding the contract method 0xfe017750.
//
// Solidity: function getLatestHeight(address host, string clientId) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetLatestHeight(host common.Address, clientId string) (uint64, bool, error) {
	return _Proxymultisigclient.Contract.GetLatestHeight(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetProxyClientState is a free data retrieval call binding the contract method 0xa65c4e74.
//
// Solidity: function getProxyClientState(address host, string clientId) view returns((string,(string,bytes),(bytes),(bytes),bool) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetProxyClientState(opts *bind.CallOpts, host common.Address, clientId string) (struct {
	ClientState IbcLightclientsProxyV1ClientStateData
	Found       bool
}, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getProxyClientState", host, clientId)

	outstruct := new(struct {
		ClientState IbcLightclientsProxyV1ClientStateData
		Found       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ClientState = out[0].(IbcLightclientsProxyV1ClientStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetProxyClientState is a free data retrieval call binding the contract method 0xa65c4e74.
//
// Solidity: function getProxyClientState(address host, string clientId) view returns((string,(string,bytes),(bytes),(bytes),bool) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientSession) GetProxyClientState(host common.Address, clientId string) (struct {
	ClientState IbcLightclientsProxyV1ClientStateData
	Found       bool
}, error) {
	return _Proxymultisigclient.Contract.GetProxyClientState(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetProxyClientState is a free data retrieval call binding the contract method 0xa65c4e74.
//
// Solidity: function getProxyClientState(address host, string clientId) view returns((string,(string,bytes),(bytes),(bytes),bool) clientState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetProxyClientState(host common.Address, clientId string) (struct {
	ClientState IbcLightclientsProxyV1ClientStateData
	Found       bool
}, error) {
	return _Proxymultisigclient.Contract.GetProxyClientState(&_Proxymultisigclient.CallOpts, host, clientId)
}

// GetProxyConsensusState is a free data retrieval call binding the contract method 0x93b21d8a.
//
// Solidity: function getProxyConsensusState(address host, string clientId, (uint64,uint64) height) view returns(((string,bytes)) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetProxyConsensusState(opts *bind.CallOpts, host common.Address, clientId string, height HeightData) (struct {
	ConsensusState IbcLightclientsProxyV1ConsensusStateData
	Found          bool
}, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getProxyConsensusState", host, clientId, height)

	outstruct := new(struct {
		ConsensusState IbcLightclientsProxyV1ConsensusStateData
		Found          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ConsensusState = out[0].(IbcLightclientsProxyV1ConsensusStateData)
	outstruct.Found = out[1].(bool)

	return *outstruct, err

}

// GetProxyConsensusState is a free data retrieval call binding the contract method 0x93b21d8a.
//
// Solidity: function getProxyConsensusState(address host, string clientId, (uint64,uint64) height) view returns(((string,bytes)) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientSession) GetProxyConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState IbcLightclientsProxyV1ConsensusStateData
	Found          bool
}, error) {
	return _Proxymultisigclient.Contract.GetProxyConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// GetProxyConsensusState is a free data retrieval call binding the contract method 0x93b21d8a.
//
// Solidity: function getProxyConsensusState(address host, string clientId, (uint64,uint64) height) view returns(((string,bytes)) consensusState, bool found)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetProxyConsensusState(host common.Address, clientId string, height HeightData) (struct {
	ConsensusState IbcLightclientsProxyV1ConsensusStateData
	Found          bool
}, error) {
	return _Proxymultisigclient.Contract.GetProxyConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xbd90991a.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, uint64 height) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) GetTimestampAtHeight(opts *bind.CallOpts, host common.Address, clientId string, height uint64) (uint64, bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "getTimestampAtHeight", host, clientId, height)

	if err != nil {
		return *new(uint64), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xbd90991a.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, uint64 height) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientSession) GetTimestampAtHeight(host common.Address, clientId string, height uint64) (uint64, bool, error) {
	return _Proxymultisigclient.Contract.GetTimestampAtHeight(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// GetTimestampAtHeight is a free data retrieval call binding the contract method 0xbd90991a.
//
// Solidity: function getTimestampAtHeight(address host, string clientId, uint64 height) view returns(uint64, bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) GetTimestampAtHeight(host common.Address, clientId string, height uint64) (uint64, bool, error) {
	return _Proxymultisigclient.Contract.GetTimestampAtHeight(&_Proxymultisigclient.CallOpts, host, clientId, height)
}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x81b4e822.
//
// Solidity: function makeChannelStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakeChannelStateSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makeChannelStateSignBytes", height, timestamp, diversifier, portID, channelID, channel, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x81b4e822.
//
// Solidity: function makeChannelStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakeChannelStateSignBytes(height uint64, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeChannelStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, portID, channelID, channel, prefix)
}

// MakeChannelStateSignBytes is a free data retrieval call binding the contract method 0x81b4e822.
//
// Solidity: function makeChannelStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string portID, string channelID, bytes channel, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakeChannelStateSignBytes(height uint64, timestamp uint64, diversifier string, portID string, channelID string, channel []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeChannelStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, portID, channelID, channel, prefix)
}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0x464cb82f.
//
// Solidity: function makeClientStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakeClientStateSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makeClientStateSignBytes", height, timestamp, diversifier, clientID, clientState, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0x464cb82f.
//
// Solidity: function makeClientStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakeClientStateSignBytes(height uint64, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeClientStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, clientID, clientState, prefix)
}

// MakeClientStateSignBytes is a free data retrieval call binding the contract method 0x464cb82f.
//
// Solidity: function makeClientStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, bytes clientState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakeClientStateSignBytes(height uint64, timestamp uint64, diversifier string, clientID string, clientState []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeClientStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, clientID, clientState, prefix)
}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x4438375e.
//
// Solidity: function makeConnectionStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakeConnectionStateSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makeConnectionStateSignBytes", height, timestamp, diversifier, connectionID, connection, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x4438375e.
//
// Solidity: function makeConnectionStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakeConnectionStateSignBytes(height uint64, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeConnectionStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, connectionID, connection, prefix)
}

// MakeConnectionStateSignBytes is a free data retrieval call binding the contract method 0x4438375e.
//
// Solidity: function makeConnectionStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string connectionID, bytes connection, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakeConnectionStateSignBytes(height uint64, timestamp uint64, diversifier string, connectionID string, connection []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeConnectionStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, connectionID, connection, prefix)
}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x4125d9e9.
//
// Solidity: function makeConsensusStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, uint64 consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakeConsensusStateSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, clientID string, consensusHeight uint64, consensusState []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makeConsensusStateSignBytes", height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x4125d9e9.
//
// Solidity: function makeConsensusStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, uint64 consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakeConsensusStateSignBytes(height uint64, timestamp uint64, diversifier string, clientID string, consensusHeight uint64, consensusState []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeConsensusStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)
}

// MakeConsensusStateSignBytes is a free data retrieval call binding the contract method 0x4125d9e9.
//
// Solidity: function makeConsensusStateSignBytes(uint64 height, uint64 timestamp, string diversifier, string clientID, uint64 consensusHeight, bytes consensusState, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakeConsensusStateSignBytes(height uint64, timestamp uint64, diversifier string, clientID string, consensusHeight uint64, consensusState []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeConsensusStateSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, clientID, consensusHeight, consensusState, prefix)
}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x0b94538d.
//
// Solidity: function makePacketAcknowledgementSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakePacketAcknowledgementSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makePacketAcknowledgementSignBytes", height, timestamp, diversifier, path, packetAcknowledgement, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x0b94538d.
//
// Solidity: function makePacketAcknowledgementSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakePacketAcknowledgementSignBytes(height uint64, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakePacketAcknowledgementSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, path, packetAcknowledgement, prefix)
}

// MakePacketAcknowledgementSignBytes is a free data retrieval call binding the contract method 0x0b94538d.
//
// Solidity: function makePacketAcknowledgementSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes packetAcknowledgement, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakePacketAcknowledgementSignBytes(height uint64, timestamp uint64, diversifier string, path [32]byte, packetAcknowledgement []byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakePacketAcknowledgementSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, path, packetAcknowledgement, prefix)
}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xfc1b3803.
//
// Solidity: function makePacketSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakePacketSignBytes(opts *bind.CallOpts, height uint64, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makePacketSignBytes", height, timestamp, diversifier, path, packetCommitment, prefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xfc1b3803.
//
// Solidity: function makePacketSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakePacketSignBytes(height uint64, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakePacketSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, path, packetCommitment, prefix)
}

// MakePacketSignBytes is a free data retrieval call binding the contract method 0xfc1b3803.
//
// Solidity: function makePacketSignBytes(uint64 height, uint64 timestamp, string diversifier, bytes32 path, bytes32 packetCommitment, bytes prefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakePacketSignBytes(height uint64, timestamp uint64, diversifier string, path [32]byte, packetCommitment [32]byte, prefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakePacketSignBytes(&_Proxymultisigclient.CallOpts, height, timestamp, diversifier, path, packetCommitment, prefix)
}

// MakeProxyCommitmentPrefix is a free data retrieval call binding the contract method 0xeef3a01a.
//
// Solidity: function makeProxyCommitmentPrefix(bytes proxyKeyPrefix, string upstreamClientId, bytes counterpartyPrefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCaller) MakeProxyCommitmentPrefix(opts *bind.CallOpts, proxyKeyPrefix []byte, upstreamClientId string, counterpartyPrefix []byte) ([]byte, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "makeProxyCommitmentPrefix", proxyKeyPrefix, upstreamClientId, counterpartyPrefix)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MakeProxyCommitmentPrefix is a free data retrieval call binding the contract method 0xeef3a01a.
//
// Solidity: function makeProxyCommitmentPrefix(bytes proxyKeyPrefix, string upstreamClientId, bytes counterpartyPrefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientSession) MakeProxyCommitmentPrefix(proxyKeyPrefix []byte, upstreamClientId string, counterpartyPrefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeProxyCommitmentPrefix(&_Proxymultisigclient.CallOpts, proxyKeyPrefix, upstreamClientId, counterpartyPrefix)
}

// MakeProxyCommitmentPrefix is a free data retrieval call binding the contract method 0xeef3a01a.
//
// Solidity: function makeProxyCommitmentPrefix(bytes proxyKeyPrefix, string upstreamClientId, bytes counterpartyPrefix) pure returns(bytes)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) MakeProxyCommitmentPrefix(proxyKeyPrefix []byte, upstreamClientId string, counterpartyPrefix []byte) ([]byte, error) {
	return _Proxymultisigclient.Contract.MakeProxyCommitmentPrefix(&_Proxymultisigclient.CallOpts, proxyKeyPrefix, upstreamClientId, counterpartyPrefix)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0x6455eb57.
//
// Solidity: function verifyChannelState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyChannelState(opts *bind.CallOpts, host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyChannelState", host, clientId, height, prefix, proof, portId, channelId, channelBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyChannelState is a free data retrieval call binding the contract method 0x6455eb57.
//
// Solidity: function verifyChannelState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyChannelState(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyChannelState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyChannelState is a free data retrieval call binding the contract method 0x6455eb57.
//
// Solidity: function verifyChannelState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, bytes channelBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyChannelState(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, channelBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyChannelState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, channelBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xeb4561fb.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyClientConsensusState(opts *bind.CallOpts, host common.Address, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyClientConsensusState", host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xeb4561fb.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyClientConsensusState(host common.Address, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyClientConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientConsensusState is a free data retrieval call binding the contract method 0xeb4561fb.
//
// Solidity: function verifyClientConsensusState(address host, string clientId, uint64 height, string counterpartyClientIdentifier, uint64 consensusHeight, bytes prefix, bytes proof, bytes consensusStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyClientConsensusState(host common.Address, clientId string, height uint64, counterpartyClientIdentifier string, consensusHeight uint64, prefix []byte, proof []byte, consensusStateBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyClientConsensusState(&_Proxymultisigclient.CallOpts, host, clientId, height, counterpartyClientIdentifier, consensusHeight, prefix, proof, consensusStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0xf30fef52.
//
// Solidity: function verifyClientState(address host, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyClientState(opts *bind.CallOpts, host common.Address, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyClientState", host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyClientState is a free data retrieval call binding the contract method 0xf30fef52.
//
// Solidity: function verifyClientState(address host, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyClientState(host common.Address, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyClientState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyClientState is a free data retrieval call binding the contract method 0xf30fef52.
//
// Solidity: function verifyClientState(address host, string clientId, uint64 height, bytes prefix, string counterpartyClientIdentifier, bytes proof, bytes clientStateBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyClientState(host common.Address, clientId string, height uint64, prefix []byte, counterpartyClientIdentifier string, proof []byte, clientStateBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyClientState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, counterpartyClientIdentifier, proof, clientStateBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xd112d0b5.
//
// Solidity: function verifyConnectionState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyConnectionState(opts *bind.CallOpts, host common.Address, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyConnectionState", host, clientId, height, prefix, proof, connectionId, connectionBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xd112d0b5.
//
// Solidity: function verifyConnectionState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyConnectionState(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyConnectionState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyConnectionState is a free data retrieval call binding the contract method 0xd112d0b5.
//
// Solidity: function verifyConnectionState(address host, string clientId, uint64 height, bytes prefix, bytes proof, string connectionId, bytes connectionBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyConnectionState(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, connectionId string, connectionBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyConnectionState(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, connectionId, connectionBytes)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x63c0d662.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyPacketAcknowledgement(opts *bind.CallOpts, host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyPacketAcknowledgement", host, clientId, height, prefix, proof, portId, channelId, sequence, acknowledgement)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x63c0d662.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyPacketAcknowledgement(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyPacketAcknowledgement(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, sequence, acknowledgement)
}

// VerifyPacketAcknowledgement is a free data retrieval call binding the contract method 0x63c0d662.
//
// Solidity: function verifyPacketAcknowledgement(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes acknowledgement) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyPacketAcknowledgement(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, acknowledgement []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyPacketAcknowledgement(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, sequence, acknowledgement)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x08795a2d.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifyPacketCommitment(opts *bind.CallOpts, host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifyPacketCommitment", host, clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x08795a2d.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifyPacketCommitment(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyPacketCommitment(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifyPacketCommitment is a free data retrieval call binding the contract method 0x08795a2d.
//
// Solidity: function verifyPacketCommitment(address host, string clientId, uint64 height, bytes prefix, bytes proof, string portId, string channelId, uint64 sequence, bytes32 commitmentBytes) view returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifyPacketCommitment(host common.Address, clientId string, height uint64, prefix []byte, proof []byte, portId string, channelId string, sequence uint64, commitmentBytes [32]byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifyPacketCommitment(&_Proxymultisigclient.CallOpts, host, clientId, height, prefix, proof, portId, channelId, sequence, commitmentBytes)
}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCaller) VerifySignature(opts *bind.CallOpts, consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	var out []interface{}
	err := _Proxymultisigclient.contract.Call(opts, &out, "verifySignature", consensusState, multisig, signBytes)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Proxymultisigclient *ProxymultisigclientSession) VerifySignature(consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifySignature(&_Proxymultisigclient.CallOpts, consensusState, multisig, signBytes)
}

// VerifySignature is a free data retrieval call binding the contract method 0xe332d75c.
//
// Solidity: function verifySignature((bytes[],string,uint64) consensusState, (bytes[],uint64) multisig, bytes signBytes) pure returns(bool)
func (_Proxymultisigclient *ProxymultisigclientCallerSession) VerifySignature(consensusState ConsensusStateData, multisig MultiSignatureData, signBytes []byte) (bool, error) {
	return _Proxymultisigclient.Contract.VerifySignature(&_Proxymultisigclient.CallOpts, consensusState, multisig, signBytes)
}
