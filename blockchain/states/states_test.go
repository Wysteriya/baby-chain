package states

import (
	"reflect"
	"testing"
)

import (
	"blockchain/block"
)

func Test(t *testing.T) {
	sd := StateData{}
	stts := New(STest)
	stts.Init(&sd)
	b := block.MBlock(block.Data{"head": "Test"}, block.HashB(), block.Data{"test1": true})
	if err := stts.Exec(&sd, b); err != nil {
		t.Fatal(err)
	}
	if save, err := sd.Save(); err != nil {
		t.Fatal(err)
	} else if _sd, err := Load(save); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(sd, _sd) {
		t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", sd, _sd)
	}
	t.Log(sd)
}
