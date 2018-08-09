package blockchain

import (
	"crypto/sha1"
	"fmt"
)

//Block construct
type Block struct {
	id                int
	hash              string
	previousBlockHash string
	Content           []byte
}

//BlockChain type
type BlockChain struct {
	currentID int
	blocks    []Block
}

func hash(ip []byte) string {
	sha1 := sha1.New()
	sha1.Write(ip)
	return fmt.Sprintf("%x", sha1.Sum(nil))
}

//CalculateHash Calculate the hash of the block
func (block *Block) calculateHash() {
	block.hash = hash([]byte(string(block.id) +
		string(block.previousBlockHash) +
		string(block.Content)))
}

//NewChain Create new chain
func NewChain() *BlockChain {
	return &BlockChain{
		currentID: 0,
		blocks:    make([]Block, 0),
	}
}

//NewBlock Create new block
func (blockchain *BlockChain) NewBlock(content []byte) *Block {
	return &Block{
		Content: content,
	}
}

//GetBlocks Get all blocks
func (blockchain *BlockChain) GetBlocks() []Block {
	return blockchain.blocks
}

//AddBlock Add the block to chain
func (blockchain *BlockChain) AddBlock(block Block) {
	blockchain.currentID = blockchain.currentID + 1
	block.id = blockchain.currentID
	block.calculateHash()
	blockchain.blocks = append(blockchain.blocks, block)
}
