package state

import (
	"baby-chain/blockchain/block"
	"baby-chain/tools"
)

var SNode = State{
	func(_ *StateData, b block.Block) bool {
		if b.Header["head"] != "Node" {
			return false
		}
		return true
	},
	func(sd *StateData, b block.Block) error {
		data, ok := sd.Data["Nodes"].(tools.Data)
		if !ok {
			data = tools.Data{}
			sd.Data["Nodes"] = data
		}
		_publicKey, _ := b.Data["public_key"].(string)
		data[_publicKey] = b.Data
		return nil
	},
}
