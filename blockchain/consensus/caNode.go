package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"errors"
	"net"
)

var CNode = Consensus{
	func(_ *blockchain.Blockchain, b block.Block) bool {
		if b.Header["head"].(string) != "Node" {
			return false
		}
		if _, ok := b.Header["signature1"].(string); !ok {
			return false
		}
		if _, ok := b.Data["public_key"].(string); !ok {
			return false
		}
		if _, ok := b.Data["ip_address"].(string); !ok {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, b block.Block) error {
		ipAddress, _ := b.Data["ip_address"].(string)
		if net.ParseIP(ipAddress) == nil {
			return errors.New("invalidIpAddress")
		}
		if err := SignCheckBlock(b); err != nil {
			return err
		}
		if err := bc.AddBlock(b); err != nil {
			return err
		}
		return nil
	},
}
