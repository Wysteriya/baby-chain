package consensus_state

import (
	"baby-chain/blockchain"
	"baby-chain/blockchain/block"
	"baby-chain/blockchain/wallet"
	"baby-chain/errors"
	. "baby-chain/tools/data"
	"encoding/hex"
)

// StateData : query keys will be upper-cased, value keys will be lower-cased, values can be any unicode
type StateData struct {
	Data
}

type ConsensusState struct {
	Check func(*blockchain.Blockchain, *StateData, block.Block) bool
	Run   func(*blockchain.Blockchain, *StateData, block.Block) error
}

type CSAlgo []ConsensusState

func (cs *CSAlgo) Exec(bc *blockchain.Blockchain, sd *StateData, b block.Block) error {
	atLeastOne := false
	if err := bc.ValidateBlock(bc.Len()-1, b); err != nil {
		return err
	}
	for _, con := range *cs {
		if con.Check(bc, sd, b) {
			if err := con.Run(bc, sd, b); err != nil {
				return err
			}
			atLeastOne = true
		}
	}
	if !atLeastOne {
		return errors.NoConsensusStateMatch
	}
	return nil
}

func (cs *CSAlgo) GenSD(bc *blockchain.Blockchain) (StateData, error) {
	newBc := blockchain.New(bc.Chain[0].Data)
	sd := NewSD()
	for _, b := range bc.Chain {
		if err := cs.Exec(&newBc, &sd, b); err != nil {
			return StateData{}, err
		}
	}
	return sd, nil
}

func NewSD() StateData {
	return StateData{Data{}}
}

func New(cons ...ConsensusState) CSAlgo {
	return append(CSAlgo{CSNode}, cons...)
}

func SignCheckBlock(b block.Block, signLabel string) error {
	_publicKey, ok := b.Data["public_key"].(string)
	if !ok {
		return errors.PublicKeyNotFound("can't check signature without public key")
	}
	_sign, ok := b.Header[signLabel].(string)
	if !ok {
		return errors.SignatureNotFound("can't find signature1")
	}
	sign, err := hex.DecodeString(_sign)
	if err != nil {
		return err
	}
	if !wallet.VerifySignature(_publicKey, b.Hash, sign) {
		return errors.InvalidSignatureHashPair
	}
	return nil
}

func GoodStateData() StateData {
	return StateData{GoodTestData()}
}
