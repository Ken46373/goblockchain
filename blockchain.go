package main

import (
	"fmt"
	"log"
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

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
