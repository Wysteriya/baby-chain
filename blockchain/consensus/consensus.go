package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/wallet"
	"encoding/hex"
	"errors"
	"fmt"
)

type Consensus struct {
	Check func(*blockchain.Blockchain, block.Block) bool
	Run   func(*blockchain.Blockchain, block.Block) error
}

type CAlgo []Consensus

func (ca *CAlgo) Exec(bc *blockchain.Blockchain, b block.Block) error {
	atLeastOne := false
	if err := bc.ValidateBlock(bc.Len()-1, b); err != nil {
		return err
	}
	for _, con := range *ca {
		if con.Check(bc, b) {
			if err := con.Run(bc, b); err != nil {
				return err
			}
			atLeastOne = true
		}
	}
	if !atLeastOne {
		return errors.New("noConsensusMatch")
	}
	return nil
}

func New(cons ...Consensus) CAlgo {
	return append(CAlgo{CGenesis, CNode}, cons...)
}

func SignCheckBlock(b block.Block) error {
	_publicKey, ok := b.Data["public_key"].(string)
	if !ok {
		return fmt.Errorf("noPublicKeyFound")
	}
	_sign, ok := b.Header["signature1"].(string)
	if !ok {
		return fmt.Errorf("nosignature1Found")
	}
	sign, err := hex.DecodeString(_sign)
	if err != nil {
		return err
	}
	if !wallet.VerifySignature(_publicKey, b.Hash, sign) {
		return fmt.Errorf("invalidSignatureHashPair")
	}
	return nil
}
