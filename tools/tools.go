package tools

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func MultiError(errs []error, label string) error {
	if len(errs) > 0 {
		err := errs[0]
		for _, err_ := range errs[1:] {
			err = fmt.Errorf("%w\n%s", err, err_)
		}
		err = fmt.Errorf("%s: \n%w", label, err)
		return err
	} else {
		return nil
	}
}

func Reverse[T any](arr []T) chan T {
	ret := make(chan T)
	go func() {
		for i := range arr {
			ret <- arr[len(arr)-1-i]
		}
		close(ret)
	}()
	return ret
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
