package main

import (
	"blockchain-go/blockchain"
	"fmt"
)

func main() {

	bchain := blockchain.NewChain()

	block1 := bchain.NewBlock([]byte("Hello BlockChain"))
	block2 := bchain.NewBlock([]byte("Greetings"))

	bchain.AddBlock(block1)
	bchain.AddBlock(block2)

	for _, block := range bchain.GetBlocks() {

		fmt.Println(block)
	}

	fmt.Println(bchain.VerifyChain())
}
