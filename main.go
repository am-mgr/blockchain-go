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

	fmt.Println(bchain.GetBlocks())
}
