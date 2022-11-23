package blockchain

import (
    "testing"
    "reflect"
    "os"
)

import "blockchain/blockchain"
import "blockchain/block"
import "blockchain/consensus"
import "db/jsoner"

func Test(t *testing.T) {
    testFile := "test.bin"
    bc := blockchain.New(block.Data{"balances": block.Data{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
    cons := consensus.New()
    if err := cons.Exec(&bc, bc.MineBlock("Test", block.Data{"test1": true})); err != nil {
        t.Fatalf("%s", err)
    } else if err := cons.Exec(&bc, bc.MineBlock("Test", block.Data{"test2": true})); err != nil {
        t.Fatalf("%s", err)
    } else if save, err := bc.Save(); err != nil {
        t.Fatalf("%s", err)
    } else if err := jsoner.WriteData("test.bin", save); err != nil {
        t.Fatalf("%s", err)
    } else if save, err := jsoner.ReadData(testFile); err != nil {
        t.Fatalf("%s", err)
    } else if _bc, err := blockchain.Load(save); err != nil {
        t.Fatalf("%s", err)
    } else if !reflect.DeepEqual(bc, _bc) {
        t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", bc, _bc)
    }
    os.Remove(testFile)
}
