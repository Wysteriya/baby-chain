package main

import (
	"blockchain/block"
	"blockchain/blockchain"
	"db/jsoner"
	"fmt"
)

func main() {
	bcFile := "../db/blockchain.bin"
	bc := blockchain.New(block.Data{})
	if save, err := bc.Save(); err != nil {
		fmt.Printf("%s", err)
	} else if err := jsoner.WriteData(bcFile, save); err != nil {
		fmt.Printf("%s", err)
	}
}
