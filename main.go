package main

import (
	"fmt"
)

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second After Genesis")
	chain.AddBlock("Third. After Genesis")
	chain.AddBlock("fourth,")
	chain.AddBlock("fifth block here")
	chain.AddBlock("6th.")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous hash:\t%x\n", block.PrevHash)
		fmt.Printf("Data in Block:\t%s\n", block.Data)
		fmt.Printf("Hash:\t\t%x\n", block.Hash)
		fmt.Println()
	}
}
