pragma solidity ^0.6.8;
pragma experimental ABIEncoderV2;

import {IClient} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IClient.sol";
import {IBCHost} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/IBCHost.sol";
import {GoogleProtobufAny as Any} from "@hyperledger-labs/yui-ibc-solidity/contracts/core/types/GoogleProtobufAny.sol";
import {
    IbcLightclientsProxyV1ClientState as ProxyClientState,
    IbcLightclientsProxyV1ConsensusState as ProxyConsensusState
} from "../../types/proxy.sol";
import {
  ConsensusState as EthMultisigConsensusState,
  ClientState as EthMultisigClientState,
  Height
} from "../../types/ethmultisig.sol";
import {MultisigClient} from "./MultisigClient.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/ECRecovery.sol";
import "@hyperledger-labs/yui-ibc-solidity/contracts/lib/Bytes.sol";

contract ProxyMultisigClient is MultisigClient {
    using Bytes for bytes;

    protoTypes proxyPts;

    constructor() public {
      // TODO The typeUrl should be defined in types/proxy.sol
      // The schema of typeUrl follows cosmos/cosmos-sdk/codec/types/any.go
      proxyPts = protoTypes({
        clientState: keccak256(abi.encodePacked("/ibc.lightclients.proxy.v1.ClientState")),
        consensusState: keccak256(abi.encodePacked("/ibc.lightclients.proxy.v1.ConsensusState")),
        header: bytes32(0)
      });
    }

    /**
     * @dev checkHeaderAndUpdateState checks if the provided header is valid
     */
    function checkHeaderAndUpdateState(
        IBCHost host,
        string memory clientId,
        bytes memory clientStateBytes,
        bytes memory headerBytes
    ) public override view returns (bytes memory newClientStateBytes, bytes memory newConsensusStateBytes, uint64 height) {
      Any.Data memory anyClientState = Any.decode(clientStateBytes);
      require(keccak256(abi.encodePacked(anyClientState.type_url)) == proxyPts.clientState, "invalid client type");
      ProxyClientState.Data memory clientState = ProxyClientState.decode(anyClientState.value);

      (newClientStateBytes,newConsensusStateBytes, height) = super.checkHeaderAndUpdateState(host, clientId, Any.encode(clientState.proxy_client_state), headerBytes);
      clientState.proxy_client_state = Any.decode(newClientStateBytes);
      anyClientState.value = ProxyClientState.encode(clientState);
      Any.Data memory anyConsensusState = Any.Data({
        type_url: "/ibc.lightclients.proxy.v1.ConsensusState",
        value: ProxyConsensusState.encode(ProxyConsensusState.Data({
          proxy_consensus_state: Any.decode(newConsensusStateBytes)
        }))
      });
      return (Any.encode(anyClientState), Any.encode(anyConsensusState), height);
    }

    function verifyClientState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        string memory counterpartyClientIdentifier,
        bytes memory proof,
        bytes memory clientStateBytes // serialized with pb
    ) public override view returns (bool) {
        (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
        require(found, "client state not found");
        return super.verifyClientState(
          host,
          clientId,
          height,
          makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
          counterpartyClientIdentifier,
          proof,
          clientStateBytes
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
    ) public override view returns (bool) {
        (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
        require(found, "client state not found");
        return super.verifyClientConsensusState(
          host,
          clientId,
          height,
          counterpartyClientIdentifier,
          consensusHeight,
          makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
          proof,
          consensusStateBytes
        );
    }

    function verifyConnectionState(
        IBCHost host,
        string memory clientId,
        uint64 height,
        bytes memory prefix,
        bytes memory proof,
        string memory connectionId,
        bytes memory connectionBytes // serialized with pb
    ) public override view returns (bool) {
      (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
      require(found, "client state not found");
      return super.verifyConnectionState(
        host,
        clientId,
        height,
        makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
        proof,
        connectionId,
        connectionBytes
      );
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
    ) public override view returns (bool) {
      (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
      require(found, "client state not found");
      return super.verifyChannelState(
        host,
        clientId,
        height,
        makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
        proof,
        portId,
        channelId,
        channelBytes
      );
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
    ) public override view returns (bool) {
      (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
      require(found, "client state not found");
      return super.verifyPacketCommitment(
        host,
        clientId,
        height,
        makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
        proof,
        portId,
        channelId,
        sequence,
        commitmentBytes
      );
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
    ) public override view returns (bool) {
        (ProxyClientState.Data memory clientState, bool found) = getProxyClientState(host, clientId);
        require(found, "client state not found");
        return super.verifyPacketAcknowledgement(
            host,
            clientId,
            height,
            makeProxyCommitmentPrefix(clientState.proxy_prefix.key_prefix, clientState.upstream_client_id, prefix),
            proof,
            portId,
            channelId,
            sequence,
            acknowledgement
        );
    }

    function getProxyClientState(IBCHost host, string memory clientId) public virtual view returns (ProxyClientState.Data memory clientState, bool found) {
      bytes memory clientStateBytes;
      (clientStateBytes, found) = host.getClientState(clientId);
      if (!found) {
        return (clientState, false);
      }
      return (ProxyClientState.decode(Any.decode(clientStateBytes).value), true);
    }

    function getProxyConsensusState(IBCHost host, string memory clientId, Height.Data memory height) public virtual view returns (ProxyConsensusState.Data memory consensusState, bool found) {
      bytes memory consensusStateBytes;
      (consensusStateBytes, found) = host.getConsensusState(clientId, height.revision_height);
      if (!found) {
        return (consensusState, false);
      }
      return (ProxyConsensusState.decode(Any.decode(consensusStateBytes).value), true);
    }

    // override underlying functions

    function getClientState(IBCHost host, string memory clientId) public virtual override view returns (EthMultisigClientState.Data memory clientState, bool found) {
      ProxyClientState.Data memory proxyClientState;
      (proxyClientState, found) = getProxyClientState(host, clientId);
      if (!found) {
        return (clientState, false);
      }
      return (
        EthMultisigClientState.decode(proxyClientState.proxy_client_state.value),
        true
      );
    }

    // getConsensusState gets underlying consensus state
    function getConsensusState(IBCHost host, string memory clientId, Height.Data memory height) public virtual override view returns (EthMultisigConsensusState.Data memory consensusState, bool found) {
      ProxyConsensusState.Data memory proxyConsensusState;
      (proxyConsensusState, found) = getProxyConsensusState(host, clientId, height);
      if (!found) {
        return (consensusState, found);
      }
      return (
        EthMultisigConsensusState.decode(proxyConsensusState.proxy_consensus_state.value),
        true
      );
    }

    function makeProxyCommitmentPrefix(bytes memory proxyKeyPrefix, string memory upstreamClientId, bytes memory counterpartyPrefix) public pure returns (bytes memory) {
      return abi.encodePacked(proxyKeyPrefix, upstreamClientId, counterpartyPrefix);
    }
}
