package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TError(err error, t *testing.T) {
	if err != nil {
		t.Fatalf("unexpected error: got %s", err)
	}
}

func TExpectedError(err error, t *testing.T) {
	if err == nil {
		t.Fatal("expected error: but got nil")
	} else {
		t.Log(fmt.Sprintf("expected error: got %s", err))
	}
}

func TExpectedPanic(err any, t *testing.T) {
	if err == nil {
		t.Fatal("expected panic: but got nil")
	} else {
		t.Log(fmt.Sprintf("expected panic: got %s", err))
	}
}

func TTestJson[T any](t1 T, t *testing.T) {
	save, err := json.Marshal(&t1)
	var t2 T
	if err != nil {
		t.Error(err)
	} else if err := json.Unmarshal(save, &t2); err != nil {
		t.Error(err)
	} else if !reflect.DeepEqual(t1, t2) {
		t.Fatalf("Saved & Loaded are not equal: \n%#v\n%#v\n", t1, t2)
	}
}
