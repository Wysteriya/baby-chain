package block

import (
	"crypto/sha256"
	"encoding/hex"
)

type Hash [32]byte

func (h Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func HashB(params ...[]byte) Hash {
	var hashee []byte
	for _, par := range params {
		for _, p := range par {
			hashee = append(hashee, p)
		}
	}
	return sha256.Sum256(hashee)
}
