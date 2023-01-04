package blockchain

import (
	"baby-chain/tools"
	"testing"
)

func TestBlockchain_MineNode(t *testing.T) {
	// todo
}

func TestBC(t *testing.T) {
	blockchain1 := New(tools.GoodTestData())
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.GoodTestData())), t)
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.GoodTestData())), t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}

func TestJson(t *testing.T) {
	blockchain1 := New(tools.GoodTestData())
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.GoodTestData())), t)
	tools.TTestJson(blockchain1, t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}
