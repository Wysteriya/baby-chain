package block

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/tools"
	"testing"
)

func GoodHeader() tools.Data {
	d := tools.GoodTestData()
	d["head"] = "Test"
	return d
}

func BadHeader() tools.Data {
	return tools.BadTestData()
}

func TestMBlocks(t *testing.T) {
	block := MGenesis(tools.GoodTestData())
	tools.TError(block.Validate(), t)

	_publicKey, _privateKey, err := wallet.GenerateKeys()
	tools.TError(err, t)
	block = MNode(_publicKey, _privateKey, tools.HashB(), tools.GoodTestData())
	tools.TError(block.Validate(), t)
}

func TestBlock_Validate(t *testing.T) {
	defer func() {
		tools.TExpectedPanic(recover(), t)
	}()

	block := MNew(GoodHeader(), tools.HashB(), tools.GoodTestData())
	tools.TError(block.Validate(), t)

	block.Hash = tools.HashB()
	block.Header = BadHeader()
	tools.TExpectedError(block.Validate(), t)
	block.Data = tools.BadTestData()
	tools.TExpectedError(block.Validate(), t)

	_ = MNew(BadHeader(), tools.HashB(), tools.BadTestData())
}

func TestJson(t *testing.T) {
	block := MNew(GoodHeader(), tools.HashB(), tools.GoodTestData())
	tools.TTestJson(block, t)
	t.Log(block.String())
}
