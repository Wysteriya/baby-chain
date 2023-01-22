package consensus_state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/errors"
	"baby-chain/tools"
	"net"
)

const (
	NODES   = "Nodes"
	Balance = "balance"
)

var CSNode = ConsensusState{
	func(_ *blockchain.Blockchain, _ *StateData, b block.Block) bool {
		if b.Header[block.Head].(string) != block.NODE {
			return false
		}
		if _, ok := b.Header[block.Signature(1)].(string); !ok {
			return false
		}
		if _, ok := b.Data[block.PublicKey].(string); !ok {
			return false
		}
		if _, ok := b.Data[block.IpAddress].(string); !ok {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, sd *StateData, b block.Block) error {
		ipAddress := b.Data[block.IpAddress].(string)
		if net.ParseIP(ipAddress) == nil {
			return errors.InvalidIP(ipAddress)
		}
		if err := SignCheckBlock(b, block.Signature(1)); err != nil {
			return err
		}
		if err := bc.AddBlock(b); err != nil {
			return err
		}

		nodes, ok := sd.Data[NODES].(tools.Data)
		if !ok {
			nodes = tools.Data{}
			sd.Data[NODES] = nodes
		}
		_publicKey := b.Data[block.PublicKey].(string)
		node, ok := nodes[_publicKey].(tools.Data)
		if !ok {
			node = tools.Data{Balance: "0"}
			nodes[_publicKey] = node
		}
		node[block.IpAddress] = b.Data[block.IpAddress]

		return nil
	},
}
