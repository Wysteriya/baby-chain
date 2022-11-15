package block

import (
    "encoding/json"
    "fmt"
    "errors"
)

type Block struct {
    header string
    timestamp Time
    prevHash Hash
    hash Hash
    data Data
}

type block struct {
    Header *string `json:"header"`
    Timestamp *Time `json:"timestamp"`
    PrevHash *Hash `json:"prev_hash"`
    Hash *Hash `json:"hash"`
    Data *Data `json:"data"`
}

func (B *block) toBlock() Block {
    return Block{*B.Header, *B.Timestamp, *B.PrevHash, *B.Hash, *B.Data}
}

func (b *Block) MarshalJSON() ([]byte, error) {
    if err := b.Validate(); err != nil {
        return []byte{}, err
    }
    return json.Marshal(block{&b.header, &b.timestamp, &b.prevHash, &b.hash, &b.data})
}

func (b *Block) UnmarshalJSON(data []byte) error {
    B := block{}
    if err := json.Unmarshal(data, &B); err != nil {
        return err
    }
    *b = B.toBlock()
    if err := b.Validate(); err != nil {
        return err
    }
    return nil
}

func (b *Block) Hash() Hash {
    return HashB([]byte(b.header), []byte{byte(b.timestamp)}, []byte(b.prevHash.Hex()), []byte(b.data.String()))
}

func (b *Block) PrevHash() Hash {
    return b.prevHash
}

func (b *Block) Validate() error {
    if b.Hash() != b.hash {
        return errors.New("hash mismatch")
    }
    if err := b.data.Validate(); err != nil {
        return err
    }
    return nil
}

func (b *Block) Print() {
    fmt.Printf("Header: %s; Timestamp: %s; PrevHash: %s...; Hash: %s...;\nData: %s\n",
               b.header, b.timestamp.String(), b.prevHash.Hex()[:16], b.hash.Hex()[:16], b.data)
}

func (b *Block) Save() ([]byte, error) {
    return json.Marshal(b)
}

func New(header string, timestamp Time, prevHash Hash, data Data) Block {
    b := Block{header, timestamp, prevHash, HashB(), data}
    b.hash = b.Hash()
    return b
}

func Load(save []byte) (Block, error) {
    b := Block{}
    if err := json.Unmarshal(save, &b); err != nil {
        return Block{}, err
    }
    return b, nil
}

func MBlock(header string, prevHash Hash, data Data) Block {
    return New(header, CurrTime(), prevHash, data)
}

func Genesis(data Data) Block {
    return MBlock("Genesis", HashB(), data)
}
