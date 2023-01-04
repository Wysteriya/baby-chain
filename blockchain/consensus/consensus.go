package consensus

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"errors"
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
