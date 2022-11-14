package blockchain

import (
    "testing"
)

import "blockchain/blockchain"
import "blockchain/block"

func Test(t *testing.T) {
    bc := blockchain.New(block.Data{})
    bc.MineBlock("Test", block.Data{"test": true})
    bc.MineBlock("Test", block.Data{"test": true})
}
