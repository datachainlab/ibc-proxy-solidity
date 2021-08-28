package types

import (
	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

func (cs *ClientState) ClientType() string {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) GetLatestHeight() exported.Height {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) Validate() error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) GetProofSpecs() []*ics23.ProofSpec {
	panic("not implemented") // TODO: Implement
}

// Initialization function
// Clients must validate the initial consensus state, and may store any client-specific metadata
// necessary for correct light client operation
func (cs *ClientState) Initialize(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.ConsensusState) error {
	panic("not implemented") // TODO: Implement
}

// Status function
// Clients must return their status. Only Active clients are allowed to process packets.
func (cs *ClientState) Status(ctx sdk.Context, clientStore sdk.KVStore, cdc codec.BinaryCodec) exported.Status {
	panic("not implemented") // TODO: Implement
}

// Genesis function
func (cs *ClientState) ExportMetadata(_ sdk.KVStore) []exported.GenesisMetadata {
	panic("not implemented") // TODO: Implement
}

// Update and Misbehaviour functions
func (cs *ClientState) CheckHeaderAndUpdateState(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Header) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) CheckMisbehaviourAndUpdateState(_ sdk.Context, _ codec.BinaryCodec, _ sdk.KVStore, _ exported.Misbehaviour) (exported.ClientState, error) {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) CheckSubstituteAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, subjectClientStore sdk.KVStore, substituteClientStore sdk.KVStore, substituteClient exported.ClientState) (exported.ClientState, error) {
	panic("not implemented") // TODO: Implement
}

// Upgrade functions
// NOTE: proof heights are not included as upgrade to a new revision is expected to pass only on the last
// height committed by the current revision. Clients are responsible for ensuring that the planned last
// height of the current revision is somehow encoded in the proof verification process.
// This is to ensure that no premature upgrades occur, since upgrade plans committed to by the counterparty
// may be cancelled or modified before the last planned height.
func (cs *ClientState) VerifyUpgradeAndUpdateState(ctx sdk.Context, cdc codec.BinaryCodec, store sdk.KVStore, newClient exported.ClientState, newConsState exported.ConsensusState, proofUpgradeClient []byte, proofUpgradeConsState []byte) (exported.ClientState, exported.ConsensusState, error) {
	panic("not implemented") // TODO: Implement
}

// Utility function that zeroes out any client customizable fields in client state
// Ledger enforced fields are maintained while all custom fields are zero values
// Used to verify upgrades
func (cs *ClientState) ZeroCustomFields() exported.ClientState {
	panic("not implemented") // TODO: Implement
}

// State verification functions
func (cs *ClientState) VerifyClientState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, counterpartyClientIdentifier string, proof []byte, clientState exported.ClientState) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyClientConsensusState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, counterpartyClientIdentifier string, consensusHeight exported.Height, prefix exported.Prefix, proof []byte, consensusState exported.ConsensusState) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyConnectionState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, connectionID string, connectionEnd exported.ConnectionI) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyChannelState(store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, prefix exported.Prefix, proof []byte, portID string, channelID string, channel exported.ChannelI) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketCommitment(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, commitmentBytes []byte) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketAcknowledgement(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64, acknowledgement []byte) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyPacketReceiptAbsence(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, sequence uint64) error {
	panic("not implemented") // TODO: Implement
}

func (cs *ClientState) VerifyNextSequenceRecv(ctx sdk.Context, store sdk.KVStore, cdc codec.BinaryCodec, height exported.Height, delayTimePeriod uint64, delayBlockPeriod uint64, prefix exported.Prefix, proof []byte, portID string, channelID string, nextSequenceRecv uint64) error {
	panic("not implemented") // TODO: Implement
}
