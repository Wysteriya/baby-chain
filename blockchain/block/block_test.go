package block

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/tools"
	"testing"
)

func GoodHeader() tools.Data {
	return tools.Data{
		"head":  "Test",
		"test1": "test",
		"test2": tools.Data{"test1": "1"},
		"test3": tools.Data{"test1": "1", "test2": tools.Data{"test1": "true"}},
	}
}

func GoodData() tools.Data {
	return tools.Data{
		"test1": "test",
		"test2": tools.Data{"test1": "1"},
		"test3": tools.Data{"test1": "1", "test2": tools.Data{"test1": "true"}},
	}
}

func BadData() tools.Data {
	return tools.Data{
		"test1": "test",
		"test2": tools.Data{"test1": 1},
		"test3": tools.Data{"test1": 1, "test2": tools.Data{"test1": true}},
	}
}

func BadHeader() tools.Data {
	return BadData()
}

func TestM(t *testing.T) {
	block := MGenesis(GoodData())
	tools.TError(block.Validate(), t)

	_publicKey, _privateKey, err := wallet.GenerateKeys()
	tools.TError(err, t)
	block = MNode(_publicKey, _privateKey, tools.HashB(), GoodData())
	tools.TError(block.Validate(), t)
}

func TestBlock_Validate(t *testing.T) {
	defer func() {
		tools.TExpectedPanic(recover(), t)
	}()

	block := MNew(GoodHeader(), tools.HashB(), GoodData())
	tools.TError(block.Validate(), t)

	block.Hash = tools.HashB()
	block.Header = BadHeader()
	block.Data = BadData()
	tools.TExpectedError(block.Validate(), t)

	_ = MNew(BadHeader(), tools.HashB(), BadData())
}

func TestJson(t *testing.T) {
	block := MNew(GoodHeader(), tools.HashB(), GoodData())
	tools.TTestJson(block, t)
	t.Log(block.String())
}
