package block

import (
	"baby-chain/blockchain/wallet"
	"baby-chain/tools"
	"encoding/hex"
	"fmt"
)

const (
	Head      = "head"
	IpAddress = "ip_address"
	PublicKey = "public_key"
	TEST      = "Test"
	NODE      = "Node"
	GENESIS   = "Genesis"
)

func Signature(i int) string {
	return fmt.Sprintf("signature_%d", i)
}

func MNew(header tools.Data, prevHash tools.Hash, data tools.Data) Block {
	return New(header, tools.CurrTime(), prevHash, data)
}

func MBlock(head string, prevHash tools.Hash, data tools.Data) Block {
	return MNew(tools.Data{Head: head}, prevHash, data)
}

func MTest(prevHash tools.Hash, data tools.Data) Block {
	return MBlock(TEST, prevHash, data)
}

func MGenesis(data tools.Data) Block {
	return MBlock(GENESIS, tools.HashB(), data)
}

func MNode(_publicKey string, _privateKey string, prevHash tools.Hash, data tools.Data) Block {
	data[PublicKey] = _publicKey
	data[IpAddress] = tools.GetOutboundIP()
	b := MBlock(NODE, prevHash, data)
	sign, _ := wallet.SignHash(_privateKey, b.Hash)
	b.Header[Signature(1)] = hex.EncodeToString(sign)
	return b
}
