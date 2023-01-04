package blockchain

import (
	"baby-chain/tools"
	"testing"
)

// todo: improve
func TestBC(t *testing.T) {
	blockchain1 := New(tools.Data{"balances": tools.Data{"amith": "1000", "yash": "500"}, "key": "10", "test": "true"})
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.Data{"test1": "true"})), t)
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.Data{"test2": "true"})), t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}

// todo: improve
func TestJson(t *testing.T) {
	blockchain1 := New(tools.Data{"test1": "test", "test2": tools.Data{"test1": "0"}, "test3": "true"})
	tools.TError(blockchain1.AddBlock(blockchain1.MineBlock("Test", tools.Data{"test1": "1"})), t)
	tools.TTestJson(blockchain1, t)
	for str := range blockchain1.StringChan() {
		t.Logf(str)
	}
}
