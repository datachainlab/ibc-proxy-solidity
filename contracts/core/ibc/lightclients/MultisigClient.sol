pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import {IClient} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IClient.sol";
import {IBCHost} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import {IBCIdentifier} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCIdentifier.sol";
import {GoogleProtobufAny as Any} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/types/GoogleProtobufAny.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/ECRecovery.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/Bytes.sol";
import {
    ClientState,
    ConsensusState,
    Header,
    MultiSignature,
    SignBytes,
    StateData,
    Height
} from "../../types/ethmultisig.sol";

// MultisigClient is a dialect of https://github.com/datachainlab/ibc-multisig-client
contract MultisigClient is IClient {
    using Bytes for bytes;

    struct protoTypes {
        bytes32 clientState;
        bytes32 consensusState;
        bytes32 header;
    }

    protoTypes pts;

    constructor() public {
        // TODO The typeUrl should be defined in types/ethmultisig.sol
        // The schema of typeUrl follows cosmos/cosmos-sdk/codec/types/any.go
        pts = protoTypes({
            clientState: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.ClientState")),
            consensusState: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.ConsensusState")),
            header: keccak256(abi.encodePacked("/ibc.lightclients.ethmultisig.v1.Header"))
        });
    }

    function getConsensusHeight(uint64 height) internal pure returns (Height.Data memory) {
      return Height.Data({revision_number: 0, revision_height: 1});
    }

    /**
     * @dev getTimestampAtHeight returns the timestamp of the consensus state at the given height.
     */
    function getTimestampAtHeight(
        IBCHost host,
        string memory clientId,
        uint64 height
    ) public override view returns (uint64, bool) {
        (ConsensusState.Data memory consensusState, bool found) = getConsensusState(host, clientId, getConsensusHeight(height));
        if (!found) {
            return (0, false);
        }
        return (consensusState.timestamp, true);
    }

    function getLatestHeight(
        IBCHost host,
        string memory clientId
    ) public override view returns (uint64, bool) {
        (ClientState.Data memory clientState, bool found) = getClientState(host, clientId);
        if (!found) {
            return (0, false);
        }
        return (clientState.latest_height.revision_height, true);
    }

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public virtual override view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height) {
        // TODO blocked by https://github.com/hyperledger-labs/yui-ibc-solidity/issues/50
        revert("UpdateClient is not supported");
    }

    function verifySignature(ConsensusState.Data memory consensusState, MultiSignature.Data memory multisig, bytes memory signBytes) public pure returns (bool) {
      // TODO implements
      return true;
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes // serialized with pb
    ) public virtual override view returns (bool) {
        ConsensusState.Data memory consensusState;
        bool found;
        (consensusState, found) = getConsensusState(host, clientId, getConsensusHeight(height));
        require(found, "consensus state not found");

        return verifySignature(
          consensusState,
          MultiSignature.decode(proof),
          getClientStateSignBytes(height, consensusState.timestamp, consensusState.diversifier, counterpartyClientIdentifier, clientStateBytes)
        );
    }

    function verifyClientConsensusState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        string memory counterpartyClientIdentifier,
        uint64 consensusHeight,
        bytes memory prefix,
        bytes memory proof,
        bytes memory consensusStateBytes // serialized with pb
    ) public virtual override view returns (bool) {
      revert("not implemented error");
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public virtual override view returns (bool) {
      revert("not implemented error");
    }

    function verifyChannelState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        bytes memory channelBytes // serialized with pb
    ) public virtual override view returns (bool) {
      revert("not implemented error");
    }

    function verifyPacketCommitment(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes32 commitmentBytes
    ) public virtual override view returns (bool) {
      revert("not implemented error");
    }

    function verifyPacketAcknowledgement(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory portId,
        string memory channelId,
        uint64 sequence,
        bytes memory acknowledgement
    ) public virtual override view returns (bool) {
      revert("not implemented error");
    }

    function getClientState(IBCHost host, string memory clientId) public virtual view returns (ClientState.Data memory clientState, bool found) {
      bytes memory clientStateBytes;
      (clientStateBytes, found) = host.getClientState(clientId);
      if (!found) {
        return (clientState, false);
      }
      return (ClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getConsensusState(IBCHost host, string memory clientId, Height.Data memory height) public virtual view returns (ConsensusState.Data memory consensusState, bool found) {
      bytes memory consensusStateBytes;
      (consensusStateBytes, found) = host.getConsensusState(clientId, height.revision_number);
      if (!found) {
        return (consensusState, false);
      }
      return (ConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    function parseHeader(bytes memory headerBytes) internal view returns (Header.Data memory) {
      Any.Data memory any = Any.decode(headerBytes);
      require(keccak256(abi.encodePacked(any.type_url)) == pts.header, "invalid header type");
      return Header.decode(any.value);
    }

    function getClientStateSignBytes(
      uint64 height,
      uint64 timestamp,
      string memory diversifier,
      string memory clientID,
      bytes memory clientState
    ) public pure returns (bytes memory) {
      return SignBytes.encode(
        SignBytes.Data({
          height: Height.Data({revision_number: 0, revision_height: height}),
          timestamp: timestamp,
          diversifier: diversifier,
          data_type: SignBytes.DataType.DATA_TYPE_CLIENT_STATE,
          data: StateData.encode(
            StateData.Data(
              {
                path: abi.encodePacked(IBCIdentifier.clientCommitmentKey(clientID)),
                value: clientState
              }
            )
          )
        })
      );
    }
}
