package blockchain

import (
    "fmt"
    "errors"
    "encoding/json"
)

import "blockchain/block"

type Blockchain struct {
    chain []block.Block
}

type blockchain struct {
    Chain []block.Block `json:"chain"`
}

func (bc *Blockchain) MineBlock(header string, data block.Data) error {
    b := block.MBlock(header, bc.HashOf(bc.Len() - 1), data)
    bc.chain = append(bc.chain, b)
    return nil
}

func (bc *Blockchain) Len() int {
    return len(bc.chain)
}

func (bc *Blockchain) HashOf(i int) block.Hash {
    return bc.chain[i].Hash()
}

func (bc *Blockchain) Validate() bool {
    return true
}

func (bc *Blockchain) ValidateBlock(b block.Block) error {
    if bc.chain[bc.Len() - 1].Hash() == b.PrevHash() {
        return nil
    } else {
        return errors.New("chain hash mismatch")
    }
}

func (bc *Blockchain) Print() {
    fmt.Println("---")
    for _, block := range bc.chain {
        block.Print()
        fmt.Println("---")
    }
}

func (bc *Blockchain) Save() ([]byte, error) {
    return json.Marshal(blockchain{bc.chain})
}

func New(data block.Data) Blockchain {
    return Blockchain{[]block.Block{block.Genesis(data)}}
}

func Load(save []byte) (Blockchain, error) {
    load := blockchain{}
    err := json.Unmarshal(save, &load)
    return Blockchain{load.Chain}, err
}
