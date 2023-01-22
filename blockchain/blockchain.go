package blockchain

import (
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/wallet"
	"baby-chain/errors"
	"baby-chain/tools"
	. "baby-chain/tools/data"
	"encoding/json"
	"fmt"
)

type Blockchain struct {
	Fork  string        `json:"fork"`
	Chain []block.Block `json:"chain"`
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	type alias Blockchain
	if err := bc.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal((*alias)(bc))
}

func (bc *Blockchain) UnmarshalJSON(save []byte) error {
	type alias Blockchain
	aux := alias{}
	if err := json.Unmarshal(save, &aux); err != nil {
		return err
	}
	*bc = Blockchain(aux)
	if err := bc.Validate(); err != nil {
		return err
	}
	return nil
}

func (bc *Blockchain) Len() int {
	return len(bc.Chain)
}

func (bc *Blockchain) CurrHash() tools.Hash {
	return bc.Chain[bc.Len()-1].Hash
}

func (bc *Blockchain) Validate() error {
	if err := bc.Chain[0].Validate(); err != nil {
		return err
	}
	for i, b := range bc.Chain[1:] {
		if err := bc.ValidateBlock(i, b); err != nil {
			return err
		}
	}
	return nil
}

func (bc *Blockchain) ValidateBlock(i int, b block.Block) error {
	if err := b.Validate(); err != nil {
		return err
	}
	if bc.Chain[i].Hash != b.PrevHash {
		return errors.HashMismatch(fmt.Sprintf("chainHash & blockPrevHash @%d: %x & %x", i, bc.Chain[i].Hash, b.PrevHash))
	}
	return nil
}

func (bc *Blockchain) AddBlock(b block.Block) error {
	if err := bc.ValidateBlock(bc.Len()-1, b); err != nil {
		return err
	}
	bc.Chain = append(bc.Chain, b)
	return nil
}

func (bc *Blockchain) MineBlock(head string, data Data) block.Block {
	return block.MBlock(head, bc.CurrHash(), data)
}

func (bc *Blockchain) MineNode(data Data) (block.Block, string, string) {
	_publicKey, _privateKey, _ := wallet.GenerateKeys()
	return block.MNode(_publicKey, _privateKey, bc.CurrHash(), data), _publicKey, _privateKey
}

func (bc *Blockchain) StringChan() chan string {
	printer := make(chan string)
	go func() {
		printer <- bc.Fork + "\n"
		for _, b := range bc.Chain {
			printer <- b.String() + "\n"
		}
		close(printer)
	}()
	return printer
}

func New(data Data) Blockchain {
	return Blockchain{"AAA1", []block.Block{block.MGenesis(data)}}
}
