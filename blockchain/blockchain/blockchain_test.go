package blockchain

import (
    "testing"
    "reflect"
)

import "blockchain/block"

func TestSaveLoad(t *testing.T) {
    bc := New(block.Data{"balances": map[string]interface {}{"amith": "1000", "yash": "500"}, "key": "10", "test": true})
    bc.MineBlock("Test", block.Data{"test1": true})
    bc.MineBlock("Test", block.Data{"test2": true})

    save, err := bc.Save()
    if err != nil {
        t.Fatalf("%s", err)
    }
    _bc, err := Load(save)
    if err != nil {
        t.Fatalf("%s", err)
    }
    if !reflect.DeepEqual(bc, _bc) {
        t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", bc, _bc)
    }
}
