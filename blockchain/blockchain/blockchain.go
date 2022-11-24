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
	Fork  string        `json:"fork"`
}

func (BC *blockchain) toBlockchain() Blockchain {
	return Blockchain{BC.Chain, BC.Fork}
}

func (bc *Blockchain) toblockchain() blockchain {
	return blockchain{bc.chain, bc.fork}
}

func (bc *Blockchain) MarshalJSON() ([]byte, error) {
	if err := bc.Validate(); err != nil {
		return []byte{}, err
	}
	return json.Marshal(bc.toblockchain())
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

func (bc *Blockchain) MineBlock(header block.Data, data block.Data) block.Block {
	return block.MBlock(header, bc.HashOf(bc.Len()-1), data)
}

func (bc *Blockchain) AddBlock(b block.Block) error {
	if err := bc.ValidateBlock(b, bc.Len()-1); err != nil {
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
	for i, b := range bc.chain[1:] {
		if err := bc.ValidateBlock(b, i); err != nil {
			return err
		}
	}
	return nil
}

func (bc *Blockchain) ValidateBlock(b block.Block, i int) error {
	if err := b.Validate(); err != nil {
		return err
	}
	if bc.HashOf(i) != b.PrevHash() {
		return errors.New(fmt.Sprintf("chain hash mismatch : %v & %v", bc.HashOf(i), b.PrevHash()))
	}
	return nil
}

func (bc *Blockchain) Print() {
	fmt.Println("---")
	for _, b := range bc.chain {
		b.Print()
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
