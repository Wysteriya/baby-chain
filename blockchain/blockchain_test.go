package blockchain

import (
	"baby-chain/tools"
	"baby-chain/tools/data"
	"testing"
)

func TestBlockchain_MineNode(t *testing.T) {
	// todo
}

func TestBC(t *testing.T) {
	blockchain1 := New(data.GoodTestData())
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", data.GoodTestData())), t)
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", data.GoodTestData())), t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}

func TestJson(t *testing.T) {
	blockchain1 := New(data.GoodTestData())
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", data.GoodTestData())), t)
	tools.TTestJson(blockchain1, t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}
