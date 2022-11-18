package blockchain

import (
	"encoding/json"
	"errors"
	"fmt"
)

import "blockchain/block"

type Blockchain struct {
	chain []block.Block
	fork  string
}

type blockchain struct {
	Chain []block.Block `json:"chain"`
	Fork  *string       `json:"fork"`
}

func (BC *blockchain) toBlockchain() Blockchain {
	return Blockchain{BC.Chain, *BC.Fork}
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	if err := bc.Validate(); err != nil {
		return []byte{}, err
	}
	return json.Marshal(blockchain{bc.chain, &bc.fork})
}

func (bc *Blockchain) UnmarshalJSON(data []byte) error {
	BC := blockchain{}
	if err := json.Unmarshal(data, &BC); err != nil {
		return err
	}
	*bc = BC.toBlockchain()
	if err := bc.Validate(); err != nil {
		return err
	}
	return nil
}

func (bc *Blockchain) MineBlock(header string, data block.Data) error {
	b := block.MBlock(header, bc.HashOf(bc.Len()-1), data)
	if err := bc.ValidateBlock(b); err != nil {
		return err
	}
	bc.chain = append(bc.chain, b)
	return nil
}

func (bc *Blockchain) Len() int {
	return len(bc.chain)
}

func (bc *Blockchain) HashOf(i int) block.Hash {
	return bc.chain[i].Hash()
}

func (bc *Blockchain) PrevHashOf(i int) block.Hash {
	return bc.chain[i].PrevHash()
}

func (bc *Blockchain) Validate() error {
	return nil
}

func (bc *Blockchain) ValidateBlock(b block.Block) error {
	if err := b.Validate(); err != nil {
		return err
	}
	if bc.HashOf(bc.Len()-1) == b.PrevHash() {
		return nil
	}
	return errors.New("chain hash mismatch")
}

func (bc *Blockchain) Print() {
	fmt.Println("---")
	for _, block := range bc.chain {
		block.Print()
		fmt.Println("---")
	}
}

func (bc *Blockchain) Save() ([]byte, error) {
	return json.Marshal(bc)
}

func New(data block.Data) Blockchain {
	return Blockchain{[]block.Block{block.Genesis(data)}, "AA1"}
}

func Load(save []byte) (Blockchain, error) {
	bc := Blockchain{}
	if err := json.Unmarshal(save, &bc); err != nil {
		return Blockchain{}, err
	}
	return bc, nil
}
