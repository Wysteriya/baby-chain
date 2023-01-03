package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/tools"
	"testing"
)

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
	bc := blockchain.New(tools.Data{}) // todo: use GoodData(...)
	b1 := bc.Chain[0]
	b2 := bc.MineBlock("BadGenesis", tools.Data{})
	TCons(CGenesis, bc, b1, b2, t)
}

func TestCaNode(t *testing.T) {
	bc := blockchain.New(tools.Data{})
	b1, _publicKey, _privateKey := bc.MineNode(tools.Data{})
	t.Logf("publicKey: %s\nprivateKey: %s", _publicKey, _privateKey)
	b2 := bc.MineBlock("BadNode", tools.Data{})
	TCons(CNode, bc, b1, b2, t)
}

func TestCAlgo_Exec(t *testing.T) {
	// todo
}
