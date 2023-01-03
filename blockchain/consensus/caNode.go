package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/wallet"
	"encoding/hex"
	"fmt"
)

var CNode = Consensus{
	func(_ *blockchain.Blockchain, b block.Block) bool {
		if b.Header["head"].(string) != "Node" {
			return false
		}
		if _, ok := b.Header["signature1"].(string); !ok {
			return false
		}
		if _, ok := b.Data["publicKey"].(string); !ok {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, b block.Block) error {
		_publicKey, _ := b.Data["publicKey"].(string)
		_sign, _ := b.Header["signature1"].(string)
		sign, err := hex.DecodeString(_sign)
		if err != nil {
			return err
		}
		if !wallet.VerifySignature(_publicKey, b.Hash, sign) {
			return fmt.Errorf("invalid signature-hash pair")
		}
		if err := bc.AddBlock(b); err != nil {
			return err
		}
		return nil
	},
}
