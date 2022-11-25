package chain

import (
	"blockchain/block"
	"blockchain/blockchain"
	"blockchain/states"
	"fmt"
)

func Sync(bc *blockchain.Blockchain, sd *states.StateData) {
	nodes, _ := (*sd)["full_nodes"].(block.Data)
	for publicKey, ip := range nodes {
		fmt.Println(publicKey, ip)
	}
}
