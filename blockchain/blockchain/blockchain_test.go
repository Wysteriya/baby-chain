package blockchain

import (
    "testing"
    "reflect"
)

import "blockchain/block"

func TestSaveLoad(t *testing.T) {
    bc := New(block.Data{"balances": block.Data{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
    if err := bc.AddBlock(bc.MineBlock("Test", block.Data{"test1": true})); err != nil {
        t.Fatalf("%s", err)
    } else if err := bc.AddBlock(bc.MineBlock("Test", block.Data{"test2": true})); err != nil {
        t.Fatalf("%s", err)
    } else if save, err := bc.Save(); err != nil {
        t.Fatalf("%s", err)
    } else if _bc, err := Load(save); err != nil {
        t.Fatalf("%s", err)
    } else if !reflect.DeepEqual(bc, _bc) {
        t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", bc, _bc)
    }
}
