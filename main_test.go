package main

import (
	"blockchain-go/blockchain"
	"testing"
)

func TestVerifyWhenValid(t *testing.T) {
	bchain := blockchain.NewChain()

	block1 := bchain.NewBlock([]byte("Hello BlockChain"))
	block2 := bchain.NewBlock([]byte("Greetings"))

	bchain.AddBlock(block1)
	bchain.AddBlock(block2)

	if !bchain.VerifyChain() {
		t.Error("Blockchain invalid without change in blocks")
	}
}

func TestVerifyWhenInValid(t *testing.T) {
	bchain := blockchain.NewChain()

	block1 := bchain.NewBlock([]byte("Hello BlockChain"))
	block2 := bchain.NewBlock([]byte("Greetings"))

	bchain.AddBlock(block1)
	bchain.AddBlock(block2)

	block2.Content = []byte("Content changed")

	if bchain.VerifyChain() {
		t.Error("Blockchain valid with change in blocks")
	}
}
