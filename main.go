package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {

	bchain := blockchain.NewChain()

	block := bchain.NewBlock([]byte("Hello BlockChain"))
	block2 := bchain.NewBlock([]byte("Greetings"))

	bchain.AddBlock(*block)
	bchain.AddBlock(*block2)

	// block2.Content = []byte("Hello")

	fmt.Println(bchain.GetBlocks())
	fmt.Println(bchain.VerifyChain())
}
