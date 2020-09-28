package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Block struct
type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

// NewBlock create Block struct
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// Print Block struct
func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	fmt.Printf("transactions   %s\n", b.transactions)
}

// Blockchain struct
type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

// NewBlockchain create Blockchain struct
func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}

// CreateBlock create new block
func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

// Print Blockchain struct
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()
}
