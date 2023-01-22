package consensus_state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/tools"
	. "baby-chain/tools/data"
	"testing"
)

// todo: improve multi b1 and b2
func TCons(con ConsensusState, bc blockchain.Blockchain, sd StateData, b1, b2 block.Block, t *testing.T) {
	if !con.Check(&bc, &sd, b1) {
		t.Fatalf("b1 check failed")
	} else {
		err := con.Run(&bc, &sd, b1)
		tools.TError(err, t)
	}
	if con.Check(&bc, &sd, b2) {
		t.Fatalf("b2 check passed")
	}
}

func TestCaNode(t *testing.T) {
	bc := blockchain.New(GoodTestData())
	sd := GoodStateData()
	b1, _publicKey, _privateKey := bc.MineNode(GoodTestData())
	t.Logf("publicKey: %s\nprivateKey: %s", _publicKey, _privateKey)
	b2 := bc.MineBlock("BadNode", GoodTestData())
	TCons(CSNode, bc, sd, b1, b2, t)
}

func TestCSAlgo_Exec(t *testing.T) {
	csAlgo := New(CSTest)
	bc := blockchain.New(GoodTestData())
	sd, err := csAlgo.GenSD(&bc)
	tools.TError(err, t)

	b, _, _ := bc.MineNode(GoodTestData())
	err = csAlgo.Exec(&bc, &sd, b)
	tools.TError(err, t)

	// todo CSTest
}

func TestCSAlgo_GenSD(t *testing.T) {
	// todo
}

func TestJson(t *testing.T) {
	sd := GoodStateData()
	tools.TTestJson(sd, t)
	t.Log(sd.String())
}
