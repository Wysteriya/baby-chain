package Proof_of_work

import (
	"blockchain/block"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math/big"
)

//default difficulty value set to 5
var difficulty = 5

type PowService struct {
	blk              *block.Block
	difficultyTarget *big.Int
}

func NewPowService(block *block.Block) *PowService {
	difficultyTarget := big.NewInt(1)
	difficultyTarget.Lsh(difficultyTarget, uint(256-difficulty))
	returnObj := &PowService{block, difficultyTarget}
	return returnObj
}

func (p *PowService) InitializeNonce(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(p.blk.PrevHash().Hex()),
			[]byte((p.blk.Data()).String()),
			ToBytes(nonce),
			ToBytes(int64(difficulty)),
		},
		[]byte{},
	)
	return data
}

func (p *PowService) Run() (int64, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for true {
		dat := p.InitializeNonce(int64(nonce))
		hash = sha256.Sum256(dat)
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(p.difficultyTarget) == -1 {
			break
		} else {
			nonce++
		}
	}
	return int64(nonce), hash[:]
}

func (p *PowService) Validate() bool {
	var hashInt big.Int
	dat := p.InitializeNonce(p.blk.Nonce())
	hash := sha256.Sum256(dat)
	hashInt.SetBytes(hash[:])
	return hashInt.Cmp(p.difficultyTarget) == -1
}

//ToBytes - utility function to convert int64 to []byte
func ToBytes(number int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, number)
	if err != nil {
		fmt.Println(err)
	}
	return buff.Bytes()
}
