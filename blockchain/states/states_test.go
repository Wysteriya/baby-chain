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
	if save, err := sd.Save(); err != nil {
		t.Fatal(err)
	} else if _sd, err := Load(save); err != nil {
		t.Fatal(err)
	} else if !reflect.DeepEqual(sd, _sd) {
		t.Fatalf("Saved and Loaded Data are not equal\n%#v\n%#v", sd, _sd)
	}
	sts := New(STest)
	b := block.MBlock(block.Data{"head": "Test"}, block.HashB(), block.Data{"test1": true})
	if err := sts.Exec(&sd, b); err != nil {
		t.Fatal(err)
	}
	t.Log(sd)
}
