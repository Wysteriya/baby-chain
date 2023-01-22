package block

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/tools"
	. "baby-chain/tools/data"
	"testing"
)

func GoodHeader() Data {
	d := GoodTestData()
	d[Head] = TEST
	return d
}

func BadHeader() Data {
	return BadTestData()
}

func TestMBlocks(t *testing.T) {
	block := MGenesis(GoodTestData())
	tools.TError(block.Validate(), t)

	_publicKey, _privateKey, err := wallet.GenerateKeys()
	tools.TError(err, t)
	block = MNode(_publicKey, _privateKey, tools.HashB(), GoodTestData())
	tools.TError(block.Validate(), t)
}

func TestBlock_Validate(t *testing.T) {
	defer func() {
		tools.TExpectedPanic(recover(), t)
	}()

	block := MNew(GoodHeader(), tools.HashB(), GoodTestData())
	tools.TError(block.Validate(), t)

	block.Hash = tools.HashB()
	block.Header = BadHeader()
	tools.TExpectedError(block.Validate(), t)
	block.Data = BadTestData()
	tools.TExpectedError(block.Validate(), t)

	_ = MNew(BadHeader(), tools.HashB(), BadTestData())
}

func TestJson(t *testing.T) {
	block := MNew(GoodHeader(), tools.HashB(), GoodTestData())
	tools.TTestJson(block, t)
	t.Log(block.String())
}
