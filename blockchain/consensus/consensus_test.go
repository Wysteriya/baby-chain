package consensus

import (
	"testing"
)

import "blockchain/block"
import "blockchain/blockchain"

func Test(t *testing.T) {
	bc := blockchain.New(block.Data{"balances": block.Data{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
	cons := New(CTest)
	if err := bc.AddBlock(bc.MineBlock(block.Data{"head": "Test"}, block.Data{"test1": true})); err != nil {
		t.Fatalf("%s", err)
	}
	b := bc.MineBlock(block.Data{"head": "Test"}, block.Data{"test2": true})
	if err := cons.Exec(&bc, b); err != nil {
		t.Fatalf("%s", err)
	}
}
