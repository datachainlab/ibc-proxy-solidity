package types

import "github.com/cosmos/ibc-go/modules/core/exported"

func (cs *ConsensusState) ClientType() string {
	panic("not implemented") // TODO: Implement
}

// GetRoot returns the commitment root of the consensus state,
// which is used for key-value pair verification.
func (cs *ConsensusState) GetRoot() exported.Root {
	panic("not implemented") // TODO: Implement
}

// GetTimestamp returns the timestamp (in nanoseconds) of the consensus state
func (cs *ConsensusState) GetTimestamp() uint64 {
	panic("not implemented") // TODO: Implement
}

func (cs *ConsensusState) ValidateBasic() error {
	panic("not implemented") // TODO: Implement
}
