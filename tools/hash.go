package tools

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
)

type Hash [32]byte

func (h Hash) Hex() string {
	return hex.EncodeToString(h[:])
}

func HashB(chunks ...[]byte) Hash {
	return sha256.Sum256(bytes.Join(chunks, nil))
}
