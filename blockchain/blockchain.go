package blockchain

import (
	"sync"
)

type Blockchain struct {
	mu     sync.Mutex
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{Blocks: []*Block{genesisBlock()}}
}

func (bc *Blockchain) AddBlock(data string) {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := generateBlock(prevBlock, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) GetBlocks() []*Block {
	return bc.Blocks
}
