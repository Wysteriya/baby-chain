package state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/tools"
	"testing"
)

func GoodData() StateData {
	return StateData{tools.GoodTestData()}
}

// todo: improve multi good's, and bad's
func TStates(state State, sd StateData, b1, b2 block.Block, t *testing.T) {
	if !state.Check(&sd, b1) {
		t.Fatalf("b1 check failed")
	} else {
		err := state.Run(&sd, b1)
		tools.TError(err, t)
	}
	if state.Check(&sd, b2) {
		t.Fatalf("b2 check passed")
	}
	t.Log(sd.String())
}

func TestSaGenesis(t *testing.T) {
	sd := GoodData()
	bc := blockchain.New(tools.GoodTestData())
	b1 := bc.Chain[0]
	b2 := bc.MineBlock("BadGenesis", tools.GoodTestData())
	TStates(SGenesis, sd, b1, b2, t)
}

func TestSaNode(t *testing.T) {
	sd := GoodData()
	bc := blockchain.New(tools.GoodTestData())
	b1, _publicKey, _privateKey := bc.MineNode(tools.GoodTestData())
	t.Logf("publicKey: %s\nprivateKey: %s", _publicKey, _privateKey)
	b2 := bc.MineBlock("BadNode", tools.GoodTestData())
	TStates(SNode, sd, b1, b2, t)
}

func TestSAlgo_Exec(t *testing.T) {
	// todo
}

func TestJson(t *testing.T) {
	sd := GoodData()
	tools.TTestJson(sd, t)
	t.Log(sd.String())
}
