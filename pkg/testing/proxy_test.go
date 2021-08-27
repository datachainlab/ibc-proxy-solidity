package testing

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/datachainlab/ibc-proxy-solidity/pkg/consts"
	"github.com/stretchr/testify/suite"
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

	// suite.Require().NoError(
	// 	suite.chain.TxSyncIfNoError(ctx)(
	// 		suite.chain.ibcHost.SetClientType(suite.chain.TxOpts(ctx, 0), "dummy-client", "dummy-ct"),
	// 	),
	// )

	bz, err := suite.chain.multisigClient.GetClientStateSignBytes(
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

func TestChainTestSuite(t *testing.T) {
	suite.Run(t, new(ProxyTestSuite))
}
