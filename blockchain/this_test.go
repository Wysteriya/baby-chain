package main

import (
	"os"
	"reflect"
	"testing"
)

import "blockchain/blockchain"
import "blockchain/block"
import "blockchain/consensus"
import "db/jsoner"

func Test(t *testing.T) {
	testFile := "test.bin"
	bc := blockchain.New(block.Data{"balances": block.Data{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
	cons := consensus.New(consensus.CTest)
	if err := cons.Exec(&bc, bc.MineBlock(block.Data{"head": "Test"}, block.Data{"test1": true})); err != nil {
		t.Fatalf("%s", err)
	} else if err := cons.Exec(&bc, bc.MineBlock(block.Data{"head": "Test"}, block.Data{"test2": true})); err != nil {
		t.Fatalf("%s", err)
	} else if save, err := bc.Save(); err != nil {
		t.Fatalf("%s", err)
	} else if err := jsoner.WriteData(testFile, save); err != nil {
		t.Fatalf("%s", err)
	} else if save, err := jsoner.ReadData(testFile); err != nil {
		t.Fatalf("%s", err)
	} else if _bc, err := blockchain.Load(save); err != nil {
		t.Fatalf("%s", err)
	} else if !reflect.DeepEqual(bc, _bc) {
		t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", bc, _bc)
	}
	if err := os.Remove(testFile); err != nil {
		t.Fatal(err)
	}
}
