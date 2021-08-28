package testing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/suite"

	ethmultisigtypes "github.com/datachainlab/ibc-proxy-solidity/modules/light-clients/xx-ethmultisig/types"
	"github.com/datachainlab/ibc-proxy-solidity/pkg/consts"
	"github.com/datachainlab/ibc-proxy-solidity/pkg/contract/multisigclient"
	gethcrypto "github.com/ethereum/go-ethereum/crypto"
)

const testMnemonicPhrase = "math razor capable expose worth grape metal sunset metal sudden usage scheme"

type ProxyTestSuite struct {
	suite.Suite

	chain *Chain
}

func (suite *ProxyTestSuite) SetupTest() {
	chain := NewChain(suite.T(), "http://127.0.0.1:8545", testMnemonicPhrase, consts.Contract)
	suite.chain = chain
}

func (suite *ProxyTestSuite) TestMultisig() {
	ctx := context.TODO()

	var clientState ethmultisigtypes.ClientState
	_ = clientState

	var signBytes = []byte("signBytes")
	h := gethcrypto.Keccak256(signBytes)

	sig, err := gethcrypto.Sign(h, suite.chain.prvKey(0))
	suite.Require().NoError(err)

	consensusState := makeConsensusState(
		[]common.Address{suite.chain.CallOpts(ctx, 0).From},
		"tester",
		uint64(time.Now().Unix()),
	)
	multisigData := ethmultisigtypes.MultiSignature{
		Signatures: [][]byte{sig},
	}

	ok, err := suite.chain.multisigClient.VerifySignature(
		suite.chain.CallOpts(ctx, 0),
		multisigclient.ConsensusStateData(consensusState),
		multisigclient.MultiSignatureData(multisigData),
		signBytes,
	)
	suite.Require().NoError(err)
	suite.Require().True(ok)

	bz, err := suite.chain.multisigClient.MakeClientStateSignBytes(
		suite.chain.CallOpts(ctx, 0),
		1,
		uint64(time.Now().Unix()),
		"tester",
		"client-0",
		[]byte("dummy"),
	)
	suite.Require().NoError(err)
	fmt.Println(bz)
}

func makeConsensusState(addresses []common.Address, diversifier string, timestamp uint64) ethmultisigtypes.ConsensusState {
	var addrs [][]byte
	for _, addr := range addresses {
		addrs = append(addrs, addr[:])
	}
	return ethmultisigtypes.ConsensusState{
		Addresses:   addrs,
		Diversifier: diversifier,
		Timestamp:   timestamp,
	}
}

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ProxyTestSuite))
}
