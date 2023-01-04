package state

import (
	"baby-chain/blockchain/block"
	"baby-chain/tools"
)

type StateData struct {
	tools.Data
}

type State struct {
	Check func(*StateData, block.Block) bool
	Run   func(*StateData, block.Block) error
}

type SAlgo []State

func (sa *SAlgo) Exec(sd *StateData, b block.Block) error {
	if err := sd.Validate(); err != nil {
		return err
	}
	if err := b.Validate(); err != nil {
		return err
	}
	for _, con := range *sa {
		if con.Check(sd, b) {
			if err := con.Run(sd, b); err != nil {
				return err
			}
		}
	}
	return nil
}

func New(states ...State) SAlgo {
	return append(SAlgo{SGenesis, SNode}, states...)
}
