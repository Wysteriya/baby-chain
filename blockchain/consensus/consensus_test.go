package consensus

import (
    "testing"
)

import "blockchain/block"

func TestHeader(t *testing.T) {
    cons := CAlgo{cGenesis("Genesis")}
    b := block.MBlock("Test", block.HashB(), block.Data{})
    if err := cons.Check(b); err != nil {
        t.Fatalf("%s", err)
    }
}
