package testing

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	conntypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	chantypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/suite"

	ethmultisigtypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
	ethmultisig "github.com/datachainlab/ibc-proxy-solidity/modules/relay/ethmultisig"
	"github.com/datachainlab/ibc-proxy-solidity/pkg/consts"
	proxytypes "github.com/datachainlab/ibc-proxy/modules/light-clients/xx-proxy/types"
	proxymoduletypes "github.com/datachainlab/ibc-proxy/modules/proxy/types"
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
	proxymoduletypes.RegisterInterfaces(registry)
	suite.cdc = codec.NewProtoCodec(registry)
}

func (suite *ProxyTestSuite) TestMultisig() {
	ctx := context.TODO()

	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, prefix)

	consensusState := makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		uint64(time.Now().UnixNano()),
	)
	anyConsensusStateBytes, err := suite.cdc.MarshalInterface(consensusState)
	suite.Require().NoError(err)
	suite.chain.TxSyncIfNoError(ctx)(
		suite.chain.ibcHost.SetConsensusState(
			suite.chain.TxOpts(ctx, 0),
			clientID,
			1,
			anyConsensusStateBytes,
		))

	// VerifyClientState
	{
		targetClientState := makeMultisigClientState(1)
		proofClient, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
		suite.Require().NoError(err)
		anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofClient)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyClientState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyClientConsensusState
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(0, 100)
		proofConsensus, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, counterpartyClientID, consensusHeight.RevisionHeight, prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyConnectionState
	{
		const connectionID = "connection-0"
		targetConnection := conntypes.NewConnectionEnd(conntypes.INIT, counterpartyClientID, conntypes.NewCounterparty(clientID, connectionID, types.NewMerklePrefix([]byte("ibc"))), []*conntypes.Version{}, 0)
		proof, err := prover.SignConnectionState(proofHeight, connectionID, targetConnection)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		connectionBytes, err := suite.cdc.Marshal(&targetConnection)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyConnectionState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, connectionID, connectionBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	const portID, channelID, cpPortID, cpChannelID = "port-0", "channel-0", "port-1", "channel-1"

	// VerifyChannelstate
	{
		targetChannel := chantypes.NewChannel(chantypes.INIT, chantypes.UNORDERED, chantypes.NewCounterparty(cpPortID, cpChannelID), []string{"connection-0"}, "1")
		proof, err := prover.SignChannelState(proofHeight, portID, channelID, targetChannel)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		channelBytes, err := suite.cdc.Marshal(&targetChannel)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyChannelState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, channelBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketCommitment
	{
		commitment := sha256.Sum256([]byte("test"))
		proof, err := prover.SignPacketState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyPacketCommitment(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, 1, commitment,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketAcknowledgement
	{
		acknowledgement := []byte("ack")
		commitment := sha256.Sum256(acknowledgement)
		proof, err := prover.SignPacketAcknowledgementState(proofHeight, portID, channelID, 1, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.multisigClient.VerifyPacketAcknowledgement(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, 1, acknowledgement,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}
}

func (suite *ProxyTestSuite) TestProxy() {
	ctx := context.TODO()

	const (
		diversifier          = "tester"
		clientID             = "testclient-0"
		counterpartyClientID = "testcounterparty-0"
		upstreamClientID     = "testupstream-0"
	)
	proofHeight := clienttypes.NewHeight(0, 1)
	prefix := []byte("ibc")

	prover := ethmultisig.NewETHMultisig(suite.cdc, diversifier, []*ecdsa.PrivateKey{suite.chain.prvKey(0)}, makeProxyCommitmentPrefix(prefix, upstreamClientID, prefix))

	proxyClient := makeProxyMultisigClientState(upstreamClientID, makeMultisigClientState(1000), prefix, prefix)
	proxyConsensus := makeProxyMultisigConsensusState(makeMultisigConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		diversifier,
		uint64(time.Now().UnixNano()),
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
			suite.chain.TxOpts(ctx, 0), clientID, proofHeight.RevisionHeight, anyProxyConsensusBytes,
		))

	// VerifyClientState
	{
		targetClientState := makeMultisigClientState(1)
		proofClient, err := prover.SignClientState(proofHeight, counterpartyClientID, targetClientState)
		suite.Require().NoError(err)
		anyClientStateBytes, err := suite.cdc.MarshalInterface(targetClientState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofClient)
		suite.Require().NoError(err)

		ok, err := suite.chain.proxymultisigClient.VerifyClientState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, proofHeight.RevisionHeight, prefix, counterpartyClientID, proofBytes, anyClientStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyConsensusState
	{
		targetConsensusState := makeMultisigConsensusState(nil, "tester2", uint64(time.Now().UnixNano()))
		consensusHeight := clienttypes.NewHeight(0, 100)
		proofConsensus, err := prover.SignConsensusState(proofHeight, counterpartyClientID, consensusHeight, targetConsensusState)
		suite.Require().NoError(err)
		anyConsensusStateBytes, err := suite.cdc.MarshalInterface(targetConsensusState)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proofConsensus)
		suite.Require().NoError(err)
		ok, err := suite.chain.proxymultisigClient.VerifyClientConsensusState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, counterpartyClientID, consensusHeight.RevisionHeight, prefix, proofBytes, anyConsensusStateBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyConnectionState
	{
		const connectionID = "connection-0"
		targetConnection := conntypes.NewConnectionEnd(conntypes.INIT, counterpartyClientID, conntypes.NewCounterparty(clientID, connectionID, types.NewMerklePrefix([]byte("ibc"))), []*conntypes.Version{}, 0)
		proof, err := prover.SignConnectionState(proofHeight, connectionID, targetConnection)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		connectionBytes, err := suite.cdc.Marshal(&targetConnection)
		suite.Require().NoError(err)
		ok, err := suite.chain.proxymultisigClient.VerifyConnectionState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, connectionID, connectionBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	const portID, channelID, cpPortID, cpChannelID = "port-0", "channel-0", "port-1", "channel-1"

	// VerifyChannelstate
	{
		targetChannel := chantypes.NewChannel(chantypes.INIT, chantypes.UNORDERED, chantypes.NewCounterparty(cpPortID, cpChannelID), []string{"connection-0"}, "1")
		proof, err := prover.SignChannelState(proofHeight, portID, channelID, targetChannel)
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		channelBytes, err := suite.cdc.Marshal(&targetChannel)
		suite.Require().NoError(err)
		ok, err := suite.chain.proxymultisigClient.VerifyChannelState(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, channelBytes,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketCommitment
	{
		commitment := sha256.Sum256([]byte("test"))
		proof, err := prover.SignPacketState(proofHeight, portID, channelID, 2, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.proxymultisigClient.VerifyPacketCommitment(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, 2, commitment,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}

	// VerifyPacketAcknowledgement
	{
		acknowledgement := []byte("ack")
		commitment := sha256.Sum256(acknowledgement)
		proof, err := prover.SignPacketAcknowledgementState(proofHeight, portID, channelID, 2, commitment[:])
		suite.Require().NoError(err)
		proofBytes, err := proto.Marshal(proof)
		suite.Require().NoError(err)
		ok, err := suite.chain.proxymultisigClient.VerifyPacketAcknowledgement(
			suite.chain.CallOpts(ctx, 0),
			suite.chain.ContractConfig.GetIBCHostAddress(),
			clientID, 1, prefix, proofBytes, portID, channelID, 2, acknowledgement,
		)
		suite.Require().NoError(err)
		suite.Require().True(ok)
	}
}

func makeProxyCommitmentPrefix(proxyKeyPrefix []byte, upstreamClientID string, counterpartyPrefix []byte) []byte {
	return []byte(fmt.Sprintf("%s/%s/%s", proxyKeyPrefix, upstreamClientID, counterpartyPrefix))
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
