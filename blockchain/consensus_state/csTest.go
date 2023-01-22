package consensus_state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	. "baby-chain/tools/data"
	"fmt"
)

const (
	TESTS   = "Tests"
	TestNum = "TestNum"
)

func Test(i int) string {
	return fmt.Sprintf("Test%d", i)
}

var CSTest = ConsensusState{
	func(_ *blockchain.Blockchain, _ *StateData, b block.Block) bool {
		if b.Header[block.Head] != block.TEST {
			return false
		}
		if _, ok := b.Data[TestNum].(string); !ok {
			return false
		}
		return true
	},
	func(bc *blockchain.Blockchain, sd *StateData, b block.Block) error {
		if err := bc.AddBlock(b); err != nil {
			return err
		}
		tests, ok := sd.Data[TESTS].(Data)
		if !ok {
			tests = Data{}
			sd.Data[TESTS] = tests
		}
		tests[b.Data[TestNum].(string)] = b.Data
		return nil
	},
}
