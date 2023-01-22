package block

import (
	"baby-chain/errors"
	"baby-chain/tools"
	"encoding/json"
	"fmt"
)

type Block struct {
	Header    tools.Data `json:"header"`
	Timestamp tools.Time `json:"timestamp"`
	PrevHash  tools.Hash `json:"prev_hash"`
	Data      tools.Data `json:"data"`
	Hash      tools.Hash `json:"hash"`
}

func (b *Block) MarshalJSON() ([]byte, error) {
	type alias Block
	if err := b.Validate(); err != nil {
		return nil, err
	}
	return json.Marshal((*alias)(b))
}

func (b *Block) UnmarshalJSON(save []byte) error {
	type alias Block
	aux := alias{}
	if err := json.Unmarshal(save, &aux); err != nil {
		return err
	}
	*b = Block(aux)
	if err := b.Validate(); err != nil {
		return err
	}
	return nil
}

func (b *Block) genHash() (tools.Hash, error) {
	dat, err := b.Data.MarshalJSON()
	if err != nil {
		return tools.HashB(), err
	}
	return tools.HashB([]byte{byte(b.Timestamp)}, b.PrevHash[:], dat), nil
}

func (b *Block) Validate() error {
	var errs []error
	if hash, err := b.genHash(); err != nil {
		return err
	} else if b.Hash != hash {
		errs = append(errs, errors.HashMismatch(fmt.Sprintf("block.Hash & block.genHash(): %x & %x", b.Hash, hash)))
	}
	if err := b.Header.Validate(); err != nil {
		errs = append(errs, fmt.Errorf("=== headerValidationErrors: %w\n===", err))
	}
	if _, ok := b.Header[Head]; !ok {
		errs = append(errs, errors.NoHead)
	}
	return errors.MultiError(errs, "blockValidationErrors")
}

func (b *Block) String() string {
	return fmt.Sprintf("Header: %s; Timestamp: %s; PrevHash: %s...; Hash: %s...; Data: %s",
		b.Header.String(), b.Timestamp.String(), b.PrevHash.Hex()[:8], b.Hash.Hex()[:8], b.Data.String())
}

func New(header tools.Data, timestamp tools.Time, prevHash tools.Hash, data tools.Data) Block {
	b := Block{header, timestamp, prevHash, data, tools.HashB()}
	if hash, err := b.genHash(); err != nil {
		panic(err)
	} else {
		b.Hash = hash
	}
	if err := b.Validate(); err != nil {
		panic(err)
	}
	return b
}
