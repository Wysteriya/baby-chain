package block

import (
	"baby-chain/tools"
	"testing"
)

func TestBlock_Validate(t *testing.T) {
	defer func() {
		tools.TExpectedPanic(recover(), t)
	}()

	block1 := MGenesis(tools.Data{})
	tools.TError(block1.Validate(), t)

	block1.Hash = tools.HashB()
	block1.Header = tools.Data{"test1": "test", "test2": tools.Data{"test1": 1}, "test3": true}
	tools.TExpectedError(block1.Validate(), t)

	_ = MGenesis(tools.Data{"test1": "test", "test2": tools.Data{"test1": 1}, "test3": true})
}

func TestJson(t *testing.T) {
	block1 := MGenesis(tools.Data{"test1": "test", "test2": tools.Data{"test1": "0"}, "test3": "true"})
	tools.TTestJson(block1, t)
	t.Log(block1.String())
}
