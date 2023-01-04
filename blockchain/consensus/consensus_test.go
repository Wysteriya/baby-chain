package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/tools"
	"testing"
)

func GoodData() tools.Data {
	return tools.Data{
		"test1": "test",
		"test2": tools.Data{"test1": "1"},
		"test3": tools.Data{"test1": "1", "test2": tools.Data{"test1": "true"}},
	}
}

// todo: improve multi good's, and bad's
func TCons(con Consensus, bc blockchain.Blockchain, b1, b2 block.Block, t *testing.T) {
	if !con.Check(&bc, b1) {
		t.Fatalf("b1 check failed")
	} else {
		err := con.Run(&bc, b1)
		tools.TError(err, t)
	}
	if con.Check(&bc, b2) {
		t.Fatalf("b2 check passed")
	}
}

func TestCaGenesis(t *testing.T) {
	bc := blockchain.New(GoodData())
	b1 := bc.Chain[0]
	b2 := bc.MineBlock("BadGenesis", GoodData())
	TCons(CGenesis, bc, b1, b2, t)
}

func TestCaNode(t *testing.T) {
	bc := blockchain.New(GoodData())
	b1, _publicKey, _privateKey := bc.MineNode(GoodData())
	t.Logf("publicKey: %s\nprivateKey: %s", _publicKey, _privateKey)
	b2 := bc.MineBlock("BadNode", GoodData())
	TCons(CNode, bc, b1, b2, t)
}

func TestCAlgo_Exec(t *testing.T) {
	// todo
}
