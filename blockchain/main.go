package main

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/db/jsoner"
	"blockchain/states"
	"fmt"
)

func main() {
	bcFile := "./db/blockchain.bin"
	sdFile := "./db/statedata.bin"

	bc := blockchain.New(block.Data{})
	if save, err := bc.Save(); err != nil {
		fmt.Printf("%s", err)
	} else if err := jsoner.WriteData(bcFile, save); err != nil {
		fmt.Printf("%s", err)
	}

	sd := states.StateData{}
	if save, err := sd.Save(); err != nil {
		fmt.Printf("%s", err)
	} else if err := jsoner.WriteData(sdFile, save); err != nil {
		fmt.Printf("%s", err)
	}
}
