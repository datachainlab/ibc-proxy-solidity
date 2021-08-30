package testing

import (
	"context"
	"crypto/ecdsa"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	ethmultisigtypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
	proxytypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-proxy/types"
	"github.com/datachainlab/ibc-proxy-solidity/pkg/consts"
	"github.com/datachainlab/ibc-proxy-solidity/pkg/contract/multisigclient"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const testMnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ProxyTestSuite struct {
	suite.Suite

	chain *Chain
	cdc   codec.ProtoCodecMarshaler
}

func (suite *ProxyTestSuite) SetupTest() {
	suite.chain = NewChain(suite.T(), "http://127.0.0.1:8545", testMnemonicPhrase, consts.Contract)
	registry := codectypes.NewInterfaceRegistry()
	ethmultisigtypes.RegisterInterfaces(registry)
	proxytypes.RegisterInterfaces(registry)
	suite.cdc = codec.NewProtoCodec(registry)
}

func (suite *ProxyTestSuite) TestMultisig() {
	ctx := context.TODO()

	const diversifier = "tester"
	const (
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	prefix := []byte("ibc")
	ts := uint64(time.Now().Unix())

	prover := NewProver(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, prefix)
	targetClientState := makeMultisigClientState(1)
	signBytes, proofClient, err := prover.ProveClientState(counterpartyClientID, 1, ts, targetClientState)
	suite.Require().NoError(err)

	consensusState := makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		ts,
	)

	// VerifySignature

	ok, err := suite.chain.multisigClient.VerifySignature(
		suite.chain.CallOpts(ctx, 0),
		multisigclient.ConsensusStateData(*consensusState),
		multisigclient.MultiSignatureData(*proofClient),
		signBytes,
	)
	suite.Require().NoError(err)
	suite.Require().True(ok)

	// VerifyClientState

	anyConsensusStateBytes, err := suite.cdc.MarshalInterface(consensusState)
	suite.Require().NoError(err)
	suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.ibcHost.SetConsensusState(
			suite.chain.TxOpts(ctx, 0),
			clientID,
			1,
			anyConsensusStateBytes,
		))
	anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
	suite.Require().NoError(err)
	proofBytes, err := proto.Marshal(proofClient)
	suite.Require().NoError(err)
	ok, err = suite.chain.multisigClient.VerifyClientState(
		suite.chain.CallOpts(ctx, 0),
		suite.chain.ContractConfig.GetIBCHostAddress(),
		clientID, 1, prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
	)
	suite.Require().NoError(err)
	suite.Require().True(ok)
}

func makeProxyCommitmentPrefix(proxyKeyPrefix []byte, upstreamClientID string, counterpartyPrefix []byte) []byte {
	return append(append(proxyKeyPrefix, []byte(upstreamClientID)...), counterpartyPrefix...)
}

func (suite *ProxyTestSuite) TestProxy() {
	ctx := context.TODO()

	const diversifier = "tester"
	const (
		clientID                    = "testclient-0"
		counterpartyClientID        = "testcounterparty-0"
		upstreamClientID            = "testupstream-0"
		proofHeight          uint64 = 1
	)
	prefix := []byte("ibc")
	ts := uint64(time.Now().Unix())

	prover := NewProver(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, makeProxyCommitmentPrefix(prefix, upstreamClientID, prefix))

	proxyClient := makeProxyMultisigClientState(upstreamClientID, makeMultisigClientState(1000), prefix, prefix)
	proxyConsensus := makeProxyMultisigConsensusState(makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		ts,
	))
	anyProxyClientBytes, err := suite.cdc.MarshalInterface(proxyClient)
	suite.Require().NoError(err)
	anyProxyConsensusBytes, err := suite.cdc.MarshalInterface(proxyConsensus)
	suite.Require().NoError(err)
	suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.ibcHost.SetClientState(
			suite.chain.TxOpts(ctx, 0), clientID, anyProxyClientBytes,
		))
	suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.ibcHost.SetConsensusState(
			suite.chain.TxOpts(ctx, 0), clientID, proofHeight, anyProxyConsensusBytes,
		))

	targetClientState := makeMultisigClientState(1)
	_, proofClient, err := prover.ProveClientState(counterpartyClientID, proofHeight, ts, targetClientState)
	suite.Require().NoError(err)
	anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
	suite.Require().NoError(err)
	proofBytes, err := proto.Marshal(proofClient)
	suite.Require().NoError(err)

	ok, err := suite.chain.proxymultisigClient.VerifyClientState(
		suite.chain.CallOpts(ctx, 0),
		suite.chain.ContractConfig.GetIBCHostAddress(),
		clientID, proofHeight, prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
	)
	suite.Require().NoError(err)
	suite.Require().True(ok)
}

func (suite *ProxyTestSuite) TestSerialization() {
	const diversifier = "tester"
	const (
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	prefix := []byte("ibc")
	ts := uint64(time.Now().Unix())
	prover := NewProver(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, prefix)

	ctx := context.TODO()
	clientState := makeMultisigClientState(1)
	signBytes, _, err := prover.ProveClientState(counterpartyClientID, 1, ts, clientState)
	suite.Require().NoError(err)

	anyClientStateBytes, err := suite.cdc.MarshalInterface(clientState)
	suite.Require().NoError(err)
	expectedBytes, err := suite.chain.multisigClient.MakeClientStateSignBytes(
		suite.chain.CallOpts(ctx, 0),
		1,
		ts,
		diversifier,
		counterpartyClientID,
		anyClientStateBytes,
		prefix,
	)
	suite.Require().NoError(err)
	suite.Require().Equal(expectedBytes, signBytes)
}

func makeMultisigClientState(latestHeight uint64) *ethmultisigtypes.ClientState {
	return &ethmultisigtypes.ClientState{
		LatestHeight: ethmultisigtypes.Height{
			RevisionNumber: 0,
			RevisionHeight: latestHeight,
		},
	}
}

func makeMultisigConsensusState(addresses []common.Address, diversifier string, timestamp uint64) *ethmultisigtypes.ConsensusState {
	var addrs [][]byte
	for _, addr := range addresses {
		addrs = append(addrs, addr[:])
	}
	return &ethmultisigtypes.ConsensusState{
		Addresses:   addrs,
		Diversifier: diversifier,
		Timestamp:   timestamp,
	}
}

func makeProxyMultisigClientState(upstreamClientID string, multisigClientState *ethmultisigtypes.ClientState, proxyPrefix, ibcPrefix []byte) *proxytypes.ClientState {
	var (
		err error
	)
	ibcMerklePrefix := types.NewMerklePrefix(ibcPrefix)
	proxyMerklePrefix := types.NewMerklePrefix(proxyPrefix)
	clientState := proxytypes.ClientState{
		UpstreamClientId: upstreamClientID,
		IbcPrefix:        &ibcMerklePrefix,
		ProxyPrefix:      &proxyMerklePrefix,
		TrustedSetup:     true,
	}
	clientState.ProxyClientState, err = clienttypes.PackClientState(multisigClientState)
	if err != nil {
		panic(err)
	}
	return &clientState
}

func makeProxyMultisigConsensusState(multisigConsensusState *ethmultisigtypes.ConsensusState) *proxytypes.ConsensusState {
	proxyConsensusState, err := clienttypes.PackConsensusState(multisigConsensusState)
	if err != nil {
		panic(err)
	}
	return &proxytypes.ConsensusState{
		ProxyConsensusState: proxyConsensusState,
	}
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ProxyTestSuite))
}

type ETHMultisigProver struct {
	cdc         codec.ProtoCodecMarshaler
	diversifier string
	keys        []*ecdsa.PrivateKey
	prefix      []byte
}

func NewProver(cdc codec.ProtoCodecMarshaler, diversifier string, keys []*ecdsa.PrivateKey, prefix []byte) ETHMultisigProver {
	return ETHMultisigProver{cdc: cdc, diversifier: diversifier, keys: keys, prefix: prefix}
}

func (p ETHMultisigProver) ProveClientState(clientID string, height uint64, timestamp uint64, clientState exported.ClientState) ([]byte, *ethmultisigtypes.MultiSignature, error) {
	bz, err := p.cdc.MarshalInterface(clientState)
	if err != nil {
		return nil, nil, err
	}
	path, err := ClientCommitmentKey(p.prefix, clientID)
	if err != nil {
		return nil, nil, err
	}
	data, err := p.cdc.Marshal(&ethmultisigtypes.StateData{
		Path:  path,
		Value: bz,
	})
	if err != nil {
		return nil, nil, err
	}
	signBytes, err := p.cdc.Marshal(&ethmultisigtypes.SignBytes{
		Height:      ethmultisigtypes.Height{RevisionNumber: 0, RevisionHeight: height},
		Timestamp:   timestamp,
		Diversifier: p.diversifier,
		DataType:    ethmultisigtypes.CLIENT,
		Data:        data,
	})
	if err != nil {
		return nil, nil, err
	}
	signHash := gethcrypto.Keccak256(signBytes)

	var proof ethmultisigtypes.MultiSignature
	for _, key := range p.keys {
		sig, err := gethcrypto.Sign(signHash, key)
		if err != nil {
			return nil, nil, err
		}
		proof.Signatures = append(proof.Signatures, sig)
	}
	return signBytes, &proof, nil
}
